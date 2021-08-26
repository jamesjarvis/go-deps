package resolve

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/tools/go/packages"
)

type resolver struct {
	// pkgs is a map of import paths to their package
	pkgs           map[string]*Package
	modules        map[string]*Module
	importPaths    map[*Package]*ModulePart
	moduleCounts   map[string]int
	rootModuleName string
}

// Package represents a single package in some module
type Package struct {
	// The full import path of this package
	ImportPath string

	// The module name this package belongs to
	Module string

	// Any other packages this package imports
	Imports []*Package
}


// Module represents a module. It includes all deps so actually represents a full module graph.
type Module struct {
	// The module name
	Name string

	Version string

	Parts []*ModulePart
}

// ModulePart essentially corresponds to a `go_module()` rule that compiles some (or all) packages from that module. In
// most cases, there's one part per module except where we need to split it out to resolve a cycle.
type ModulePart struct {
	Module *Module
	// The packages in this module
	Packages map[*Package]struct{}
	// The index of this module part
	Index int
}


func ruleName(path, suffix string) string {
	return 	strings.ReplaceAll(path, "/", ".") + suffix
}

func partName(part *ModulePart) string {
	displayIndex := len(part.Module.Parts) - part.Index

	if displayIndex > 0 {
		return ruleName(part.Module.Name, fmt.Sprintf("_%d", displayIndex))
	}
	return ruleName(part.Module.Name, "")
}

func (r *resolver) dependsOn(done map[*Package]struct{}, pkg *Package, module *ModulePart) bool {
	if _, ok := done[pkg]; ok {
		return false
	}
	done[pkg] = struct{}{}
	pkgModule := r.importPaths[pkg]

	if module == pkgModule {
		return true
	}
	for pkg := range pkgModule.Packages {
		for _, i := range pkg.Imports {
			if r.dependsOn(done, i, module) {
				return true
			}
		}
	}

	return false
}

// getOrCreateModulePart gets or create a module part that we can add this package to without causing a cycle
func (r *resolver) getOrCreateModulePart(m *Module, pkg *Package) *ModulePart {
	var validPart *ModulePart
	for _, part := range m.Parts {
		valid := true
		done := map[*Package]struct{}{}
		for _, i := range pkg.Imports {
			// Check all the imports that leave the current part
			if r.importPaths[i] != part {
				if r.dependsOn(done, i, part) {
					valid = false
					break
				}
			}
		}
		if valid {
			validPart = part
			break
		}
	}
	if validPart == nil {
		validPart = &ModulePart{
			Packages: map[*Package]struct{}{},
			Module: m,
			Index: len(m.Parts) + 1,
		}
		m.Parts = append(m.Parts, validPart)
	}
	return validPart
}

func (r *resolver) addPackageToModuleGraph(done map[*Package]struct{}, pkg *Package) {
	if _, ok := done[pkg]; ok {
		return
	}

	for _, i := range pkg.Imports {
		r.addPackageToModuleGraph(done, i)
	}

	// We don't need to add the current module to the module graph
	if r.rootModuleName == pkg.Module {
		return
	}


	part := r.getOrCreateModulePart(r.getModule(pkg.Module), pkg)
	part.Packages[pkg] = struct{}{}
	r.importPaths[pkg] = part

	done[pkg] = struct{}{}
}

func getCurrentModuleName(config *packages.Config) string {
	pkgs, err := packages.Load(config, ".")
	if err != nil {
		panic(fmt.Errorf("failed to get root package name: %v", err))
	}
	return pkgs[0].Module.Path
}

func (r *resolver) addPackagesToModules() {
	processed := 0
	done := map[*Package]struct{}{}
	for _, pkg := range r.pkgs {
		r.addPackageToModuleGraph(done, pkg)
		processed++
		fmt.Fprintf(os.Stderr, "%sBuilding module graph... %d of %d packages.", clearLineSequence, processed, len(r.pkgs))
	}
}

func toInstall(pkg *Package) string {
	install := strings.Trim(strings.TrimPrefix(pkg.ImportPath, pkg.Module), "/")
	if install == "" {
		return "."
	}
	return install
}

func (r *resolver) generateModules() ([]*ModuleRules, error) {
	processed := 0
	ret := make([]*ModuleRules, 0, len(r.modules))
	for _, m := range r.modules {
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
					dep := r.importPaths[i]
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
		fmt.Fprintf(os.Stderr, "%sGenerating rules... %d of %d modules.", clearLineSequence, processed, len(r.modules))
	}
	fmt.Fprintln(os.Stderr)
	return ret, nil
}

// ResolveGet resolves a `go get` style wildcard into a graph of packages
func ResolveGet(getPaths []string) ([]*ModuleRules, error) {
	fmt.Fprintf(os.Stderr, "Analysing packages...")
	config := &packages.Config{
		Mode: packages.NeedImports|packages.NeedModule|packages.NeedName,
	}
	pkgs, err := packages.Load(config, getPaths...)
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(os.Stderr, " Done.\n")


	r := newResolver(getCurrentModuleName(config))

	r.resolve(pkgs)
	r.addPackagesToModules()
	return r.generateModules()
}

func newResolver(rootModuleName string) *resolver {
	return &resolver{
		pkgs:         map[string]*Package{},
		modules:      map[string]*Module{},
		importPaths:  map[*Package]*ModulePart{},
		moduleCounts: map[string]int{},
		rootModuleName: rootModuleName,
	}
}

const clearLineSequence = "\x1b[1G\x1b[2K"

func (r *resolver) resolve(pkgs []*packages.Package) {
	for _, p := range pkgs {
		pkg, _ := r.getOrCreatePackage(p.PkgPath)
		pkg.Module = p.Module.Path
		if pkg.Module == "" {
			panic(fmt.Errorf("no module for %v", p.PkgPath))
		}

		newPackages := make([]*packages.Package, 0, len(p.Imports))
		for importName, importedPkg := range p.Imports {
			if _, ok := KnownImports[importName]; ok {
				continue
			}
			if strings.HasPrefix(importName, "vendor/") {
				continue
			}
			newPkg, created := r.getOrCreatePackage(importName)
			m := r.getModule(p.Module.Path)
			m.Version = p.Module.Version

			if importedPkg.Module.Path != p.Module.Path {
				pkg.Imports = append(pkg.Imports, newPkg)
			}
			if created {
				newPackages = append(newPackages, importedPkg)
			}
		}
		r.resolve(newPackages)
	}
}

// getOrCreatePackage gets an existing package or creates a new one. Returns true when a new package was creawed.
func (r *resolver) getOrCreatePackage(path string) (*Package, bool) {
	if pkg, ok := r.pkgs[path]; ok {
		return pkg, false
	}
	pkg := &Package{ImportPath: path, Imports: []*Package{}}
	r.pkgs[path] = pkg
	return pkg, true
}

func (r *resolver) getModule(path string) *Module {
	m, ok := r.modules[path]
	if !ok {
		m = &Module{
			Name: path,
		}
		r.modules[path] = m
	}
	return m
}

type DownloadResponse struct {
	Version string
}

func getVersion(module string) string {
	cmd := exec.Command("go", "mod", "download", "--json", module)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("failed to get module version for %v: %v\n%v", module, err, string(out)))
	}

	resp := new(DownloadResponse)
	err = json.Unmarshal(out, resp)
	if err != nil {
		panic(err)
	}
	return resp.Version
}