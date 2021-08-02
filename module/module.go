package module

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/jamesjarvis/go-deps/host"
)

// Module is the module object we want to add to the project, essentially just the module path
// and any required information for fetching the module (such as version).
type Module struct {
	Path string
	Version string
}

// Download downloads the go module into a temporary directory
func (m *Module) Download() error {
	goTool := host.FindGoTool()

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to determine working directory: %w", err)
	}
	dir := path.Join(currentDir, "tmp")

	env := append(os.Environ(), "GO111MODULE=on", fmt.Sprintf("GOPATH=%s", dir))

	cmd := exec.Command(goTool, "get", m.Path)
	cmd.Env = env
	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}
