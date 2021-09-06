package rules

import (
	"os"
	"path/filepath"

	"github.com/jamesjarvis/go-deps/resolve"
	"github.com/jamesjarvis/go-deps/resolve/model"

	"github.com/bazelbuild/buildtools/build"
)

type BuildGraph struct {
	File *build.File
	Modules *resolve.Modules
	ModRules map[*model.ModulePart]*build.Rule
	ModDownloadRules map[*model.Module]*build.Rule
}

func ReadRules(buildFile string) (*BuildGraph, error) {
	data, err := os.ReadFile(buildFile)
	if err != nil {
		return nil, err
	}
	f, err := build.ParseBuild(buildFile, data)
	if err != nil {
		return nil, err
	}

	ret := &BuildGraph {
		File: f,
		Modules: &resolve.Modules{
			Pkgs:        map[string]*model.Package{},
			Mods:        map[string]*model.Module{},
			ImportPaths: map[*model.Package]*model.ModulePart{},
		},
		ModRules: map[*model.ModulePart]*build.Rule{},
		ModDownloadRules: map[*model.Module]*build.Rule{},
	}
	for _, rule := range f.Rules("go_module") {
		moduleName := rule.AttrString("module")
		module := ret.Modules.GetModule(moduleName)


		pkgs := map[*model.Package]struct{}{}
		part := &model.ModulePart{
			Module:   ret.Modules.GetModule(moduleName),
			Packages: pkgs,
			Index:    len(module.Parts)+1,
		}
		ret.ModRules[part] = rule

		install := getStrListList(rule, "install")
		if len(install) == 0 {
			install = []string{"."}
		}
		for _, i := range install {
			importPath := filepath.Join(moduleName, i)
			pkg := ret.Modules.GetPackage(importPath)
			pkg.Module = moduleName
			pkgs[pkg] = struct{}{}
			ret.Modules.ImportPaths[pkg] = part
		}

		module.Parts = append(module.Parts, part)
	}

	for _, rule := range f.Rules("go_mod_download") {
		moduleName := rule.AttrString("module")
		ret.ModDownloadRules[ret.Modules.GetModule(moduleName)] = rule
	}
	return ret, nil
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
