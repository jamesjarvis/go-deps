package model

import (
	"path/filepath"
	"strings"
)

// Package represents a single package in some module
type Package struct {
	// The full import path of this package
	ImportPath string

	// The module name this package belongs to
	Module string

	// Any other packages this package imports
	Imports []*Package

	Resolved bool
}


// Module represents a module. It includes all deps so actually represents a full module graph.
type Module struct {
	// The module name
	Name string

	Version string
	Licence string

	Parts []*ModulePart
}

// ModulePart essentially corresponds to a `go_module()` rule that compiles some (or all) packages from that module. In
// most cases, there's one part per module except where we need to split it out to resolve a cycle.
type ModulePart struct {
	Module *Module

	// Any packages in the install list matched with "..." N.B the package doesn't have the /... on the end
	InstallWildCards []string

	// The packages in this module
	Packages map[*Package]struct{}
	// The index of this module part
	Index int

	Modified bool
}

func (p *ModulePart) IsWildcardImport(pkg *Package) bool {
	return p.GetWildcardImport(pkg) != ""
}

func (p *ModulePart) GetWildcardImport(pkg *Package) string {
	for _, i := range p.InstallWildCards {
		wildCardPath := filepath.Join(pkg.Module, i)
		if strings.HasPrefix(pkg.ImportPath, wildCardPath){
			return filepath.Join(i, "...")
		}
	}
	return ""
}