package resolve

import (
	"encoding/json"
	"fmt"
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

func (pkg *Package) toInstall() string {
	install := strings.Trim(strings.TrimPrefix(pkg.ImportPath, pkg.Module), "/")
	if install == "" {
		return "."
	}
	return install
}

// Module represents a module. It includes all deps so actually represents a full module graph.
type Module struct {
	// The module name
	Name string

	Parts []*ModulePart
}

type ModulePart struct {
	Module *Module
	// The packages in this module
	Packages map[*Package]struct{}
	// The index of this module part
	Index int
}


type ModuleRules struct {
	Download *DownloadRule
	Version string
	Mods []*ModuleRule
}

type ModuleRule struct {
	Name string
	Module string
	Installs []string
	Deps []string
	ExportedDeps []string
}

type DownloadRule struct {
	Name string
	Module string
	Version string
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

func (r *resolver) dependsOn(pkg *Package, module *ModulePart, leftModule bool) bool {
	if leftModule && module == r.importPaths[pkg] {
		return true
	}
	if module != r.importPaths[pkg] {
		leftModule = true
	}
	for _, i := range pkg.Imports {
		if r.dependsOn(i, module, leftModule) {
			return true
		}
	}
	return false
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

	m := r.getModule(pkg.Module)


	var validPart *ModulePart
	for _, part := range m.Parts {
		valid := true
		for _, i := range pkg.Imports {
			if r.dependsOn(i, part, false) {
				valid = false
				break
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

	validPart.Packages[pkg] = struct{}{}
	r.importPaths[pkg] = validPart

	done[pkg] = struct{}{}
}

func getModuleName(config *packages.Config) string {
	pkgs, err := packages.Load(config, ".")
	if err != nil {
		panic(fmt.Errorf("failed to get root package name: %v", err))
	}
	return pkgs[0].Module.Path
}

// ResolveGet resolves a `go get` style wildcard into a graph of packages
func ResolveGet(getPaths []string) ([]*ModuleRules, error) {
	config := &packages.Config{
		Mode: packages.NeedImports|packages.NeedModule|packages.NeedName,
	}
	pkgs, err := packages.Load(config, getPaths...)
	if err != nil {
		return nil, err
	}

	r := &resolver{
		pkgs:         map[string]*Package{},
		modules:      map[string]*Module{},
		importPaths:  map[*Package]*ModulePart{},
		moduleCounts: map[string]int{},
		rootModuleName: getModuleName(config),
	}

	r.resolve(pkgs)

	done := make(map[*Package]struct{}, len(pkgs))
	for _, pkg := range r.pkgs {
		r.addPackageToModuleGraph(done, pkg)
	}

	modules := make([]*Module, 0, len(r.modules))
	for _, m := range r.modules {
		modules = append(modules, m)
	}


	ret := make([]*ModuleRules, 0, len(r.modules))
	for _, m := range r.modules {
		rules := new(ModuleRules)
		ret = append(ret, rules)

		if len(m.Parts) > 1 {
			rules.Download = &DownloadRule{
				Name:    ruleName(m.Name, "_dl"),
				Module:  m.Name,
				Version: getVersion(m.Name),
			}
		} else {
			rules.Version = getVersion(m.Name)
		}

		for _, part := range m.Parts {
			modRule := &ModuleRule{
				Name:   partName(part),
				Module: m.Name,
			}

			done := map[string]struct{}{}
			for pkg := range part.Packages {
				install := pkg.toInstall()
				// TODO(jpoole): we should probably just sort these alphabetically with an exception to put "." at the
				//  front
				if install == "." {
					modRule.Installs = append([]string{install}, modRule.Installs...)
				} else {
					modRule.Installs = append(modRule.Installs, pkg.toInstall())
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

			if part.Index == len(m.Parts) {
				// TODO(jpoole): is this a safe assumption? Can we have more than 2 parts ot a
				for _, part := range m.Parts[:(len(m.Parts)-1)] {
					modRule.ExportedDeps = append(modRule.ExportedDeps, partName(part))
				}
			}

			rules.Mods = append(rules.Mods, modRule)
		}
	}


	return ret, nil
}


func (r *resolver) resolve(pkgs []*packages.Package) {

	for _, p := range pkgs {
		pkg, _ := r.getOrCreatePackage(p.PkgPath)
		pkg.Module = p.Module.Path
		if pkg.Module == "" {
			panic(fmt.Errorf("no module for %v", p.PkgPath))
		}

		newPackages := make([]*packages.Package, 0, len(p.Imports))
		for i, iPkg := range p.Imports {
			if _, ok := KnownImports[i]; ok {
				continue
			}
			if strings.HasPrefix(i, "vendor/") {
				continue
			}
			newPkg, created := r.getOrCreatePackage(i)
			pkg.Imports = append(pkg.Imports, newPkg)
			if created {
				newPackages = append(newPackages, iPkg)
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