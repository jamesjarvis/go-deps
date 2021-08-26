package resolve

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDependsOn(t *testing.T) {
	r := newResolver("")

	p1, _ := r.getOrCreatePackage("m1/p1")
	p2, _ := r.getOrCreatePackage("m2/p2")
	p3, _ := r.getOrCreatePackage("m3/p3")
	p4, _ := r.getOrCreatePackage("m4/p4")
	p5, _ := r.getOrCreatePackage("m4/p5")

	// Package structure:
	// m1/p1 -> m2/p2 -> m3/p3 -> m4/p4
	// m4/p5 -> m1/p1

	p1.Module = "m1"
	p2.Module = "m2"
	p3.Module = "m3"
	p4.Module = "m4"
	p5.Module = "m4"

	p1.Imports = []*Package{p2}
	p2.Imports = []*Package{p3}
	p3.Imports = []*Package{p4}
	p5.Imports = []*Package{p1} // Causes a module cycle

	r.addPackageToModuleGraph(map[*Package]struct{}{}, p1)

	require.True(t, r.dependsOn(map[*Package]struct{}{}, p5, r.importPaths[p5.Imports[0]], false))

	r.addPackageToModuleGraph(map[*Package]struct{}{}, p5)
	_, ok := r.getModule("m4").Parts[1].Packages[p5]
	require.True(t, ok)

	r = newResolver("")

	r.addPackageToModuleGraph(map[*Package]struct{}{}, p3)

	r.addPackageToModuleGraph(map[*Package]struct{}{}, p1)
	r.addPackageToModuleGraph(map[*Package]struct{}{}, p5)
	require.True(t, ok)
}

func TestResolvesCycle(t *testing.T) {
	ps := map[string][]string{
		"google.golang.org/grpc/codes": {},
		"google.golang.org/grpc": {},
		"google.golang.org/grpc/status": {},
		"google.golang.org/grpc/metadata": {},
		"golang.org/x/oauth2": {},
		"cloud.google.com/go/compute/metadata": {},
		"golang.org/x/oauth2/google": {"cloud.google.com/go/compute/metadata"},
		"golang.org/x/oauth2/jwt": {},
		"google.golang.org/grpc/credentials/oauth": {"golang.org/x/oauth2", "golang.org/x/oauth2/google", "golang.org/x/oauth2/jwt"},
		"github.com/googleapis/gax-go/v2": {"google.golang.org/grpc/codes", "google.golang.org/grpc/status", "google.golang.org/grpc"},
		"cloud.google.com/go/talent/apiv4beta1": {"google.golang.org/grpc/codes", "github.com/googleapis/gax-go/v2", "google.golang.org/grpc", "google.golang.org/grpc/metadata"},
	}

	r := newResolver(".")

	getModuleNameFor := func(path string) string {
		modules := []string{"google.golang.org/grpc", "cloud.google.com/go", "golang.org/x/oauth2", "github.com/googleapis/gax-go/v2"}
		for _, m := range modules {
			if strings.HasPrefix(path, m) {
				return m
			}
		}
		t.Fatalf("can't determine module for %v", path)
		return ""
	}

	for importPath, imports := range ps {
		pkg, _ := r.getOrCreatePackage(importPath)
		pkg.Module = getModuleNameFor(importPath)
		for _, i := range imports {
			importedPackage, _ := r.getOrCreatePackage(i)
			pkg.Imports = append(pkg.Imports, importedPackage)
		}
	}

	r.addPackagesToModules()

	// Check we don't have a cycle
	module, ok := r.modules["cloud.google.com/go"]
	require.True(t, ok)

	for _, part := range module.Parts {
		deps := map[*ModulePart] struct{}{}
		findModuleDeps(r, part, part, deps)

		_, hasSelfDep := deps[part]
		require.False(t, hasSelfDep, "found dependency cycle")
	}
}

// findModuleDeps will return all the module parts (i.e. the go_module()) rules a module part depends on
func findModuleDeps(r *resolver, from *ModulePart, part *ModulePart, parts map[*ModulePart] struct{}) {
	for pkg := range part.Packages {
		for _, i := range pkg.Imports {
			mod := r.importPaths[i]
			if mod == from {
				continue
			}
			if _, ok := parts[mod]; !ok {
				parts[mod] = struct{}{}
				findModuleDeps(r, from, mod, parts)
			}
		}
	}
}
