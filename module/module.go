package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/jamesjarvis/go-deps/host"
)

// Module is the module object we want to add to the project, essentially just the module path
// and any required information for fetching the module (such as version).
type Module struct {
	Path string
	Version string

	downloaded bool
}

func (m *Module) String() string {
	if m.Version == "" {
		return m.Path
	}
	return fmt.Sprintf("%s@%s", m.Path, m.Version)
}

// Download downloads the go module into a temporary directory
func (m *Module) Download() error {
	goTool := host.FindGoTool()
	dir := host.MustGetCacheDir()
	env := append(os.Environ(), "GO111MODULE=on", fmt.Sprintf("GOPATH=%s", dir))

	cmd := exec.Command(goTool, "get", m.String())
	cmd.Env = env
	cmd.Dir = dir
	if _, err := cmd.Output(); err != nil {
		return err
	}

	m.downloaded = true
	if m.Version == "" {
		// Find downloaded path
		moduleCachePath, err := m.FindPath()
		if err != nil {
			return err
		}
		m.Version = strings.Split(moduleCachePath, "@")[1]
	}

	return nil
}

// FindPath looks for the directory containing the code for the downloaded module.
func (m *Module) FindPath() (string, error) {
	cacheDir := host.MustGetCacheDir()
	pathSlice := []string{cacheDir, "pkg", "mod"}
	modulePathSlice := strings.Split(m.Path, "/")
	pathSlice = append(pathSlice, modulePathSlice[:len(modulePathSlice)-1]...)
	moduleCachePath := path.Join(pathSlice...)

	baseName := modulePathSlice[len(modulePathSlice)-1]

	var finalPiece string
	var exactMatch bool

	err := filepath.Walk(moduleCachePath, func(path string, info fs.FileInfo, err error) error {
		if exactMatch {
			// This short circuits if we have already found the final piece.
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			return nil
		}

		if info.Name() == fmt.Sprintf("%s@%s", baseName, m.Version) {
			finalPiece = info.Name()
			exactMatch = true
			return filepath.SkipDir
		}

		if strings.HasPrefix(info.Name(), baseName) {
			finalPiece = info.Name()
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	moduleCachePath = path.Join(moduleCachePath, finalPiece)

	return moduleCachePath, nil
}

func (m *Module) GetDependencies() ([]*Module, error) {
	if !m.downloaded {
		return nil, fmt.Errorf("module %s has not been downloaded yet", m.String())
	}

	modulePath, err := m.FindPath()
	if err != nil {
		return nil, fmt.Errorf("error while finding module path for %s", m.String())
	}

	if err = os.Chdir(modulePath); err != nil {
		return nil, fmt.Errorf("failed to change directory to %s", modulePath)
	}

	goTool := host.FindGoTool()
	env := append(os.Environ(), "GO111MODULE=on")

	cmd := exec.Command(goTool, "list", "-mod=readonly", "-m", "-json", "all")
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	cmd.Env = env
	cmd.Dir = modulePath
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(stderr)
		return nil, err
	}

	type module struct {
		Path, Version, Sum string
		Main             bool
		Indirect bool
		Replace            *struct {
			Path, Version string
		}
	}

	dec := json.NewDecoder(bytes.NewReader(out))
	modules := []*Module{}
	for dec.More() {
		mod := new(module)
		if err := dec.Decode(mod); err != nil {
			return nil, err
		}
		if mod.Main || mod.Indirect {
			continue
		}
		modules = append(modules, &Module{
			Path: mod.Path,
			Version: mod.Version,
		})
	}

	return modules, nil
}
