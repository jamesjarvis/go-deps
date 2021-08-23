package main

import (
	"context"
	"fmt"
	"github.com/jamesjarvis/go-deps/resolve"
	"os"
)



// This is janky and mostly just to test if this thing works.
func main() {
	rules, err := resolve.ResolveGet(context.Background(), resolve.KnownImports, os.Args[1])
	if err != nil {
		panic(err)
	}


	for _, module := range rules {
		if module.Download != nil {
			fmt.Println("go_mod_download(")
			fmt.Printf("    name = \"%s\",\n", module.Download.Name)
			fmt.Printf("    module = \"%s\",\n", module.Download.Module)
			fmt.Printf("    version = \"%s\",\n", module.Download.Version)
			fmt.Println(")")
			fmt.Println()
		}

		for _, modRule := range module.Mods {
			install := ""
			for _, i := range modRule.Installs {
				install += fmt.Sprintf("        \"%s\",\n", i)
			}

			deps := ""
			for _, dep := range modRule.Deps {
				deps += fmt.Sprintf("        \":%s\",\n", dep)
			}

			exportedDeps := ""
			for _, dep := range modRule.ExportedDeps {
				exportedDeps += fmt.Sprintf("        \":%s\",\n", dep)
			}

			fmt.Println("go_module(")
			fmt.Printf("    name = \"%s\",\n", modRule.Name)
			fmt.Printf("    module = \"%s\",\n", modRule.Module)
			if module.Download != nil {
				fmt.Printf("    download = \"%s\",\n", ":" + module.Download.Name)
			} else {
				fmt.Printf("    version = \"%s\",\n", module.Version)
			}
			if len(modRule.Installs) != 1 || modRule.Installs[0] != "." {
				fmt.Printf("    install = [\n")
				fmt.Printf(install)
				fmt.Printf("    ],\n")
			}
			if len(modRule.Deps) != 0 {
				fmt.Printf("    deps = [\n")
				fmt.Printf(deps)
				fmt.Printf("    ],\n")
			}
			if len(modRule.ExportedDeps) != 0 {
				fmt.Printf("    exported_deps = [\n")
				fmt.Printf(exportedDeps)
				fmt.Printf("    ],\n")
			}

			fmt.Println(")")
			fmt.Println()
		}
	}
}
