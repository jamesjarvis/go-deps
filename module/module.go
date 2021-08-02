package module

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jamesjarvis/go-deps/host"
)

// Module is the module object we want to add to the project, essentially just the module path
// and any required information for fetching the module (such as version).
type Module struct {
	Path string
	Version string
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
	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}
