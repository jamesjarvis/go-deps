package resolve

import (
	"fmt"
	"github.com/jamesjarvis/go-deps/resolve/model"
	"github.com/jamesjarvis/go-deps/rules"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/tools/go/packages"
)

const clearLineSequence = "\x1b[1G\x1b[2K"

type resolver struct {
	// pkgs is a map of import paths to their package
	pkgs           map[string]*model.Package
	modules        map[string]*model.Module
	importPaths    map[*model.Package]*model.ModulePart
	moduleCounts   map[string]int
	rootModuleName string
}

func newResolver(rootModuleName string) *resolver {
	return &resolver{
		pkgs:         map[string]*model.Package{},
		modules:      map[string]*model.Module{},
		importPaths:  map[*model.Package]*model.ModulePart{},
		moduleCounts: map[string]int{},
		rootModuleName: rootModuleName,
	}
}

func (r *resolver) dependsOn(done map[*model.Package]struct{}, pkg *model.Package, module *model.ModulePart) bool {
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
func (r *resolver) getOrCreateModulePart(m *model.Module, pkg *model.Package) *model.ModulePart {
	var validPart *model.ModulePart
	for _, part := range m.Parts {
		valid := true
		done := map[*model.Package]struct{}{}
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
		validPart = &model.ModulePart{
			Packages: map[*model.Package]struct{}{},
			Module: m,
			Index: len(m.Parts) + 1,
		}
		m.Parts = append(m.Parts, validPart)
	}
	return validPart
}

func (r *resolver) addPackageToModuleGraph(done map[*model.Package]struct{}, pkg *model.Package) {
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
	if pkgs[0].Module == nil {
		return ""
	}
	return pkgs[0].Module.Path
}

func (r *resolver) addPackagesToModules() {
	processed := 0
	done := map[*model.Package]struct{}{}
	for _, pkg := range r.pkgs {
		r.addPackageToModuleGraph(done, pkg)
		processed++
		fmt.Fprintf(os.Stderr, "%sBuilding module graph... %d of %d packages.", clearLineSequence, processed, len(r.pkgs))
	}
}


// ResolveGet resolves a `go get` style wildcard into a graph of packages
func ResolveGet(getPaths []string) ([]*rules.ModuleRules, error) {
	fmt.Fprintf(os.Stderr, "Analysing packages...")

	config := &packages.Config{
		Mode: packages.NeedImports|packages.NeedModule|packages.NeedName|packages.NeedFiles,
	}
	r := newResolver(getCurrentModuleName(config))

	pkgs, err := packages.Load(config, getPaths...)
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(os.Stderr, " Done.\n")


	r.resolve(pkgs)
	r.addPackagesToModules()
	r.setVersions()
	return rules.GenerateModules(r.modules, r.importPaths)
}

func (r *resolver) resolve(pkgs []*packages.Package) {
	for _, p := range pkgs {
		if len(p.GoFiles) + len(p.OtherFiles) == 0 {
			continue
		}
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
			newPkg, created := r.getOrCreatePackage(importName)
			m := r.getModule(p.Module.Path)
			m.Version = p.Module.Version
			if p.Module == nil {
				panic(fmt.Sprintf("no module for %v. Perhaps you need to go get something?", pkg.ImportPath))
			}
			if importedPkg.Module == nil {
				panic(fmt.Sprintf("no module for imported package %v. Perhaps you need to go get something?", importedPkg.PkgPath))
			}
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
func (r *resolver) getOrCreatePackage(path string) (*model.Package, bool) {
	if pkg, ok := r.pkgs[path]; ok {
		return pkg, false
	}
	pkg := &model.Package{ImportPath: path, Imports: []*model.Package{}}
	r.pkgs[path] = pkg
	return pkg, true
}

func (r *resolver) getModule(path string) *model.Module {
	m, ok := r.modules[path]
	if !ok {
		m = &model.Module{
			Name: path,
		}
		r.modules[path] = m
	}
	return m
}

func (r *resolver) setVersions() {
	var moduleNames []string
	for _, m := range r.modules {
		if m.Version != "" {
			continue
		}

		if m.Name == r.rootModuleName {
			continue
		}
		moduleNames = append(moduleNames, m.Name)
	}

	cmd := exec.Command("go", append([]string{"list", "-m"}, moduleNames...)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("failed to get module versions: %v\n%v", err, string(out)))
	}

	for _, moduleVersion := range strings.Split(string(out), "\n") {
		if moduleVersion == "" {
			continue
		}

		parts := strings.Split(moduleVersion, " ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid module version tuple: %v", moduleVersion))
		}
		r.modules[parts[0]].Version = parts[1]
	}
}