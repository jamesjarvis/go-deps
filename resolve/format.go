package resolve

import "fmt"

// ModuleRules contains all the build rules for a given module
type ModuleRules struct {
	// Download represents the `download = ""` param for this module if any
	Download *DownloadRule
	// Version represents the version if we're not using Downlaod above
	Version string
	// Mods represents the `go_module()` rules for this module. Usually there's only one except when there's a cycle.
	Mods []*ModuleRule
}

// ModuleRule corresponds to a `go_module()` rule. We generate one of these for each ModulePart in the Module.
type ModuleRule struct {
	Name string
	Module string
	Installs []string
	Deps []string
	ExportedDeps []string
}

// DownloadRule represents the `go_mod_download()` rule for a module with cyclic deps.
type DownloadRule struct {
	Name string
	Module string
	Version string
}

// Print will print the rules for this module
func (module *ModuleRules) Print() {
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

		if module.Version == "" {
			module.Version = getVersion(modRule.Module)
		}
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