# Example Repo

This repo contains the output of running:

```bash
go run ../main.go -m github.com/grpc-ecosystem/grpc-gateway -v v1.16.0
```

This was arbitrarily chosen as it's a kinda hard module to download (as there are many dependencies of dependencies and some are outdated).

As of 7/8/21 This does not generate an output that builds with `plz build //third_party/go/...` without manual intervention.
