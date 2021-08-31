package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    version = "1.16",
)


# To generate: plz run //resolve/cmd -- ./...
go_module(
    name = "golang.org.x.tools",
    module = "golang.org/x/tools",
    version = "v0.0.0-20210106214847-113979e3529a",
    install = [
        "internal/typesinternal",
        "internal/event/keys",
        "go/gcexportdata",
        "internal/packagesinternal",
        "go/internal/gcimporter",
        "internal/event/core",
        "internal/event/label",
        "internal/event",
        "go/internal/packagesdriver",
        "internal/gocommand",
        "go/packages",
    ],
    deps = [
        ":golang.org.x.mod",
        ":golang.org.x.xerrors",
    ],
    licences = ["BSD-3-Clause"]
)

go_module(
    name = "github.com.urfave.cli.v2",
    module = "github.com/urfave/cli/v2",
    version = "v2.3.0",
    deps = [
        ":github.com.cpuguy83.go-md2man.v2",
    ],
    licences = ["MIT"]
)

go_module(
    name = "github.com.russross.blackfriday.v2",
    module = "github.com/russross/blackfriday/v2",
    version = "v2.0.1",
    deps = [
        ":github.com.shurcooL.sanitized_anchor_name",
    ],
    licences = ["BSD-2-Clause"]
)

go_module(
    name = "github.com.google.licenseclassifier",
    module = "github.com/google/licenseclassifier",
    version = "v0.0.0-20210722185704-3043a050f148",
    install = [
        ".",
        "licenses",
        "stringclassifier/searchset/tokenizer",
        "stringclassifier",
        "stringclassifier/internal/pq",
        "internal/sets",
        "stringclassifier/searchset",
    ],
    deps = [
        ":github.com.sergi.go-diff",
    ],
    licences = ["Apache-2.0"]
)

go_module(
    name = "gopkg.in.src-d.go-git.v4",
    module = "gopkg.in/src-d/go-git.v4",
    version = "v4.13.1",
    install = [
        ".",
        "plumbing/format/pktline",
        "plumbing/storer",
        "plumbing/format/index",
        "storage/filesystem",
        "plumbing/format/packfile",
        "plumbing/transport/internal/common",
        "storage",
        "storage/filesystem/dotgit",
        "utils/merkletrie/filesystem",
        "plumbing/object",
        "internal/url",
        "config",
        "plumbing/cache",
        "plumbing/protocol/packp/sideband",
        "internal/revision",
        "utils/merkletrie/index",
        "utils/merkletrie",
        "utils/merkletrie/noder",
        "plumbing/format/config",
        "plumbing/revlist",
        "plumbing/transport/http",
        "plumbing/format/gitignore",
        "plumbing/transport/git",
        "plumbing/format/diff",
        "utils/diff",
        "utils/binary",
        "utils/ioutil",
        "plumbing/protocol/packp",
        "plumbing/transport/server",
        "plumbing/transport/client",
        "plumbing",
        "plumbing/transport/ssh",
        "plumbing/protocol/packp/capability",
        "utils/merkletrie/internal/frame",
        "plumbing/transport",
        "storage/memory",
        "plumbing/format/objfile",
        "plumbing/transport/file",
        "plumbing/format/idxfile",
        "plumbing/filemode",
    ],
    deps = [
        ":gopkg.in.src-d.go-billy.v4",
        ":github.com.emirpasic.gods",
        ":github.com.sergi.go-diff",
        ":golang.org.x.crypto",
        ":github.com.src-d.gcfg",
        ":github.com.jbenet.go-context",
        ":github.com.xanzy.ssh-agent",
        ":github.com.kevinburke.ssh_config",
        ":golang.org.x.net",
        ":github.com.mitchellh.go-homedir",
    ],
    licences = ["Apache-2.0"]
)

go_module(
    name = "github.com.xanzy.ssh-agent",
    module = "github.com/xanzy/ssh-agent",
    version = "v0.2.1",
    deps = [
        ":golang.org.x.crypto",
    ],
    licences = ["Apache-2.0"]
)

go_module(
    name = "gopkg.in.src-d.go-billy.v4",
    module = "gopkg.in/src-d/go-billy.v4",
    version = "v4.3.2",
    install = [
        ".",
        "util",
        "osfs",
        "helper/chroot",
        "helper/polyfill",
    ],
    deps = [
        ":golang.org.x.sys",
    ],
    licences = ["Apache-2.0"]
)

go_module(
    name = "github.com.shurcooL.sanitized_anchor_name",
    module = "github.com/shurcooL/sanitized_anchor_name",
    version = "v1.0.0",
    licences = ["MIT"]
)

go_module(
    name = "gopkg.in.warnings.v0",
    module = "gopkg.in/warnings.v0",
    version = "v0.1.2",
    licences = ["BSD-2-Clause"]
)

go_module(
    name = "golang.org.x.xerrors",
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
    install = [
        ".",
        "internal",
    ],
    licences = ["BSD-3-Clause"]
)

go_module(
    name = "github.com.cpuguy83.go-md2man.v2",
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.0",
    install = [
        "md2man",
    ],
    deps = [
        ":github.com.russross.blackfriday.v2",
    ],
    licences = ["MIT"]
)

go_module(
    name = "github.com.mitchellh.go-homedir",
    module = "github.com/mitchellh/go-homedir",
    version = "v1.1.0",
    licences = ["MIT"]
)

go_module(
    name = "github.com.kevinburke.ssh_config",
    module = "github.com/kevinburke/ssh_config",
    version = "v0.0.0-20190725054713-01f96b0aa0cd",
    licences = ["MIT"]
)

go_module(
    name = "golang.org.x.crypto",
    module = "golang.org/x/crypto",
    version = "v0.0.0-20200622213623-75b288015ac9",
    install = [
        "ssh/knownhosts",
        "openpgp/s2k",
        "openpgp/armor",
        "blowfish",
        "chacha20",
        "openpgp/errors",
        "ssh/agent",
        "ssh/internal/bcrypt_pbkdf",
        "cast5",
        "openpgp/elgamal",
        "internal/subtle",
        "ssh",
        "curve25519",
        "poly1305",
        "ed25519",
        "openpgp",
        "openpgp/packet",
    ],
    licences = ["BSD-3-Clause"]
)

go_module(
    name = "github.com.emirpasic.gods",
    module = "github.com/emirpasic/gods",
    version = "v1.12.0",
    install = [
        "lists",
        "trees",
        "lists/arraylist",
        "containers",
        "trees/binaryheap",
        "utils",
    ],
    licences = ["BSD-2-Clause"]
)

go_module(
    name = "github.com.jbenet.go-context",
    module = "github.com/jbenet/go-context",
    version = "v0.0.0-20150711004518-d14ea06fba99",
    install = [
        "io",
    ],
    deps = [
        ":golang.org.x.net",
    ],
    licences = ["MIT"]
)

go_module(
    name = "golang.org.x.sys",
    module = "golang.org/x/sys",
    version = "v0.0.0-20200930185726-fdedc70b468f",
    install = [
        "unix",
        "internal/unsafeheader",
    ],
    licences = ["BSD-3-Clause"]
)

go_module(
    name = "github.com.sergi.go-diff",
    module = "github.com/sergi/go-diff",
    version = "v1.2.0",
    install = [
        "diffmatchpatch",
    ],
    licences = ["MIT"]
)

go_module(
    name = "golang.org.x.mod",
    module = "golang.org/x/mod",
    version = "v0.4.0",
    install = [
        "internal/lazyregexp",
        "module",
        "semver",
        "modfile",
    ],
    deps = [
        ":golang.org.x.xerrors",
    ],
    licences = ["BSD-3-Clause"]
)

go_module(
    name = "golang.org.x.net",
    module = "golang.org/x/net",
    version = "v0.0.0-20201021035429-f5854403a974",
    install = [
        "proxy",
        "context",
        "internal/socks",
    ],
    licences = ["BSD-3-Clause"]
)

go_module(
    name = "github.com.golang.glog",
    module = "github.com/golang/glog",
    version = "v0.0.0-20160126235308-23def4e6c14b",
    licences = ["Apache-2.0"]
)

go_module(
    name = "github.com.google.go-licenses",
    module = "github.com/google/go-licenses",
    version = "v0.0.0-20210816172045-3099c18c36e1",
    install = [
        "licenses",
    ],
    deps = [
        ":github.com.google.licenseclassifier",
        ":github.com.golang.glog",
        ":gopkg.in.src-d.go-git.v4",
        ":golang.org.x.tools",
    ],
    licences = ["Apache-2.0"]
)

go_module(
    name = "github.com.src-d.gcfg",
    module = "github.com/src-d/gcfg",
    version = "v1.4.0",
    install = [
        ".",
        "token",
        "scanner",
        "types",
    ],
    deps = [
        ":gopkg.in.warnings.v0",
    ],
    licences = ["BSD-3-Clause"]
)

