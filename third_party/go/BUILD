package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    version = "1.16",
)

# To generate: plz run //:go-deps -- --module ./...

go_module(
    name = "xerrors",
    install = [
        "internal",
        ".",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_module(
    name = "licenseclassifier",
    install = [
        ".",
        "stringclassifier/searchset",
        "stringclassifier/internal/pq",
        "licenses",
        "stringclassifier/searchset/tokenizer",
        "internal/sets",
        "stringclassifier",
    ],
    licences = ["Apache-2.0"],
    module = "github.com/google/licenseclassifier",
    version = "v0.0.0-20210722185704-3043a050f148",
    deps = [":go-diff"],
)

go_module(
    name = "sys",
    install = [
        "unix",
        "internal/unsafeheader",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/sys",
    version = "v0.0.0-20210514084401-e8d321eab015",
)

go_module(
    name = "sanitized_anchor_name",
    licences = ["MIT"],
    module = "github.com/shurcooL/sanitized_anchor_name",
    version = "v1.0.0",
)

go_module(
    name = "go-git.v4",
    install = [
        "plumbing/object",
        "utils/binary",
        "plumbing/transport/client",
        "plumbing/format/packfile",
        "plumbing",
        "plumbing/format/objfile",
        "plumbing/format/pktline",
        "storage/memory",
        "utils/merkletrie/noder",
        "plumbing/transport/server",
        "plumbing/format/config",
        "plumbing/cache",
        "utils/ioutil",
        "storage/filesystem/dotgit",
        "storage",
        "plumbing/format/index",
        "plumbing/format/diff",
        "plumbing/protocol/packp",
        "utils/merkletrie/internal/frame",
        "plumbing/revlist",
        "plumbing/filemode",
        "plumbing/transport/http",
        "plumbing/protocol/packp/sideband",
        "plumbing/transport",
        "plumbing/transport/file",
        "plumbing/transport/ssh",
        "storage/filesystem",
        "utils/merkletrie",
        "utils/diff",
        "plumbing/format/idxfile",
        "plumbing/format/gitignore",
        "plumbing/transport/git",
        "utils/merkletrie/index",
        "plumbing/protocol/packp/capability",
        "plumbing/transport/internal/common",
        "utils/merkletrie/filesystem",
        "internal/revision",
        "internal/url",
        ".",
        "plumbing/storer",
        "config",
    ],
    licences = ["Apache-2.0"],
    module = "gopkg.in/src-d/go-git.v4",
    version = "v4.13.1",
    deps = [
        ":crypto",
        ":gcfg",
        ":go-billy.v4",
        ":go-context",
        ":go-diff",
        ":go-homedir",
        ":gods",
        ":net",
        ":ssh-agent",
        ":ssh_config",
    ],
)

go_module(
    name = "go-context",
    install = ["io"],
    licences = ["MIT"],
    module = "github.com/jbenet/go-context",
    version = "v0.0.0-20150711004518-d14ea06fba99",
    deps = [":net"],
)

go_module(
    name = "glog",
    licences = ["Apache-2.0"],
    module = "github.com/golang/glog",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_module(
    name = "crypto",
    install = [
        "ssh",
        "poly1305",
        "blowfish",
        "chacha20",
        "openpgp/packet",
        "openpgp/armor",
        "openpgp/elgamal",
        "internal/subtle",
        "ssh/internal/bcrypt_pbkdf",
        "openpgp",
        "ssh/agent",
        "ssh/knownhosts",
        "cast5",
        "ed25519",
        "curve25519",
        "openpgp/s2k",
        "openpgp/errors",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20200622213623-75b288015ac9",
)

go_module(
    name = "go-billy.v4",
    install = [
        "osfs",
        ".",
        "util",
        "helper/polyfill",
        "helper/chroot",
    ],
    licences = ["Apache-2.0"],
    module = "gopkg.in/src-d/go-billy.v4",
    version = "v4.3.2",
    deps = [":sys"],
)

go_module(
    name = "go-md2man.v2",
    install = ["md2man"],
    licences = ["MIT"],
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.0",
    deps = [":blackfriday.v2"],
)

go_module(
    name = "warnings.v0",
    licences = ["BSD-2-Clause"],
    module = "gopkg.in/warnings.v0",
    version = "v0.1.2",
)

go_module(
    name = "go-licenses",
    install = ["licenses"],
    licences = ["Apache-2.0"],
    module = "github.com/google/go-licenses",
    version = "v0.0.0-20210816172045-3099c18c36e1",
    deps = [
        ":glog",
        ":go-git.v4",
        ":licenseclassifier",
        ":tools",
    ],
)

go_module(
    name = "ssh-agent",
    licences = ["Apache-2.0"],
    module = "github.com/xanzy/ssh-agent",
    version = "v0.2.1",
    deps = [":crypto"],
)

go_module(
    name = "buildtools",
    install = [
        "file",
        "lang",
        "build",
        "edit",
        "build_proto",
        "api_proto",
        "labels",
        "tables",
        "wspace",
    ],
    licences = ["Apache-2.0"],
    module = "github.com/bazelbuild/buildtools",
    version = "v0.0.0-20210824143317-0a897f9af678",
    deps = [
        ":google.golang.org.protobuf",
        ":protobuf",
    ],
)

go_module(
    name = "google.golang.org.protobuf",
    install = [
        "proto",
        "internal/flags",
        "internal/mapsort",
        "internal/set",
        "reflect/protoregistry",
        "runtime/protoimpl",
        "reflect/protoreflect",
        "encoding/prototext",
        "internal/pragma",
        "internal/errors",
        "internal/detrand",
        "internal/descopts",
        "internal/genid",
        "internal/strs",
        "internal/encoding/tag",
        "internal/filedesc",
        "internal/encoding/messageset",
        "encoding/protowire",
        "internal/impl",
        "internal/version",
        "internal/encoding/defval",
        "internal/descfmt",
        "runtime/protoiface",
        "internal/filetype",
        "internal/fieldsort",
        "internal/encoding/text",
    ],
    licences = ["BSD-3-Clause"],
    module = "google.golang.org/protobuf",
    version = "v1.25.0",
)

go_module(
    name = "protobuf",
    install = ["proto"],
    licences = ["BSD-3-Clause"],
    module = "github.com/golang/protobuf",
    version = "v1.4.3",
    deps = [":google.golang.org.protobuf"],
)

go_module(
    name = "cli.v2",
    licences = ["MIT"],
    module = "github.com/urfave/cli/v2",
    version = "v2.3.0",
    deps = [":go-md2man.v2"],
)

go_module(
    name = "go-diff",
    install = ["diffmatchpatch"],
    licences = ["MIT"],
    module = "github.com/sergi/go-diff",
    version = "v1.2.0",
)

go_module(
    name = "mod",
    install = ["semver"],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/mod",
    version = "v0.4.0",
)

go_module(
    name = "gcfg",
    install = [
        ".",
        "types",
        "scanner",
        "token",
    ],
    licences = ["BSD-3-Clause"],
    module = "github.com/src-d/gcfg",
    version = "v1.4.0",
    deps = [":warnings.v0"],
)

go_module(
    name = "net",
    install = [
        "internal/socks",
        "context",
        "proxy",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/net",
    version = "v0.0.0-20201021035429-f5854403a974",
)

go_module(
    name = "go-homedir",
    licences = ["MIT"],
    module = "github.com/mitchellh/go-homedir",
    version = "v1.1.0",
)

go_module(
    name = "ssh_config",
    licences = ["MIT"],
    module = "github.com/kevinburke/ssh_config",
    version = "v0.0.0-20190725054713-01f96b0aa0cd",
)

go_module(
    name = "tools",
    install = [
        "internal/event",
        "internal/typesinternal",
        "go/packages",
        "go/internal/packagesdriver",
        "internal/event/label",
        "internal/event/keys",
        "go/gcexportdata",
        "internal/packagesinternal",
        "internal/event/core",
        "internal/gocommand",
        "go/internal/gcimporter",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/tools",
    version = "v0.0.0-20210106214847-113979e3529a",
    deps = [
        ":mod",
        ":xerrors",
    ],
)

go_module(
    name = "gods",
    install = [
        "lists/arraylist",
        "containers",
        "utils",
        "trees/binaryheap",
        "lists",
        "trees",
    ],
    licences = ["BSD-2-Clause"],
    module = "github.com/emirpasic/gods",
    version = "v1.12.0",
)

go_module(
    name = "blackfriday.v2",
    licences = ["BSD-2-Clause"],
    module = "github.com/russross/blackfriday/v2",
    version = "v2.0.1",
    deps = [":sanitized_anchor_name"],
)
