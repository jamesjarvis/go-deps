package main

import (
	"context"
	"fmt"
	"github.com/jamesjarvis/go-deps/resolve"
	"os"
)



// This is janky and mostly just to test if this thing works.
func main() {
	modules, err := resolve.ResolveGet(context.Background(), resolve.KnownImports, os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, dlRule := range modules.DownloadRules {
		fmt.Println("go_mod_download(")
		fmt.Printf("    name = \"%s\",\n", dlRule.Name)
		fmt.Printf("    module = \"%s\",\n", dlRule.Module)
		fmt.Printf("    version = \"%s\",\n", dlRule.Version)
		fmt.Println(")")
	}


	for _, module := range modules.ModuleRules {
		install := ""
		for _, i := range module.Installs {
			install += fmt.Sprintf("        \"%s\",\n", i)
		}

		deps := ""
		for _, dep := range module.Deps {
			deps += fmt.Sprintf("        \":%s\",\n", dep)
		}


		fmt.Println("go_module(")
		fmt.Printf("    name = \"%s\",\n", module.Name)
		fmt.Printf("    module = \"%s\",\n", module.Module)
		if module.Download != "" {
			fmt.Printf("    download = \"%s\",\n", ":" + module.Download)
		} else {
			fmt.Printf("    version = \"%s\",\n", module.Version)
		}
		if len(module.Installs) != 1 || module.Installs[0] != "" {
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
		fmt.Println()
	}
}
