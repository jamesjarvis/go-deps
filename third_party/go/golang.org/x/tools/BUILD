go_mod_download(
    name = "my_fork",
    module = "github.com/tatskaari/tools",
    version = "v0.1.8-0.20211116155447-1211c0d87e32",
)

go_module(
    name = "tools",
    download = ":my_fork",
    install = [
        "go/gcexportdata",
        "go/internal/gcimporter",
        "go/internal/packagesdriver",
        "go/packages",
        "internal/event",
        "internal/event/core",
        "internal/event/keys",
        "internal/event/label",
        "internal/gocommand",
        "internal/packagesinternal",
        "internal/typeparams",
        "internal/typesinternal",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/tools",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go/golang.org/x/mod",
        "//third_party/go/golang.org/x/sys",
        "//third_party/go/golang.org/x/xerrors",
    ],
)
