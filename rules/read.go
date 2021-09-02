package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jamesjarvis/go-deps/resolve"
	"github.com/jamesjarvis/go-deps/resolve/model"

	"github.com/bazelbuild/buildtools/build"
	"github.com/bazelbuild/buildtools/edit"
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

func (graph *BuildGraph) Save() error {
	for _, m := range graph.Modules.Mods {
		dlRule, ok := graph.ModDownloadRules[m]
		if len(m.Parts) > 1 {
			if !ok {
				dlRule = NewRule(graph.File, "go_mod_download", downloadRuleName(m))
			}
			dlRule.SetAttr("version", NewStringExpr(m.Version))
			dlRule.SetAttr("module", NewStringExpr(m.Name))
			dlRule.SetAttr("licences", NewStringList(m.Licence))
		} else if ok {
			graph.File.DelRules(dlRule.Kind(), dlRule.Name())
		}

		for _, part := range m.Parts {
			if !part.Modified {
				continue
			}
			modRule, ok := graph.ModRules[part]
			if !ok {
				modRule = NewRule(graph.File, "go_module", partName(part))
			}
			modRule.DelAttr("version")
			modRule.DelAttr("download")
			modRule.DelAttr("install")
			modRule.DelAttr("deps")
			modRule.DelAttr("exported_deps")

			modRule.SetAttr("module", NewStringExpr(m.Name))

			if len(m.Parts) > 1 {
				modRule.SetAttr("download", NewStringExpr(downloadRuleName(m)))
			} else {
				modRule.SetAttr("licences", NewStringList(m.Licence))
				modRule.SetAttr("version", NewStringExpr(m.Version))
			}

			installs := make([]string, 0, len(part.Packages))
			deps := make([]string, 0, len(part.Packages))
			var exportedDeps []string

			doneDeps := map[string]struct{}{}
			for pkg := range part.Packages {
				installs = append(installs, toInstall(pkg))

				for _, i := range pkg.Imports {
					dep := graph.Modules.ImportPaths[i]
					depRuleName := partName(dep)
					if _, ok := doneDeps[depRuleName]; ok || dep.Module == m {
						continue
					}
					doneDeps[depRuleName] = struct{}{}
					deps = append(deps, ":" + depRuleName)
				}
			}

			// The last part is the namesake and should export the rest of the parts.
			if part.Index == len(m.Parts) {
				for _, part := range m.Parts[:(len(m.Parts)-1)] {
					exportedDeps = append(exportedDeps, ":" + partName(part))
				}
			}


			if len(installs) > 1 || (len(installs) == 1 && installs[0] != ".") {
				modRule.SetAttr("install", NewStringList(installs...))
			}

			if len(deps) > 0 {
				modRule.SetAttr("deps", NewStringList(deps...))
			}

			if len(exportedDeps) > 0 {
				modRule.SetAttr("exported_deps", NewStringList(exportedDeps...))
			}

		}
	}
	fmt.Println(string(build.Format(graph.File)))
	return nil
}


func NewRule(f *build.File, kind, name string) *build.Rule {
	rule, _ := edit.ExprToRule(&build.CallExpr{
		X: &build.Ident{Name: kind},
		List: []build.Expr{},
	}, kind)

	rule.SetAttr("name", NewStringExpr(name))

	f.Stmt = append(f.Stmt, rule.Call)
	return rule
}

func NewStringExpr(s string) *build.StringExpr {
	return &build.StringExpr{Value: s}
}

func NewStringList(ss ...string) *build.ListExpr {
	l := new(build.ListExpr)
	for _, s := range ss{
		l.List = append(l.List, NewStringExpr(s))
	}
	return l
}