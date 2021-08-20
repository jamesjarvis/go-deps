package host

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type GoListResponse struct {
	Module string
	Imports []string
}

func GoListAll(ctx context.Context, path string) ([]string, error) {
	goTool := FindGoTool()
	cmd := exec.CommandContext(ctx, goTool, "list", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to go get %v: %v\n%v", path, err, string(out))
	}

	return strings.Split(strings.Trim(string(out), "\n"), "\n"), nil
}

func GoList(ctx context.Context, path string) (*GoListResponse, error) {
	goTool := FindGoTool()
	cmd := exec.CommandContext(ctx, goTool, "list", "-f", "{{.Module.Path}}\n{{ range .Imports }}{{.}} {{end}}", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to go get %v: %v\n%v", path, err, string(out))
	}

	lines := strings.Split(strings.Trim(string(out), " \n"), "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("go list didn't produce the expected number of lines")
	}

	resp := &GoListResponse{
		Module: lines[0],
	}

	if len(lines) == 2 {
		resp.Imports = strings.Split(lines[1], " ")
	}

	return resp, nil
}
