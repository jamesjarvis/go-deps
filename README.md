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
