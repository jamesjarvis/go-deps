go_library(
    name = "resolve",
    srcs = [
        "known_imports.go",
        "resolve.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//progress",
        "//resolve/model",
        "//third_party/go/github.com/google/go-licenses",
        "//third_party/go/golang.org/x/tools",
        "//third_party/go/golang.org/x/mod",
    ],
)

go_test(
    name = "resolve_test",
    srcs = ["resolve_test.go"],
    deps = [
        ":resolve",
        "//third_party/go/github.com/stretchr/testify:testify",
    ],
)