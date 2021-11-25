genrule(
    name = "version",
    srcs = ["VERSION"],
    cmd = "echo VERSION = \\\"$(cat $SRCS)\\\" > $OUT",
    outs = ["version.build_defs"],
    visibility = ["PUBLIC"],
)

go_binary(
    name = "go-deps",
    srcs = ["main.go"],
    deps = [
        "//resolve",
        "//resolve/driver",
        "//rules",
        "//third_party/go/github.com/jessevdk/go-flags",
    ],
    visibility = ["PUBLIC"],
)
