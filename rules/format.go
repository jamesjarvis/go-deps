package rules

import (
	"fmt"
	"strings"

	resolve "github.com/jamesjarvis/go-deps/resolve/model"

	"github.com/bazelbuild/buildtools/build"
	"github.com/bazelbuild/buildtools/edit"
)


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

func downloadRuleName(module *resolve.Module) string {
	return ruleName(module.Name, "_dl")
}

func toInstall(pkg *resolve.Package) string {
	install := strings.Trim(strings.TrimPrefix(pkg.ImportPath, pkg.Module), "/")
	if install == "" {
		return "."
	}
	return install
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
				modRule.SetAttr("download", NewStringExpr(":" + downloadRuleName(m)))
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