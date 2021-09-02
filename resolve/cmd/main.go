package main

import (
	"os"

	"github.com/jamesjarvis/go-deps/resolve"
	"github.com/jamesjarvis/go-deps/rules"
)



// This is janky and mostly just to test if this thing works.
func main() {
	moduleGraph, err := rules.ReadRules(os.Args[1])
	if err != nil {
		panic(err)
	}
	err = resolve.UpdateModules(moduleGraph.Modules, os.Args[2:])
	if err != nil {
		panic(err)
	}

	err = moduleGraph.Save()
	if err != nil {
		panic(err)
	}
}
