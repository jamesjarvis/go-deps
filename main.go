package main

import (
	"context"
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
		},
		Action: func(c *cli.Context) error {
			ctx := context.TODO()
			fmt.Println("Please Go Get v0.0.1")

			alreadyExists, err := host.CreateGoMod(ctx)
			if err != nil {
				return err
			}
			if !alreadyExists {
				defer host.TearDownGoMod(ctx)
			}

			m := &module.Module{
				Path:    c.String(moduleFlag),
				Version: c.String(versionFlag),
			}

			fmt.Printf("So, you want to add %q?\n", m.String())

			err = m.Download(ctx)
			if err != nil {
				return err
			}

			_, err = m.GetDependenciesRecursively(ctx)
			if err != nil {
				return err
			}

			module.GlobalCache.Sync()
			module.GlobalCache.Print()

			return module.GlobalCache.ExportBuildRules()
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
