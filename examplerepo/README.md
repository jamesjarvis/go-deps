# Example Repo

This repo contains the output of running:

```bash
go run ../main.go -m github.com/grpc-ecosystem/grpc-gateway -v v1.16.0
```

This was arbitrarily chosen as it's a kinda hard module to download (as there are many dependencies of dependencies and some are outdated).

As of 7/8/21 This does not generate an output that builds with `plz build //third_party/go/...` without manual intervention.

Example Tests:

- ✅ `go run ../main.go -m github.com/stretchr/testify -v v1.6.1`
- ❌ `go run ../main.go -m github.com/grpc-ecosystem/grpc-gateway -v v1.16.0`
  - This is because some dependencies do not actually contain go.mod files at the requested version, so their builds fail.
