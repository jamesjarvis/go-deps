package module

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/jamesjarvis/go-deps/host"
	"golang.org/x/mod/modfile"
)

// Module is the module object we want to add to the project, essentially just the module path
// and any required information for fetching the module (such as version).
type Module struct {
	Path string
	Version string

	Deps []*Module

	downloaded bool
	info string
	goMod string
	dir string
	sum string
	goModSum string
}

func (m *Module) String() string {
	if m.Version == "" {
		return m.Path
	}
	return fmt.Sprintf("%s@%s", m.Path, m.Version)
}

// Download downloads the go module into a temporary directory
func (m *Module) Download(ctx context.Context) error {
	goTool := host.FindGoTool()
	dir := host.MustGetCacheDir()
	env := append(os.Environ(), fmt.Sprintf("GOPATH=%s", dir))

	cmd := exec.CommandContext(ctx, goTool, "mod", "download", "-json", m.String())
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	cmd.Env = env
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil && len(out) == 0 {
		return fmt.Errorf("command failed: %s: %w", stderr.String(), err)
	}

	type module struct {
		Path, Version, Info, GoMod, Zip, Dir, Sum, GoModSum string
		Error string
	}

	mod := new(module)
	err = json.Unmarshal(out, mod)
	if err != nil {
		return fmt.Errorf("failed to unmarshal output: %w", err)
	}
	if mod.Error != "" {
		// This is a little bit hacky, but if it's failed because we didn't run go get,
		// we just do that here and retry.
		// Download downloads the go module into a temporary directory
		if strings.Contains(mod.Error, "not a known dependency") {
			cmd := exec.CommandContext(ctx, goTool, "get", m.String())
			cmd.Env = env
			cmd.Dir = dir
			if _, err := cmd.Output(); err != nil {
				return fmt.Errorf("failed to run go get: %w", err)
			}
			return m.Download(ctx)
		}
	
		return fmt.Errorf("failed to download module: %s", mod.Error)
	}

	m.downloaded = true
	if m.Version == "" {
		m.Version = mod.Version
	}
	m.info = mod.Info
	m.goMod = mod.GoMod
	m.goModSum = mod.GoModSum
	m.sum = mod.Sum
	m.dir = mod.Dir

	// Add self to cache
	storedModule := GlobalCache.SetModule(m)
	if storedModule != m {
		log.Printf("Dependencies change! We started with %s and now have %s", m, storedModule)
		m = storedModule
	}

	return nil
}

// GetDependencies returns the direct dependencies of Module m.
func (m *Module) GetDependencies() ([]*Module, error) {
	if !m.downloaded {
		return nil, fmt.Errorf("module %s has not been downloaded yet", m.String())
	}

	modulePath := m.goMod

	goModBytes, err := ioutil.ReadFile(modulePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read go.mod file: %w", err)
	}

	goMod, err := modfile.ParseLax(modulePath, goModBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.mod file: %w", err)
	}

	modules := []*Module{}
	for _, mod := range goMod.Require {
		if mod.Indirect {
			// We don't care about indirect modules tbh.
			continue
		}

		modules = append(modules, &Module{
			Path: mod.Mod.Path,
			Version: mod.Mod.Version,
		})
	}

	m.Deps = modules

	return modules, nil
}

func (m *Module) GetDependenciesRecursively(ctx context.Context) ([]*Module, error) {
	// We start a goroutine to pass the modules we want to fetch to.
	// This goroutine is then self populated by the dependencies it then fetches.
	// Each time it fetches one, it calls wg.Done, and each time it adds one, it
	// calls wg.Add.
	// Once the waitgroup is done, the channel is closed, killing the worker.
	// This implementation technically has a deadlock case, as if there are modules
	// with thousands of dependencies, it will get stuck waiting to send on the 
	// channel it is consuming from.
	allModules := []*Module{}
	modules := make(chan *Module, 1000)

	seenMap := map[string]struct{}{}

	var wg sync.WaitGroup
	var groupError error
	go func(ctx context.Context){
		for mod := range modules {
			ctx, _ = context.WithTimeout(ctx, 30*time.Second)
			if _, seen := seenMap[mod.String()]; seen {
				// We have seen this before...
				wg.Done()
				continue
			}
			if groupError != nil {
				// If we encounter an error from any one of the dependency fetchers, we short circuit.
				wg.Done()
				continue
			}
			err := mod.Download(ctx)
			if err != nil {
				groupError = err
				wg.Done()
				continue
			}
			fetchedModules, err := mod.GetDependencies()
			if err != nil {
				groupError = err
				wg.Done()
				continue
			}
			for _, fetchedMod := range fetchedModules {
				wg.Add(1)
				modules <- fetchedMod
			}
			allModules = append(allModules, fetchedModules...)
			// Mark this module as seen.
			seenMap[mod.String()] = struct{}{}
			wg.Done()
		}
	}(ctx)

	// Send initial module.
	wg.Add(1)
	modules <- m

	// Wait for results and close sending channel.
	wg.Wait()
	close(modules)
	
	if groupError != nil {
		return nil, groupError
	}
	return allModules, nil
}
