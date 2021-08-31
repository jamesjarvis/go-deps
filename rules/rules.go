package rules

import (
	"fmt"
	"os"
	"strings"

	resolve "github.com/jamesjarvis/go-deps/resolve/model"
)

const clearLineSequence = "\x1b[1G\x1b[2K"

// TODO(jpoole, jamesjarvis): We should probably be using the bazel tools Rule structs rather than our own ones here


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

func ruleName(path, suffix string) string {
	return 	strings.ReplaceAll(path, "/", ".") + suffix
}

func partName(part *resolve.ModulePart) string {
	displayIndex := len(part.Module.Parts) - part.Index

	if displayIndex > 0 {
		return ruleName(part.Module.Name, fmt.Sprintf("_%d", displayIndex))
	}
	return ruleName(part.Module.Name, "")
}

func toInstall(pkg *resolve.Package) string {
	install := strings.Trim(strings.TrimPrefix(pkg.ImportPath, pkg.Module), "/")
	if install == "" {
		return "."
	}
	return install
}

func GenerateModules(modules map[string]*resolve.Module, importPaths map[*resolve.Package]*resolve.ModulePart) ([]*ModuleRules, error) {
	processed := 0
	ret := make([]*ModuleRules, 0, len(modules))
	for _, m := range modules {
		rules := new(ModuleRules)
		ret = append(ret, rules)

		if len(m.Parts) > 1 {
			rules.Download = &DownloadRule{
				Name:    ruleName(m.Name, "_dl"),
				Module:  m.Name,
				Version: m.Version,
			}
		} else {
			rules.Version = m.Version
		}

		for _, part := range m.Parts {
			modRule := &ModuleRule{
				Name:   partName(part),
				Module: m.Name,
			}

			done := map[string]struct{}{}
			for pkg := range part.Packages {
				install := toInstall(pkg)
				// TODO(jpoole): we should probably just sort these alphabetically with an exception to put "." at the
				//  front
				if install == "." {
					modRule.Installs = append([]string{install}, modRule.Installs...)
				} else {
					modRule.Installs = append(modRule.Installs, toInstall(pkg))
				}
				for _, i := range pkg.Imports {
					dep := importPaths[i]
					depRuleName := partName(dep)
					if _, ok := done[depRuleName]; ok || dep.Module == m {
						continue
					}
					done[depRuleName] = struct{}{}

					modRule.Deps = append(modRule.Deps, depRuleName)
				}
			}

			// The last part is the namesake and should export the rest of the parts.
			if part.Index == len(m.Parts) {
				for _, part := range m.Parts[:(len(m.Parts)-1)] {
					modRule.ExportedDeps = append(modRule.ExportedDeps, partName(part))
				}
			}

			// Add them in reverse order so the namesake appears first
			rules.Mods = append([]*ModuleRule{modRule}, rules.Mods...)
		}
		processed++
		fmt.Fprintf(os.Stderr, "%sGenerating rules... %d of %d modules.", clearLineSequence, processed, len(modules))
	}
	fmt.Fprintln(os.Stderr)
	return ret, nil
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