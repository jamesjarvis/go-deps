package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    version = "1.16",
)

# To generate: plz run //resolve/cmd -- ./...
go_module(
    name = "golang.org.x.tools",
    install = [
        "internal/gocommand",
        "internal/event/keys",
        "go/gcexportdata",
        "internal/packagesinternal",
        "internal/event",
        "go/internal/packagesdriver",
        "internal/event/label",
        "internal/typesinternal",
        "go/internal/gcimporter",
        "internal/event/core",
        "go/packages",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/tools",
    version = "v0.0.0-20210106214847-113979e3529a",
    deps = [
        ":golang.org.x.mod",
        ":golang.org.x.xerrors",
    ],
)

go_module(
    name = "github.com.urfave.cli.v2",
    licences = ["MIT"],
    module = "github.com/urfave/cli/v2",
    version = "v2.3.0",
    deps = [":github.com.cpuguy83.go-md2man.v2"],
)

go_module(
    name = "github.com.russross.blackfriday.v2",
    licences = ["BSD-2-Clause"],
    module = "github.com/russross/blackfriday/v2",
    version = "v2.0.1",
    deps = [":github.com.shurcooL.sanitized_anchor_name"],
)

go_module(
    name = "github.com.google.licenseclassifier",
    install = [
        "stringclassifier/searchset/tokenizer",
        "stringclassifier",
        "stringclassifier/internal/pq",
        "internal/sets",
        "stringclassifier/searchset",
        ".",
        "licenses",
    ],
    licences = ["Apache-2.0"],
    module = "github.com/google/licenseclassifier",
    version = "v0.0.0-20210722185704-3043a050f148",
    deps = [":github.com.sergi.go-diff"],
)

go_module(
    name = "gopkg.in.src-d.go-git.v4",
    install = [
        "plumbing/format/pktline",
        "plumbing/format/gitignore",
        "utils/merkletrie/filesystem",
        "plumbing/cache",
        "plumbing/storer",
        "plumbing/transport/git",
        "utils/binary",
        "utils/merkletrie",
        "plumbing/protocol/packp/capability",
        "plumbing/transport",
        "plumbing/filemode",
        "plumbing/transport/client",
        "storage/filesystem/dotgit",
        "plumbing/transport/http",
        "plumbing",
        "plumbing/format/diff",
        "plumbing/format/index",
        "plumbing/object",
        "internal/revision",
        "utils/merkletrie/index",
        "utils/ioutil",
        "utils/diff",
        "utils/merkletrie/internal/frame",
        "plumbing/transport/server",
        "internal/url",
        "plumbing/protocol/packp/sideband",
        "config",
        "plumbing/protocol/packp",
        "utils/merkletrie/noder",
        ".",
        "plumbing/format/packfile",
        "plumbing/revlist",
        "storage",
        "plumbing/transport/ssh",
        "plumbing/format/objfile",
        "plumbing/format/idxfile",
        "plumbing/transport/internal/common",
        "storage/filesystem",
        "plumbing/format/config",
        "storage/memory",
        "plumbing/transport/file",
    ],
    licences = ["Apache-2.0"],
    module = "gopkg.in/src-d/go-git.v4",
    version = "v4.13.1",
    deps = [
        ":github.com.emirpasic.gods",
        ":github.com.jbenet.go-context",
        ":github.com.kevinburke.ssh_config",
        ":github.com.mitchellh.go-homedir",
        ":github.com.sergi.go-diff",
        ":github.com.src-d.gcfg",
        ":github.com.xanzy.ssh-agent",
        ":golang.org.x.crypto",
        ":golang.org.x.net",
        ":gopkg.in.src-d.go-billy.v4",
    ],
)

go_module(
    name = "github.com.xanzy.ssh-agent",
    licences = ["Apache-2.0"],
    module = "github.com/xanzy/ssh-agent",
    version = "v0.2.1",
    deps = [":golang.org.x.crypto"],
)

go_module(
    name = "gopkg.in.src-d.go-billy.v4",
    install = [
        "helper/chroot",
        "helper/polyfill",
        ".",
        "util",
        "osfs",
    ],
    licences = ["Apache-2.0"],
    module = "gopkg.in/src-d/go-billy.v4",
    version = "v4.3.2",
    deps = [":golang.org.x.sys"],
)

go_module(
    name = "github.com.shurcooL.sanitized_anchor_name",
    licences = ["MIT"],
    module = "github.com/shurcooL/sanitized_anchor_name",
    version = "v1.0.0",
)

go_module(
    name = "gopkg.in.warnings.v0",
    licences = ["BSD-2-Clause"],
    module = "gopkg.in/warnings.v0",
    version = "v0.1.2",
)

go_module(
    name = "golang.org.x.xerrors",
    install = [
        "internal",
        ".",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_module(
    name = "github.com.cpuguy83.go-md2man.v2",
    install = ["md2man"],
    licences = ["MIT"],
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.0",
    deps = [":github.com.russross.blackfriday.v2"],
)

go_module(
    name = "github.com.mitchellh.go-homedir",
    licences = ["MIT"],
    module = "github.com/mitchellh/go-homedir",
    version = "v1.1.0",
)

go_module(
    name = "github.com.kevinburke.ssh_config",
    licences = ["MIT"],
    module = "github.com/kevinburke/ssh_config",
    version = "v0.0.0-20190725054713-01f96b0aa0cd",
)

go_module(
    name = "golang.org.x.crypto",
    install = [
        "ssh/knownhosts",
        "openpgp",
        "blowfish",
        "openpgp/armor",
        "poly1305",
        "cast5",
        "openpgp/elgamal",
        "ssh",
        "curve25519",
        "ed25519",
        "openpgp/errors",
        "openpgp/packet",
        "ssh/agent",
        "internal/subtle",
        "openpgp/s2k",
        "ssh/internal/bcrypt_pbkdf",
        "chacha20",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20200622213623-75b288015ac9",
)

go_module(
    name = "github.com.emirpasic.gods",
    install = [
        "utils",
        "lists",
        "trees",
        "lists/arraylist",
        "containers",
        "trees/binaryheap",
    ],
    licences = ["BSD-2-Clause"],
    module = "github.com/emirpasic/gods",
    version = "v1.12.0",
)

go_module(
    name = "github.com.jbenet.go-context",
    install = ["io"],
    licences = ["MIT"],
    module = "github.com/jbenet/go-context",
    version = "v0.0.0-20150711004518-d14ea06fba99",
    deps = [":golang.org.x.net"],
)

go_module(
    name = "golang.org.x.sys",
    install = [
        "unix",
        "internal/unsafeheader",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/sys",
    version = "v0.0.0-20200930185726-fdedc70b468f",
)

go_module(
    name = "github.com.sergi.go-diff",
    install = ["diffmatchpatch"],
    licences = ["MIT"],
    module = "github.com/sergi/go-diff",
    version = "v1.2.0",
)

go_module(
    name = "golang.org.x.mod",
    install = [
        "modfile",
        "internal/lazyregexp",
        "module",
        "semver",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/mod",
    version = "v0.4.0",
    deps = [":golang.org.x.xerrors"],
)

go_module(
    name = "golang.org.x.net",
    install = [
        "proxy",
        "context",
        "internal/socks",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/net",
    version = "v0.0.0-20201021035429-f5854403a974",
)

go_module(
    name = "github.com.golang.glog",
    licences = ["Apache-2.0"],
    module = "github.com/golang/glog",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_module(
    name = "github.com.google.go-licenses",
    install = ["licenses"],
    licences = ["Apache-2.0"],
    module = "github.com/google/go-licenses",
    version = "v0.0.0-20210816172045-3099c18c36e1",
    deps = [
        ":github.com.golang.glog",
        ":github.com.google.licenseclassifier",
        ":golang.org.x.tools",
        ":gopkg.in.src-d.go-git.v4",
    ],
)

go_module(
    name = "github.com.src-d.gcfg",
    install = [
        "types",
        ".",
        "token",
        "scanner",
    ],
    licences = ["BSD-3-Clause"],
    module = "github.com/src-d/gcfg",
    version = "v1.4.0",
    deps = [":gopkg.in.warnings.v0"],
)

go_module(
    name = "github.com.bazelbuild.buildtools",
    install = [
        "build_proto",
        "wspace",
        "edit",
        "api_proto",
        "labels",
        "build",
        "lang",
        "tables",
        "file",
    ],
    licences = ["Apache-2.0"],
    module = "github.com/bazelbuild/buildtools",
    version = "v0.0.0-20210824143317-0a897f9af678",
    deps = [
        ":github.com.golang.protobuf",
        ":google.golang.org.protobuf",
    ],
)

go_module(
    name = "github.com.golang.protobuf",
    install = ["proto"],
    licences = ["BSD-3-Clause"],
    module = "github.com/golang/protobuf",
    version = "v1.4.3",
    deps = [":google.golang.org.protobuf"],
)

go_module(
    name = "google.golang.org.protobuf",
    install = [
        "internal/encoding/messageset",
        "internal/detrand",
        "internal/mapsort",
        "runtime/protoimpl",
        "internal/genid",
        "internal/fieldsort",
        "proto",
        "internal/descfmt",
        "reflect/protoregistry",
        "encoding/protowire",
        "encoding/prototext",
        "internal/encoding/tag",
        "internal/flags",
        "internal/set",
        "reflect/protoreflect",
        "internal/errors",
        "internal/filedesc",
        "internal/strs",
        "internal/filetype",
        "internal/impl",
        "internal/version",
        "internal/descopts",
        "internal/pragma",
        "internal/encoding/text",
        "runtime/protoiface",
        "internal/encoding/defval",
    ],
    licences = ["BSD-3-Clause"],
    module = "google.golang.org/protobuf",
    version = "v1.25.0",
)
