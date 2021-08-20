package main

import (
	"context"
	"fmt"
	"github.com/jamesjarvis/go-deps/resolve"
	"os"
	"path/filepath"
)

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

		fmt.Println("go_module(")
		fmt.Printf("    name = \"%s\",\n", ruleName(module.Name))
		fmt.Printf("    module = \"%s\",\n", module.Name)
		fmt.Printf("    version = \"latest\",\n")
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


func ruleName(path string) string {
	name := filepath.Base(path)
	if name == "v2" {
		name = filepath.Base(filepath.Dir(path)) + "." + name
	}
	return name
}