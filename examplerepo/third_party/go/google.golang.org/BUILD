
go_module(
  name = "appengine",
  module = "google.golang.org/appengine",
  version = "v1.4.0",
  deps = [
    "//third_party/go/github.com/golang:protobuf",
    "//third_party/go/golang.org/x:net",
    "//third_party/go/golang.org/x:text",
  ],
  visibility = ["PUBLIC"],
  install = ["..."],
)

go_module(
  name = "genproto",
  module = "google.golang.org/genproto",
  version = "v0.0.0-20200513103714-09dca8ec2884",
  deps = [
    "//third_party/go/github.com/golang:protobuf",
    "//third_party/go/golang.org/x:lint",
    "//third_party/go/golang.org/x:tools",
    "//third_party/go/google.golang.org:grpc",
    "//third_party/go/honnef.co/go:tools",
  ],
  visibility = ["PUBLIC"],
  install = ["..."],
)

go_module(
  name = "grpc",
  module = "google.golang.org/grpc",
  version = "v1.33.1",
  deps = [
    "//third_party/go/github.com/cncf:udpa_go",
    "//third_party/go/github.com/envoyproxy:go-control-plane",
    "//third_party/go/github.com/golang:glog",
    "//third_party/go/github.com/golang:protobuf",
    "//third_party/go/github.com/google:go-cmp",
    "//third_party/go/github.com/google:uuid",
    "//third_party/go/golang.org/x:net",
    "//third_party/go/golang.org/x:oauth2",
    "//third_party/go/golang.org/x:sys",
    "//third_party/go/google.golang.org:genproto",
  ],
  visibility = ["PUBLIC"],
  install = ["..."],
)
