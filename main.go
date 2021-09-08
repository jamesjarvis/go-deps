package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jamesjarvis/go-deps/host"
	"github.com/jamesjarvis/go-deps/module"
	"github.com/urfave/cli/v2"
)

const (
	moduleFlag  = "module"
	versionFlag = "version"
	thirdPartyFlag = "third_party"
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
				Name:    versionFlag,
				Aliases: []string{"v"},
				Usage:   "Version of the module to add",
			},
			&cli.StringFlag{
				Name:    thirdPartyFlag,
				DefaultText: "third_party/go",
				Usage:   "The third party folder to write rules to",
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println("Please Go Get v0.0.1")

			alreadyExists, err := host.CreateGoMod(ctx.Context)
			if err != nil {
				return err
			}
			if !alreadyExists {
				defer host.TearDownGoMod(ctx.Context)
			}

			m := &module.Module{
				Path:    ctx.String(moduleFlag),
				Version: ctx.String(versionFlag),
			}

			fmt.Printf("So, you want to add %q?\n", m.String())

			err = m.Download(ctx.Context)
			if err != nil {
				return err
			}

			_, err = m.GetDependenciesRecursively(ctx.Context)
			if err != nil {
				return err
			}

			module.GlobalCache.Sync()
			module.GlobalCache.Print()

			return module.GlobalCache.ExportBuildRules(ctx.String(thirdPartyFlag))
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
