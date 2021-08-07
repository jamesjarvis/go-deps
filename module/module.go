package module

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/jamesjarvis/go-deps/host"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
)

const goModuleTemplateString = `
go_module(
	name = "{{ .GetName }}",
	module = "{{ .Path }}",
	version = "{{ .Version }}",
	deps = [
		{{- range .Deps }}
		"{{ .GetFullyQualifiedName }}",
		{{- end }}
	],
)
	`

var goModuleTemplater = template.Must(template.New("go_module").Parse(goModuleTemplateString))

// Module is the module object we want to add to the project, essentially just the module path
// and any required information for fetching the module (such as version).
type Module struct {
	Path string
	Version string
	Name string

	Deps []*Module

	downloaded bool
	nameWithVersion bool
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

func (m *Module) GetName() string {
	if m.Name != "" {
		return m.Name
	}
	splitPath := strings.Split(m.Path, "/")
	modName := splitPath[len(splitPath)-1]
	if m.nameWithVersion {
		return modName + semver.Major(m.Version)
	}
	return modName
}

func (m *Module) GetBuildPath() string {
	splitPath := strings.Split(m.Path, "/")
	pathMinusEnd := strings.Join(splitPath[:len(splitPath)-1], "/")
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/third_party/go/%s/BUILD", currentDir, pathMinusEnd)
}

func (m *Module) GetFullyQualifiedName() string {
	splitPath := strings.Split(m.Path, "/")
	pathMinusEnd := strings.Join(splitPath[:len(splitPath)-1], "/")
	buildDir := fmt.Sprintf("third_party/go/%s", pathMinusEnd)
	return "//" + buildDir + ":" + m.GetName()
}

func (m *Module) WriteGoModuleRule(wr io.Writer) error {
	return goModuleTemplater.Execute(wr, m)
}

// Download downloads the go module into a temporary directory
func (m *Module) Download(ctx context.Context) error {
	downloadedModule, err := host.GoModDownload(ctx, m.String())
	if err != nil {
		return fmt.Errorf("failed to download go module: %w", err)
	}

	m.downloaded = true
	if m.Version == "" {
		m.Version = downloadedModule.Version
	}
	m.info = downloadedModule.Info
	m.goMod = downloadedModule.GoMod
	m.goModSum = downloadedModule.GoModSum
	m.sum = downloadedModule.Sum
	m.dir = downloadedModule.Dir

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
