# Go-Deps

The goal of this project is to build a migration tool to convert a [Please](https://github.com/thought-machine/please)
based repo using deprecated `go_get` rules to use the now standard `go_module` rules.

The difference between please and bazel is that please requires the user to specify the module dependencies,
which can be a little bit annoying.

## Rough idea

The rough idea of this project is to have two modes:

1. Please just add this module to the project
   - Basically just run a binary, passing in the specified module + version + optionally the name for it
   - This will then parse the existing build files to check if it already exists, and then add this module + it's
     dependencies to the repo.
   - The idea is that in order to find out the dependencies of a particular module, it will run `go get` (or `go mod download` idk)
     and then cd into the cache directory and run `go list -m -json all` on that stuff, to work out what the dependencies are.
2. Please just convert the existing go_get definitions into go_modules (should theoretically just call the above mentioned binary).

## Current Progress

Currently, we are at step 1.5, ie: we can pass in a module + optionally version and it will resolve the dependencies
for it and it's dependencies:

```bash
itis@whatitis go-deps % go run cmd/main.go -m github.com/hashicorp/go-hclog
Please Go Get v0.0.1
So, you want to add "github.com/hashicorp/go-hclog"?
Congrats, you just downloaded "github.com/hashicorp/go-hclog@v0.16.2"
2021/08/04 23:22:56 Dependencies change! We started with github.com/mattn/go-isatty@v0.0.8 and now have github.com/mattn/go-isatty@v0.0.10
2021/08/04 23:22:57 Dependencies change! We started with golang.org/x/sys@v0.0.0-20190222072716-a9d3bda3a223 and now have golang.org/x/sys@v0.0.0-20191008105621-543471e840be
2021/08/04 23:22:57 Synced github.com/mattn/go-isatty@v0.0.8 --> github.com/mattn/go-isatty@v0.0.10
MODULE: github.com/hashicorp/go-hclog
 VERSION: v0.16.2
  github.com/hashicorp/go-hclog@v0.16.2
  |
  |---- github.com/fatih/color@v1.7.0
  |---- github.com/mattn/go-colorable@v0.1.4
  |---- github.com/mattn/go-isatty@v0.0.10
  |---- github.com/stretchr/testify@v1.2.2
MODULE: github.com/fatih/color
 VERSION: v1.7.0
  github.com/fatih/color@v1.7.0
MODULE: github.com/mattn/go-colorable
 VERSION: v0.1.4
  github.com/mattn/go-colorable@v0.1.4
  |
  |---- github.com/mattn/go-isatty@v0.0.10
MODULE: github.com/mattn/go-isatty
 VERSION: v0.0.10
  github.com/mattn/go-isatty@v0.0.10
  |
  |---- golang.org/x/sys@v0.0.0-20191008105621-543471e840be
MODULE: github.com/stretchr/testify
 VERSION: v1.2.2
  github.com/stretchr/testify@v1.2.2
MODULE: golang.org/x/sys
 VERSION: v0.0.0-20191008105621-543471e840be
  golang.org/x/sys@v0.0.0-20191008105621-543471e840be
```
