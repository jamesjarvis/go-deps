package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/tatskaari/go-deps/resolve"
	"github.com/tatskaari/go-deps/rules"

	"github.com/urfave/cli/v2"
)

const (
	moduleFlag     = "module"
	versionFlag    = "version"
	thirdPartyFlag = "third_party"
	writeFlag      = "write"
	structuredFlag = "structured"
)

// This binary will accept a module name and optionally a semver or commit hash, and will add this module to a BUILD file.
func main() {
	app := &cli.App{
		Name:  "please-go-get",
		Usage: "Add a Go Module to an existing Please Monorepo",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     moduleFlag,
				Aliases:  []string{"m"},
				Usage:    "Module to add",
				Required: true,
			},
			&cli.StringFlag{
				Name:    thirdPartyFlag,
				DefaultText: "third_party/go",
				Usage:   "The third party folder to write rules to",
			},
			&cli.BoolFlag{
				Name:    writeFlag,
				Aliases: []string{"w"},
				Usage:   "Whether to update the BUILD file(s), or just print to stdout",
			},
			&cli.BoolFlag{
				Name:    structuredFlag,
				Aliases: []string{"s"},
				Usage:   "Whether to put each module in a directory matching the module path, or write all module to a single file.",
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println("Please Go Get v0.0.1")

			thirdPartyFolder := ctx.String(thirdPartyFlag)
			if thirdPartyFolder == "" {
				thirdPartyFolder = "third_party/go"
			}

			moduleGraph := rules.NewGraph()
			if ctx.Bool(structuredFlag) {
				err := filepath.Walk(thirdPartyFolder, func(path string, info fs.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}
					if filepath.Base(path) == "BUILD" {
						if err := moduleGraph.ReadRules(path); err != nil {
							return err
						}
					}
					return nil
				})
				if err != nil {
					return err
				}
			} else {
				if err := moduleGraph.ReadRules(filepath.Join(thirdPartyFolder, "BUILD")); err != nil {
					return err
				}
			}

			err := resolve.UpdateModules(moduleGraph.Modules, []string{ctx.String(moduleFlag)})
			if err != nil {
				return err
			}

			return moduleGraph.Save(ctx.Bool(structuredFlag), ctx.Bool(writeFlag), thirdPartyFolder)
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
