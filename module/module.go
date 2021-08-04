package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/jamesjarvis/go-deps/host"
	"golang.org/x/mod/modfile"
)

// Module is the module object we want to add to the project, essentially just the module path
// and any required information for fetching the module (such as version).
type Module struct {
	Path string
	Version string

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
func (m *Module) Download() error {
	goTool := host.FindGoTool()
	dir := host.MustGetCacheDir()
	env := append(os.Environ(), fmt.Sprintf("GOPATH=%s", dir))

	cmd := exec.Command(goTool, "mod", "download", "-json", m.String())
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	cmd.Env = env
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to download module: %s: %w", stderr.String(), err)
	}

	type module struct {
		Path, Version, Info, GoMod, Zip, Dir, Sum, GoModSum string
	}

	mod := new(module)
	err = json.Unmarshal(out, mod)
	if err != nil {
		return fmt.Errorf("failed to unmarshal output: %w", err)
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

	fmt.Println(goMod.Require)

	return nil, nil
}
