package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jamesjarvis/go-deps/resolve"
	"os"
	"os/exec"
	"strings"
)

type DownloadResponse struct {
	Version string
}

// This is janky and mostly just to test if this thing works.
func main() {
	modules, err := resolve.ResolveGet(context.Background(), resolve.KnownImports, os.Args[1])
	if err != nil {
		panic(err)
	}


	for _, module := range modules {
		install := ""
		for _, i := range module.Install {
			install += fmt.Sprintf("        \"%s\",\n", i)
		}

		deps := ""
		for _, dep := range module.Deps {
			deps += fmt.Sprintf("        \":%s\",\n", ruleName(dep.Name))
		}

		version := getVersion(module.Name)

		fmt.Println("go_module(")
		fmt.Printf("    name = \"%s\",\n", ruleName(module.Name))
		fmt.Printf("    module = \"%s\",\n", module.Name)
		fmt.Printf("    version = \"%s\",\n", version)
		if len(module.Install) != 1 || module.Install[0] != "" {
			fmt.Printf("    install = [\n")
			fmt.Printf(install)
			fmt.Printf("    ],\n")
		}
		if len(module.Deps) != 0 {
			fmt.Printf("    deps = [\n")
			fmt.Printf(deps)
			fmt.Printf("    ],\n")
		}

		fmt.Println(")")
	}
}

func getVersion(module string) string {
	cmd := exec.Command("go", "mod", "download", "--json", module)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	resp := new(DownloadResponse)
	err = json.Unmarshal(out, resp)
	if err != nil {
		panic(err)
	}
	return resp.Version
}

func ruleName(path string) string {
	return strings.ReplaceAll(path, "/", ".")
}