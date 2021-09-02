package model

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
	// The packages in this module
	Packages map[*Package]struct{}
	// The index of this module part
	Index int

	Modified bool
}