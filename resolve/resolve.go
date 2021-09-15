package resolve

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/go-licenses/licenses"
	"github.com/jamesjarvis/go-deps/progress"
	. "github.com/jamesjarvis/go-deps/resolve/model"

	"golang.org/x/tools/go/packages"
)

type Modules struct {
	Pkgs        map[string]*Package
	Mods        map[string]*Module
	ImportPaths map[*Package]*ModulePart
}

type resolver struct {
	*Modules
	moduleCounts   map[string]int
	rootModuleName string
	config *packages.Config
}

func newResolver(rootModuleName string, config *packages.Config) *resolver {
	return &resolver{
		Modules: &Modules{
			Pkgs:        map[string]*Package{},
			Mods:        map[string]*Module{},
			ImportPaths: map[*Package]*ModulePart{},
		},
		moduleCounts: map[string]int{},
		rootModuleName: rootModuleName,
		config: config,
	}
}

func (r *resolver) dependsOn(done map[*Package]struct{}, pkg *Package, module *ModulePart) bool {
	if _, ok := done[pkg]; ok {
		return false
	}
	done[pkg] = struct{}{}
	pkgModule, ok := r.ImportPaths[pkg]
	if !ok {
		panic("not okay")
	}
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
			if r.ImportPaths[i] != part {
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


	part := r.getOrCreateModulePart(r.GetModule(pkg.Module), pkg)
	part.Packages[pkg] = struct{}{}
	r.ImportPaths[pkg] = part
	part.Modified = true

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

func (r *resolver) addPackagesToModules(done map[*Package]struct{}) {
	processed := 0

	for _, pkg := range r.Pkgs {
		r.addPackageToModuleGraph(done, pkg)
		processed++
		progress.PrintUpdate("Building module graph... %d of %d packages.", processed, len(r.Pkgs))
	}
}

// UpdateModules resolves a `go get` style wildcard and updates the modules passed in to it
func UpdateModules(modules *Modules, getPaths []string) error {
	defer progress.Clear()
	pkgs, r, err := load(getPaths)
	if err != nil {
		return err
	}

	r.Modules = modules

	done := map[*Package]struct{}{}
	if modules != nil {
		for _, pkg := range modules.Pkgs {
			done[pkg] = struct{}{}
		}
	}


	r.resolve(pkgs)
	r.addPackagesToModules(done)

	if err := r.resolveModifiedPackages(); err != nil {
		return err
	}

	if err := r.setVersions(); err != nil {
		return err
	}

	if err := r.setLicence(pkgs); err != nil {
		return err
	}

	return nil
}

func load(getPaths []string) ([]*packages.Package, *resolver, error) {
	progress.PrintUpdate( "Analysing packages...")

	config := &packages.Config{
		Mode: packages.NeedImports|packages.NeedModule|packages.NeedName|packages.NeedFiles,
	}
	r := newResolver(getCurrentModuleName(config), config)

	pkgs, err := packages.Load(config, getPaths...)
	if err != nil {
		return nil, nil, err
	}

	return pkgs, r, nil
}

func (r *resolver) resolveModifiedPackages() error {
	var modifiedPackages []string
	for _, m := range r.Mods {
		for _, part := range m.Parts {
			if part.Modified {
				for pkg := range part.Packages {
					if !pkg.Resolved {
						modifiedPackages = append(modifiedPackages, pkg.ImportPath)
					}
				}
			}
		}
	}

	pkgs, err := packages.Load(r.config, modifiedPackages...)
	if err != nil {
		return err
	}

	r.resolve(pkgs)
	return nil
}

func (r *resolver) resolve(pkgs []*packages.Package) {
	for _, p := range pkgs {
		if len(p.GoFiles) + len(p.OtherFiles) == 0 {
			continue
		}
		pkg := r.GetPackage(p.PkgPath)
		pkg.Module = p.Module.Path
		if pkg.Module == "" {
			panic(fmt.Errorf("no module for %v", p.PkgPath))
		}

		newPackages := make([]*packages.Package, 0, len(p.Imports))
		for importName, importedPkg := range p.Imports {
			if _, ok := KnownImports[importName]; ok {
				continue
			}
			newPkg := r.GetPackage(importName)
			m := r.GetModule(p.Module.Path)
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
			if !newPkg.Resolved {
				newPackages = append(newPackages, importedPkg)
			}
		}
		pkg.Resolved = true
		r.resolve(newPackages)
	}
}

// GetPackage gets an existing package or creates a new one
func (mods *Modules) GetPackage(path string) *Package {
	if pkg, ok := mods.Pkgs[path]; ok {
		return pkg
	}
	pkg := &Package{ImportPath: path, Imports: []*Package{}}
	mods.Pkgs[path] = pkg
	return pkg
}

func (mods *Modules) GetModule(path string) *Module {
	m, ok := mods.Mods[path]
	if !ok {
		m = &Module{
			Name: path,
		}
		mods.Mods[path] = m
	}
	return m
}

func (r *resolver) setLicence(pkgs []*packages.Package) (err error) {
	c, _ := licenses.NewClassifier(0.9)

	done := 0 // start at 1 to ignore the root module
	packages.Visit(pkgs, nil, func(p *packages.Package) {
		if err != nil {
			return
		}
		if _, ok := r.Pkgs[p.PkgPath]; !ok  {
			return
		}
		m := r.Mods[p.Module.Path]
		if m.Licence != "" || m.Name == r.rootModuleName {
			return
		}

		done++
		progress.PrintUpdate("Adding licenses... %d of %d modules.", done, len(r.Mods))


		var pkgDir string
		switch {
		case len(p.GoFiles) > 0:
			pkgDir = filepath.Dir(p.GoFiles[0])
		case len(p.CompiledGoFiles) > 0:
			pkgDir = filepath.Dir(p.CompiledGoFiles[0])
		case len(p.OtherFiles) > 0:
			pkgDir = filepath.Dir(p.OtherFiles[0])
		default:
			// This package is empty - nothing to do.
			return
		}

		path, e := licenses.Find(pkgDir, c)
		if e != nil {
			err = fmt.Errorf("failed to find licence for %v in %v: %v", m.Name, pkgDir, err)
			return
		}
		name, _, e := c.Identify(path)
		if e != nil {
			err = fmt.Errorf("failed to identify licence %v: %v", path, err)
			return
		}
		m.Licence = name
	})
	return
}

func (r *resolver) setVersions() error {
	var moduleNames []string
	for _, m := range r.Mods {
		if m.Version != "" {
			continue
		}

		if m.Name == r.rootModuleName {
			continue
		}
		moduleNames = append(moduleNames, m.Name)
	}
	if len(moduleNames) == 0 {
		return nil
	}

	cmd := exec.Command("go", append([]string{"list", "-m"}, moduleNames...)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("failed to get module versions: %v\n%v", err, string(out)))
	}

	vs := strings.Split(string(out), "\n")
	for i, moduleVersion := range vs {
		if moduleVersion == "" {
			continue
		}

		progress.PrintUpdate("Setting versions... %d of %d modules.", i+1, len(vs))

		parts := strings.Split(moduleVersion, " ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid module version tuple: %v", moduleVersion))
		}
		r.Mods[parts[0]].Version = parts[1]
	}

	return nil
}