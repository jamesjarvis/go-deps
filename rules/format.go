package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	resolve "github.com/tatskaari/go-deps/resolve/model"

	"github.com/bazelbuild/buildtools/build"
	"github.com/bazelbuild/buildtools/edit"
)

var semverRegex = regexp.MustCompile("^v[0-9]+$")

func split(path string) (string, string) {
	dir, base := filepath.Split(path)
	return filepath.Clean(dir), base
}

func (file *BuildFile) assignName(originalPath, suffix string, structured bool) string {
	path, base := split(originalPath)
	name := base + suffix
	if !structured && semverRegex.MatchString(name) {
		path, base = split(path)
		name = base + "." + name
	}

	for {
		extantPath, ok := file.usedNames[name]
		if !ok {
			break
		}
		if extantPath == originalPath {
			return name
		}
		path, base = split(path)
		name = base + "." + name
	}
	file.usedNames[name] = originalPath
	return name
}

func (file *BuildFile) partName(part *resolve.ModulePart, structured bool) string {
	if name, ok := file.partNames[part]; ok {
		return name
	}
	displayIndex := len(part.Module.Parts) - part.Index
	suffix := ""
	if displayIndex > 0 {
		suffix = fmt.Sprintf("_%d", displayIndex)
	}

	name := file.assignName(part.Module.Name, suffix, structured)
	file.partNames[part] = name
	return name
}

func (file *BuildFile) downloadRuleName(module *resolve.Module,  structured bool) string {
	if name, ok := file.downloadNames[module]; ok {
		return name
	}

	name := file.assignName(module.Name, "_dl", structured)
	file.downloadNames[module] = name
	return name
}

func toInstall(part *resolve.ModulePart, pkg *resolve.Package) string {
	if wildCard := part.GetWildcardImport(pkg); wildCard != "" {
		return wildCard
	}
	install := strings.Trim(strings.TrimPrefix(pkg.ImportPath, pkg.Module), "/")
	if install == "" {
		return "."
	}
	return install
}

func (g *BuildGraph) file(mod *resolve.Module, structured bool, thirdPartyFolder string) (*BuildFile, error) {
	path := ""
	if structured {
		path = filepath.Join(thirdPartyFolder, mod.Name, "BUILD")
	} else {
		path = filepath.Join(thirdPartyFolder, "BUILD")
	}

	if f, ok := g.Files[path]; ok {
		g.ModFiles[mod] = f
		return f, nil
	} else {
		// TODO create the build file
		file, err := newFile(path)
		if err != nil {
			return nil, err
		}

		g.ModFiles[mod] = file
		g.Files[path] = file

		return file, nil
	}
}


func cannonicalise(name, modpath, thirdParty string, structured bool) string {
	if !structured {
		return ":" + name
	}
	return "//" + filepath.Join(thirdParty, modpath)
}

func (g *BuildGraph) Save(structured, write bool, thirdPartyFolder string) error {
	for _, m := range g.Modules.Mods {
		file, err := g.file(m, structured, thirdPartyFolder)
		if err != nil {
			return err
		}
		dlRule, ok := file.ModDownloadRules[m]
		if len(m.Parts) > 1 {
			if !ok {
				dlRule = NewRule(file.File, "go_mod_download", file.downloadRuleName(m, structured))
				file.ModDownloadRules[m] = dlRule
			}
			dlRule.SetAttr("module", NewStringExpr(m.Name))
			if m.Version != "" {
				dlRule.SetAttr("version", NewStringExpr(m.Version))
			}
			if m.Licence != "" {
				dlRule.SetAttr("licences", NewStringList(m.Licence))
			}
		}

		for _, part := range m.Parts {
			if !part.Modified {
				continue
			}
			modRule, ok := file.ModRules[part]
			if !ok {
				modRule = NewRule(file.File, "go_module", file.partName(part, structured))
				file.ModRules[part] = modRule
			}
			modRule.DelAttr("install")
			modRule.DelAttr("deps")
			modRule.DelAttr("exported_deps")
			modRule.DelAttr("visibility")

			modRule.SetAttr("module", NewStringExpr(m.Name))

			if len(m.Parts) > 1 {
				modRule.DelAttr("version")
				modRule.SetAttr("download", NewStringExpr(":" + file.downloadRuleName(m, structured)))
			} else {
				if m.Licence != "" {
					modRule.SetAttr("licences", NewStringList(m.Licence))
				}
				if m.Version != "" {
					modRule.SetAttr("version", NewStringExpr(m.Version))
				}
			}

			installs := make([]string, 0, len(part.Packages))
			deps := make([]string, 0, len(part.Packages))
			var exportedDeps []string

			doneDeps := map[string]struct{}{}
			doneInstalls := map[string]struct{}{}

			for pkg := range part.Packages {
				i := toInstall(part, pkg)
				if _, ok := doneInstalls[i]; !ok {
					installs = append(installs, i)
					doneInstalls[i] = struct{}{}
				}

				for _, i := range pkg.Imports {
					dep := g.Modules.Import(i)
					depRuleName := file.partName(dep, structured)
					if _, ok := doneDeps[depRuleName]; ok || dep.Module == m {
						continue
					}
					doneDeps[depRuleName] = struct{}{}
					deps = append(deps, cannonicalise(depRuleName, dep.Module.Name, thirdPartyFolder, structured))
				}
			}

			// The last part is the namesake and should export the rest of the parts.
			if part.Index == len(m.Parts) {
				modRule.SetAttr("visibility", NewStringList("PUBLIC"))

				for _, part := range m.Parts[:(len(m.Parts) - 1)] {
					exportedDeps = append(exportedDeps, ":" + file.partName(part, structured))
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

	for path, f := range g.Files {
		if write {
			if err := os.MkdirAll(filepath.Dir(f.File.Path), os.ModeDir | 0775); err != nil {
				return err
			}

			osFile, err := os.Create(f.File.Path)
			if err != nil {
				return err
			}

			if _, err := osFile.Write(build.Format(f.File)); err != nil {
				return err
			}
			osFile.Close()
		} else {
			fmt.Println("# " + path)
			fmt.Println(string(build.Format(f.File)))
		}

	}
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