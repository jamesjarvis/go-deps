
go_module(
  name = "go-control-plane",
  module = "github.com/envoyproxy/go-control-plane",
  version = "v0.9.4",
  deps = [
    "//third_party/go/github.com/census-instrumentation:opencensus-proto",
    "//third_party/go/github.com/cncf:udpa_go",
    "//third_party/go/github.com/envoyproxy:protoc-gen-validate",
    "//third_party/go/github.com/golang:protobuf",
    "//third_party/go/github.com/google:go-cmp",
    "//third_party/go/github.com/prometheus:client_model",
    "//third_party/go/google.golang.org:genproto",
    "//third_party/go/google.golang.org:grpc",
  ],
  visibility = ["PUBLIC"],
  install = ["..."],
)

go_module(
  name = "protoc-gen-validate",
  module = "github.com/envoyproxy/protoc-gen-validate",
  version = "v0.1.0",
  deps = [
  ],
  visibility = ["PUBLIC"],
  install = ["..."],
)
