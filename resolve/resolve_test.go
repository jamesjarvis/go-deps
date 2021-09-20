package resolve

import (
	. "github.com/jamesjarvis/go-deps/resolve/model"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDependsOn(t *testing.T) {
	r := newResolver(".", nil)

	// Package structure:
	// m1/p1 --> m2/p2 --> m3/p3 --> m4/p4
	// m1/p1   	 <--------------	 m4/p5

	m1p1 := r.GetPackage("m1/p1")
	m2p2 := r.GetPackage("m2/p2")
	m3p3 := r.GetPackage("m3/p3")
	m4p4 := r.GetPackage("m4/p4")
	m4p5 := r.GetPackage("m4/p5")

	m1p1.Module = "m1"
	m2p2.Module = "m2"
	m3p3.Module = "m3"
	m4p4.Module = "m4"
	m4p5.Module = "m4"

	m1p1.Imports = []*Package{m2p2}
	m2p2.Imports = []*Package{m3p3}
	m3p3.Imports = []*Package{m4p4}
	m4p5.Imports = []*Package{m1p1} // Causes a module cycle

	// Add the packages to the graph
	r.addPackageToModuleGraph(map[*Package]struct{}{}, m1p1)
	r.addPackageToModuleGraph(map[*Package]struct{}{}, m4p5)

	// Check that m4/p5 has an import that depends on m4/p4 (creating a module cycle)
	require.True(t, r.dependsOn(map[*Package]struct{}{}, m4p5.Imports[0], r.ImportPaths[m4p4]))

	// Check that we resolved that by creating a new part
	require.Len(t, r.GetModule("m4").Parts, 2)
	_, ok := r.GetModule("m4").Parts[1].Packages[m4p5]
	require.True(t, ok)
}

func TestResolvesCycle(t *testing.T) {
	// This package structure is a simplified form of the google.golang.com/go module
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

	r := newResolver(".", nil)

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
		pkg := r.GetPackage(importPath)
		pkg.Module = getModuleNameFor(importPath)
		for _, i := range imports {
			importedPackage := r.GetPackage(i)
			pkg.Imports = append(pkg.Imports, importedPackage)
		}
	}

	r.addPackagesToModules(map[*Package]struct{}{})

	// Check we don't have a cycle
	module, ok := r.Mods["cloud.google.com/go"]
	require.True(t, ok)

	// TODO(jpoole): Make the generated module graph deterministic so we don't have to have a complicated assertion here
	for _, part := range module.Parts {
		deps := map[*ModulePart] struct{}{}
		findModuleDeps(r, part, part, deps)

		_, hasSelfDep := deps[part]
		require.False(t, hasSelfDep, "found dependency cycle")
	}
}

// findModuleDeps will return all the module parts (i.e. the go_module()) rules a module part depends on
func findModuleDeps(r *resolver, from *ModulePart, currentPart *ModulePart, parts map[*ModulePart] struct{}) {
	for pkg := range currentPart.Packages {
		for _, i := range pkg.Imports {
			mod := r.ImportPaths[i]
			// Ignore self imports
			if mod == currentPart {
				continue
			}
			// We found a cycle, return so we don't stack overflow
			if mod == from {
				parts[mod] = struct{}{}
				return
			}
			if _, ok := parts[mod]; !ok {
				parts[mod] = struct{}{}
				findModuleDeps(r, from, mod, parts)
			}
		}
	}
}
