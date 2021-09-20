package rules

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/tatskaari/go-deps/resolve"
	"github.com/tatskaari/go-deps/resolve/model"

	"github.com/bazelbuild/buildtools/build"
)

type BuildGraph struct {
	Modules *resolve.Modules

	ModFiles map[*model.Module]*BuildFile
	Files    map[string]*BuildFile
}

type BuildFile struct {
	File             *build.File
	ModRules         map[*model.ModulePart]*build.Rule
	ModDownloadRules map[*model.Module]*build.Rule

	usedNames     map[string]string
	partNames     map[*model.ModulePart]string
	downloadNames map[*model.Module]string
}

func NewGraph() *BuildGraph {
	return &BuildGraph{
		Modules: &resolve.Modules{
			Pkgs:        map[string]*model.Package{},
			Mods:        map[string]*model.Module{},
			ImportPaths: map[*model.Package]*model.ModulePart{},
		},
		ModFiles: map[*model.Module]*BuildFile{},
		Files:    map[string]*BuildFile{},
	}
}

func newFile(path string) (*BuildFile, error) {
	// Ignore errors here as the file doesn't have to exist
	data, _ := os.ReadFile(path)
	f, err := build.ParseBuild(path, data)
	if err != nil {
		return nil, err
	}

	return &BuildFile{
		File:             f,
		ModRules:         map[*model.ModulePart]*build.Rule{},
		ModDownloadRules: map[*model.Module]*build.Rule{},

		usedNames:        map[string]string{},
		downloadNames:    map[*model.Module]string{},
		partNames:        map[*model.ModulePart]string{},
	}, nil
}

func (g *BuildGraph) ReadRules(buildFile string) error {
	file, err := newFile(buildFile)
	if err != nil {
		return err
	}

	g.Files[buildFile] = file

	for _, rule := range file.File.Rules("go_module") {
		moduleName := rule.AttrString("module")
		module := g.Modules.GetModule(moduleName)
		g.ModFiles[module] = file

		pkgs := map[*model.Package]struct{}{}
		part := &model.ModulePart{
			Module:   g.Modules.GetModule(moduleName),
			Packages: pkgs,
			Index:    len(module.Parts) + 1,
		}
		file.ModRules[part] = rule
		file.usedNames[rule.Name()] = part.Module.Name
		file.partNames[part] = rule.Name()

		module.Version = rule.AttrString("version")

		install := getStrListList(rule, "install")
		if len(install) == 0 {
			install = []string{"."}
		}
		for _, i := range install {
			// Add these here to be resolved later
			if strings.HasSuffix(i, "...") {
				pkgPath := strings.TrimSuffix(strings.TrimSuffix(i, "..."), "/")
				part.InstallWildCards = append(part.InstallWildCards, pkgPath)
				continue
			}

			importPath := filepath.Join(moduleName, i)

			pkg := g.Modules.GetPackage(importPath)
			pkg.Module = moduleName

			part.Packages[pkg] = struct{}{}
			g.Modules.ImportPaths[pkg] = part
		}

		module.Parts = append(module.Parts, part)
	}

	for _, rule := range file.File.Rules("go_mod_download") {
		moduleName := rule.AttrString("module")
		module := g.Modules.GetModule(moduleName)
		file.ModDownloadRules[module] = rule

		file.usedNames[rule.Name()] = moduleName
		file.downloadNames[module] = rule.Name()

		module.Version = rule.AttrString("version")
	}

	return nil
}

func getStrListList(rule *build.Rule, attr string) []string {
	list, ok := rule.Attr(attr).(*build.ListExpr)
	if !ok {
		return nil
	}
	ret := make([]string, 0, len(list.List))
	for _, i := range list.List {
		ret = append(ret, i.(*build.StringExpr).Value)
	}
	return ret
}
