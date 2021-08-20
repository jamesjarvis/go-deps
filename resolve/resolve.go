package resolve

import (
	"context"
	"github.com/jamesjarvis/go-deps/host"
	"strings"
)

// TODO(jpoole): these could come from a .importconfig which might be cleaner
var KnownImports = map[string]struct{}{
	"archive/tar": {},
	"archive/zip": {},
	"bufio": {},
	"bytes": {},
	"compress/bzip2": {},
	"compress/flate": {},
	"compress/gzip": {},
	"compress/lzw": {},
	"compress/zlib": {},
	"container/heap": {},
	"container/list": {},
	"container/ring": {},
	"context": {},
	"crypto/aes": {},
	"crypto/cipher": {},
	"crypto/des": {},
	"crypto/dsa": {},
	"crypto/ecdsa": {},
	"crypto/ed25519": {},
	"crypto/ed25519/internal/edwards25519": {},
	"crypto/elliptic": {},
	"crypto/elliptic/internal/fiat/please": {},
	"crypto/hmac": {},
	"crypto": {},
	"crypto/internal/randutil": {},
	"crypto/internal/subtle": {},
	"crypto/md5": {},
	"crypto/rand": {},
	"crypto/rc4": {},
	"crypto/rsa": {},
	"crypto/sha1": {},
	"crypto/sha256": {},
	"crypto/sha512": {},
	"crypto/subtle": {},
	"crypto/tls": {},
	"crypto/x509": {},
	"crypto/x509/pkix": {},
	"database/sql/driver": {},
	"database/sql": {},
	"debug/dwarf": {},
	"debug/elf": {},
	"debug/gosym": {},
	"debug/macho": {},
	"debug/pe": {},
	"debug/plan9obj": {},
	"embed": {},
	"encoding/ascii85": {},
	"encoding/asn1": {},
	"encoding/base32": {},
	"encoding/base64": {},
	"encoding/binary": {},
	"encoding/csv": {},
	"encoding/gob": {},
	"encoding/hex": {},
	"encoding": {},
	"encoding/json": {},
	"encoding/pem": {},
	"encoding/xml": {},
	"errors": {},
	"expvar": {},
	"flag": {},
	"fmt": {},
	"go/ast": {},
	"go/build/constraint": {},
	"go/build": {},
	"go/constant": {},
	"go/doc": {},
	"go/format": {},
	"go/importer": {},
	"go/internal/gccgoimporter": {},
	"go/internal/gcimporter": {},
	"go/internal/srcimporter": {},
	"go/internal/typeparams": {},
	"go/parser": {},
	"go/printer": {},
	"go/scanner": {},
	"go/token": {},
	"go/types": {},
	"hash/adler32": {},
	"hash/crc32": {},
	"hash/crc64": {},
	"hash/fnv": {},
	"hash": {},
	"hash/maphash": {},
	"html": {},
	"html/template": {},
	"image/color": {},
	"image/color/palette": {},
	"image/draw": {},
	"image/gif": {},
	"image": {},
	"image/internal/imageutil": {},
	"image/jpeg": {},
	"image/png": {},
	"index/suffixarray": {},
	"internal/abi": {},
	"internal/buildcfg": {},
	"internal/bytealg": {},
	"internal/cfg": {},
	"internal/cpu": {},
	"internal/execabs": {},
	"internal/fmtsort": {},
	"internal/goexperiment": {},
	"internal/goroot": {},
	"internal/goversion": {},
	"internal/itoa": {},
	"internal/lazyregexp": {},
	"internal/lazytemplate": {},
	"internal/nettrace": {},
	"internal/obscuretestdata": {},
	"internal/oserror": {},
	"internal/poll": {},
	"internal/profile": {},
	"internal/race": {},
	"internal/reflectlite": {},
	"internal/singleflight": {},
	"internal/syscall/execenv": {},
	"internal/syscall/unix": {},
	"internal/sysinfo": {},
	"internal/testenv": {},
	"internal/testlog": {},
	"internal/trace": {},
	"internal/unsafeheader": {},
	"internal/xcoff": {},
	"io/fs": {},
	"io": {},
	"io/ioutil": {},
	"log": {},
	"log/syslog": {},
	"math/big": {},
	"math/bits": {},
	"math/cmplx": {},
	"math": {},
	"math/rand": {},
	"mime": {},
	"mime/multipart": {},
	"mime/quotedprintable": {},
	"net": {},
	"net/http/cgi": {},
	"net/http/cookiejar": {},
	"net/http/fcgi": {},
	"net/http": {},
	"net/http/httptest": {},
	"net/http/httptrace": {},
	"net/http/httputil": {},
	"net/http/internal/ascii/please": {},
	"net/http/internal": {},
	"net/http/internal/testcert/please": {},
	"net/http/pprof": {},
	"net/internal/socktest": {},
	"net/mail": {},
	"net/rpc": {},
	"net/rpc/jsonrpc": {},
	"net/smtp": {},
	"net/textproto": {},
	"net/url": {},
	"os/exec": {},
	"os": {},
	"os/signal": {},
	"os/signal/internal/pty/please": {},
	"os/user": {},
	"path/filepath": {},
	"path": {},
	"plugin": {},
	"reflect": {},
	"reflect/internal/example1": {},
	"reflect/internal/example2": {},
	"regexp": {},
	"regexp/syntax": {},
	"runtime/cgo": {},
	"runtime/debug": {},
	"runtime": {},
	"runtime/internal/atomic": {},
	"runtime/internal/math": {},
	"runtime/internal/sys": {},
	"runtime/metrics": {},
	"runtime/pprof": {},
	"runtime/race": {},
	"runtime/trace": {},
	"sort": {},
	"strconv": {},
	"strings": {},
	"sync/atomic": {},
	"sync": {},
	"syscall": {},
	"testing/fstest": {},
	"testing": {},
	"testing/internal/testdeps": {},
	"testing/iotest": {},
	"testing/quick": {},
	"text/scanner": {},
	"text/tabwriter": {},
	"text/template": {},
	"text/template/parse": {},
	"time": {},
	"time/tzdata": {},
	"unicode": {},
	"unicode/utf16": {},
	"unicode/utf8": {},
	"unsafe": {},
}

type resolver struct {
	// pkgs is a map of import paths to their package
	pkgs map[string]*Package
	modules map[string]*Module
	queue []string
	KnownImports map[string]struct{}
}

type Package struct {
	ImportPath string
	Module string
	Imports []*Package
	resolved bool
}

type Module struct {
	Name string
	Deps []*Module
	deps map[string]struct{}

	Install []string
	install map[string]struct{}
}


// ResolveGet resolves a `go get` style wildcard into a graph of packages
func ResolveGet(ctx context.Context, knownImports map[string]struct{}, getPath string) ([]*Module, error) {
	pkgs, err := host.GoListAll(ctx, getPath)
	if err != nil {
		return nil, err
	}

	r := &resolver{
		pkgs: map[string]*Package{},
		modules: map[string]*Module{},
		queue: pkgs,
		KnownImports: knownImports,
	}

	err = r.resolve(ctx)

	for _, pkg := range r.pkgs {
		module := r.getModule(pkg.Module)

		for _, i := range pkg.Imports {
			dep := r.getModule(i.Module)

			install := strings.Trim(strings.TrimPrefix(i.ImportPath, i.Module), "/")
			if _, ok := dep.install[install]; !ok {
				dep.Install = append(dep.Install, install)
				dep.install[install] = struct {}{}
			}


			if _, ok := module.deps[dep.Name]; i.Module != module.Name && !ok {
				module.Deps = append(module.Deps, dep)
				module.deps[dep.Name] = struct {}{}
			}
		}
	}

	ret := make([]*Module, 0, len(r.modules))
	for _, m := range r.modules {
		ret = append(ret, m)
	}

	return ret, err
}

func (r *resolver) resolve(ctx context.Context) error {
	for len(r.queue) > 0 {
		next := r.queue[0]
		r.queue = r.queue[1:]

		pkg := r.getPackage(next)
		if pkg.resolved {
			continue
		}

		resp, err := host.GoList(ctx, next)
		if err != nil {
			return err
		}

		pkg.Module = resp.Module

		for _, i := range resp.Imports {
			if _, ok := r.KnownImports[i]; ok {
				continue
			}
			if strings.HasPrefix(i, "vendor/") {
				continue
			}
			pkg.Imports = append(pkg.Imports, r.getPackage(i))
		}
		pkg.resolved = true
	}
	return nil
}

func (r *resolver) getPackage(path string) *Package {
	if pkg, ok := r.pkgs[path]; ok {
		return pkg
	}
	pkg := &Package{ImportPath: path, Imports: []*Package{}}
	r.pkgs[path] = pkg
	r.queue = append(r.queue, path)
	return pkg
}

func (r *resolver) getModule(path string) *Module {
	m, ok := r.modules[path]
	if !ok {
		m = &Module{
			Name: path,
			install: map[string]struct{}{},
			deps: map[string]struct{}{},
		}
		r.modules[path] = m
	}
	return m
}