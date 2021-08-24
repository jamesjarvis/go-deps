package(default_visibility = ["PUBLIC"])

# To generate: plz run //resolve/cmd -- ./...

go_module(
    name = "golang.org.x.mod",
    module = "golang.org/x/mod",
    version = "v0.4.2",
    install = [
        "module",
        "modfile",
        "semver",
        "internal/lazyregexp",
    ],
    deps = [
        ":golang.org.x.xerrors",
    ],
)

go_module(
    name = "golang.org.x.sys",
    module = "golang.org/x/sys",
    version = "v0.0.0-20210603081109-ebe580a85c40",
    install = [
        "execabs",
    ],
)

go_module(
    name = "golang.org.x.tools",
    module = "golang.org/x/tools",
    version = "v0.1.2",
    install = [
        "internal/event/core",
        "go/gcexportdata",
        "go/packages",
        "internal/gocommand",
        "internal/packagesinternal",
        "internal/typesinternal",
        "go/internal/gcimporter",
        "go/internal/packagesdriver",
        "internal/event/label",
        "internal/event/keys",
        "internal/event",
    ],
    deps = [
        ":golang.org.x.xerrors",
        ":golang.org.x.sys",
        ":golang.org.x.mod",
    ],
)

go_module(
    name = "golang.org.x.xerrors",
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
    install = [
        ".",
        "internal",
    ],
)

go_module(
    name = "github.com.russross.blackfriday.v2",
    module = "github.com/russross/blackfriday/v2",
    version = "v2.1.0",
)

go_module(
    name = "github.com.cpuguy83.go-md2man.v2",
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.1",
    install = [
        "md2man",
    ],
    deps = [
        ":github.com.russross.blackfriday.v2",
    ],
)

go_module(
    name = "github.com.urfave.cli.v2",
    module = "github.com/urfave/cli/v2",
    version = "v2.3.0",
    deps = [
        ":github.com.cpuguy83.go-md2man.v2",
    ],
)

