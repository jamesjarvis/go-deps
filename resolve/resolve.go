package resolve

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jamesjarvis/go-deps/host"
	"os/exec"
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
	pkgs         map[string]*Package
	modules      map[string]*Module
	queue        []string
	knownImports map[string]struct{}
	importPaths  map[*Package]*ModulePart
	moduleCounts map[string]int
}

// Package represents a single package in some module
type Package struct {
	// The full import path of this package
	ImportPath string

	// The module name this package belongs to
	Module string

	// Any other packages this package imports
	Imports []*Package

	// Whether this package has been resolved and the above fields are populated
	resolved bool
}

func (pkg *Package) toInstall() string {
	install := strings.Trim(strings.TrimPrefix(pkg.ImportPath, pkg.Module), "/")
	if install == "" {
		return "."
	}
	return install
}

// Module represents a module. It includes all deps so actually represents a full module graph.
type Module struct {
	// The module name
	Name string

	Parts []*ModulePart
}

type ModulePart struct {
	Module *Module
	// The packages in this module
	Packages map[*Package]struct{}
	// The index of this module part
	Index int
}


type ModuleRules struct {
	Download *DownloadRule
	Version string
	Mods []*ModuleRule
}

type ModuleRule struct {
	Name string
	Module string
	Installs []string
	Deps []string
	ExportedDeps []string
}

type DownloadRule struct {
	Name string
	Module string
	Version string
}

func ruleName(path, suffix string) string {
	return 	strings.ReplaceAll(path, "/", ".") + suffix
}

func partName(part *ModulePart) string {
	displayIndex := len(part.Module.Parts) - part.Index

	if displayIndex > 0 {
		return ruleName(part.Module.Name, fmt.Sprintf("_%d", displayIndex))
	}
	return ruleName(part.Module.Name, "")
}

func (r *resolver) dependsOn(pkg *Package, module *ModulePart, leftModule bool) bool {
	if leftModule && module == r.importPaths[pkg] {
		return true
	}
	if module != r.importPaths[pkg] {
		leftModule = true
	}
	for _, i := range pkg.Imports {
		if r.dependsOn(i, module, leftModule) {
			return true
		}
	}
	return false
}


func (r *resolver) addPackageToModuleGraph(done map[*Package]struct{}, pkg *Package) {
	if _, ok := done[pkg]; ok {
		return
	}

	for _, i := range pkg.Imports {
		r.addPackageToModuleGraph(done, i)
	}

	m := r.getModule(pkg.Module)


	var validPart *ModulePart
	for _, part := range m.Parts {
		valid := true
		for _, i := range pkg.Imports {
			if r.dependsOn(i, part, false) {
				valid = false
				break
			}
		}
		if valid {
			validPart = part
			break
		}
	}
	if validPart == nil {
		validPart = &ModulePart{
			Packages: map[*Package]struct{}{},
			Module: m,
			Index: len(m.Parts) + 1,
		}
		m.Parts = append(m.Parts, validPart)
	}

	validPart.Packages[pkg] = struct{}{}
	r.importPaths[pkg] = validPart

	done[pkg] = struct{}{}
}


// ResolveGet resolves a `go get` style wildcard into a graph of packages
func ResolveGet(ctx context.Context, knownImports map[string]struct{}, getPath string) ([]*ModuleRules, error) {
	pkgs, err := host.GoListAll(ctx, getPath)
	if err != nil {
		return nil, err
	}

	r := &resolver{
		pkgs:         map[string]*Package{},
		modules:      map[string]*Module{},
		importPaths:  map[*Package]*ModulePart{},
		moduleCounts: map[string]int{},
		queue:        pkgs,
		knownImports: knownImports,
	}

	err = r.resolve(ctx)
	if err != nil {
		return nil, err
	}

	done := make(map[*Package]struct{}, len(pkgs))
	for _, pkg := range r.pkgs {
		r.addPackageToModuleGraph(done, pkg)
	}

	modules := make([]*Module, 0, len(r.modules))
	for _, m := range r.modules {
		modules = append(modules, m)
	}

	ret := make([]*ModuleRules, 0, len(r.modules))
	for _, m := range r.modules {
		rules := new(ModuleRules)
		ret = append(ret, rules)

		if len(m.Parts) > 1 {
			rules.Download = &DownloadRule{
				Name:    ruleName(m.Name, "_dl"),
				Module:  m.Name,
				Version: getVersion(m.Name),
			}
		} else {
			rules.Version = getVersion(m.Name)
		}

		for _, part := range m.Parts {
			modRule := &ModuleRule{
				Name:   partName(part),
				Module: m.Name,
			}

			done := map[string]struct{}{}
			for pkg := range part.Packages {
				modRule.Installs = append(modRule.Installs, pkg.toInstall())
				for _, i := range pkg.Imports {
					dep := r.importPaths[i]
					depRuleName := partName(dep)
					if _, ok := done[depRuleName]; ok || dep.Module == m {
						continue
					}
					done[depRuleName] = struct{}{}

					modRule.Deps = append(modRule.Deps, depRuleName)
				}
			}

			if part.Index == len(m.Parts) {
				// TODO(jpoole): is this a safe assumption? Can we have more than 2 parts ot a
				for _, part := range m.Parts[:(len(m.Parts)-1)] {
					modRule.ExportedDeps = append(modRule.ExportedDeps, partName(part))
				}
			}

			rules.Mods = append(rules.Mods, modRule)
		}
	}


	return ret, nil
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
			if _, ok := r.knownImports[i]; ok {
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
		}
		r.modules[path] = m
	}
	return m
}

type DownloadResponse struct {
	Version string
}

func getVersion(module string) string {
	cmd := exec.Command("go", "mod", "download", "--json", module)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("failed to get module version: %v\n%v", err, string(out)))
	}

	resp := new(DownloadResponse)
	err = json.Unmarshal(out, resp)
	if err != nil {
		panic(err)
	}
	return resp.Version
}