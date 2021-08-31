go_library(
    name = "resolve",
    srcs = [
        "resolve.go",
        "known_imports.go",
    ],
    deps = [
        "//host",
        "//rules",
        "//resolve/model",
        "//third_party/go:golang.org.x.tools",
        "//third_party/go:github.com.google.go-licenses",
    ],
    visibility = ["//resolve/..."],
)