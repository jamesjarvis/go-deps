package main

import (
	"os"

	"github.com/jamesjarvis/go-deps/resolve"
)



// This is janky and mostly just to test if this thing works.
func main() {
	rules, err := resolve.ResolveGet(os.Args[1:])
	if err != nil {
		panic(err)
	}


	for _, module := range rules {
		module.Print()
	}
}
