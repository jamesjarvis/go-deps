# Go-Deps

This tool is used to help maintian your `go_module()` rules in a [Please](https://please.build) project.

# Features

Go-deps can be used to updates existing mdoules to newer version, or add new modules to your project. It 
works by parsing your existing BUILD files in your third party folder (by default `third_party/go/BUILD`, 
and updating them non-destructively. 

Go-deps has two modes of operation. It can generate a flat BUILD file, e.g. `third_party/go/BUILD`, or it
can split out each module into it's own build file e.g. `third_party/go/github.com/example/module/BUILD`.
That later can be very useful to improve maintainabilty in larger mono-repos, especially if you use `OWNER`
files to assign reviewers to branches of the source tree. 

# Installation

To install this in your project, add the following to your project:

```
GO_DEPS_VERSION = < version here, check https://github.com/Tatskaari/go-deps/releases >

remote_file(
    name = "go-deps",
    binary = True,
    url = f"https://github.com/Tatskaari/go-deps/releases/download/{GO_DEPS_VERSION}/go-deps",
)
```

# Usage
Note: go-deps works best with a `go.mod`.

First, install the module with `go get github.com/example/module/...`, then simply run `go-deps -w -m github.com/example/module/...`.
To add the `go_module()` rules into separate `BUILD` files for each module, pass the `--structured, -s` flag. 

```
NAME:
   please-go-get - Add a Go Module to an existing Please Monorepo

USAGE:
   go-deps [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --module value, -m value  Module to add
   --third_party value       The third party folder to write rules to (default: third_party/go)
   --write, -w               Whether to update the BUILD file(s), or just print to stdout (default: false)
   --structured, -s          Whether to put each module in a directory matching the module path, or write all module to a single file. (default: false)
   --help, -h                show help (default: false)
```

