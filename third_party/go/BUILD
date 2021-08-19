package(default_visibility = ["PUBLIC"])

go_module(
    name = "mod",
    install = [
        "semver",
        "modfile",
        "internal/lazyregexp",
        "module",
    ],
    module = "golang.org/x/mod",
    version = "v0.5.0",
    deps = [":xerrors"],
)

go_module(
    name = "xerrors",
    install = [
        ".",
        "internal",
    ],
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_module(
    name = "cli.v2",
    module = "github.com/urfave/cli/v2",
    version = "v2.3.0",
    deps = [":md2man.v2"],
)

go_module(
    name = "md2man.v2",
    install = ["..."],
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.1",
    deps = [":blackfriday.v2"],
)

go_module(
    name = "blackfriday.v2",
    module = "github.com/russross/blackfriday/v2",
    version = "v2.1.0",
)
