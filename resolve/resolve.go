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
	importPaths  map[string]*Module
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
	return strings.Trim(strings.TrimPrefix(pkg.ImportPath, pkg.Module), "/")
}

// Module represents a module. It includes all deps so actually represents a full module graph.
type Module struct {
	// The module name
	Name string

	// The postfix to apply to the rule
	Postfix string

	// Externally download with go_mod_download()
	External bool

	// The packages in this module
	Packages map[*Package]struct{}
}

type ModuleGraph struct {
	ModuleRules []*ModuleRule
	DownloadRules []*DownloadRule
}

type ModuleRule struct {
	Name string
	Module string
	Installs []string
	Deps []string
	Download string
	Version string
}

type DownloadRule struct {
	Name string
	Module string
	Version string
}

// cycleResolution represents the imports that need to be moved out of the "to" module in order to resolve this cycle
type cycleResolution struct {
	from *Module
	to *Module
	imports []*Package
}

func ruleName(path string) string {
	return strings.ReplaceAll(path, "/", ".")
}

func (m *Module) Imports(importPaths map[string]*Module) map[*Package]struct{} {
	ret := map[*Package]struct{}{}
	for pkg := range m.Packages {
		for _, i := range pkg.Imports {
			mod := importPaths[i.ImportPath]
			if mod == m {
				continue
			}
			ret[i] = struct{}{}
		}
	}
	return ret
}

func (r *resolver) resolveCycles(modules []*Module) []*Module {
	for _, m := range modules {
		for i := range m.Imports(r.importPaths) {
			dep := r.importPaths[i.ImportPath]
			c := r.findCycle(m, m, dep)
			if c == nil {
				continue
			}

			oldRule := c.to

			// Mark the old rule as needing an external go_mod_download()
			oldRule.External = true

			count := r.moduleCounts[oldRule.Name]
			count++
			r.moduleCounts[oldRule.Name] = count

			newRule := &Module{
				Name:     oldRule.Name,
				Postfix:  fmt.Sprintf("_%v", count), //TODO(jpoole): find out what index we're on
				External: true,
				Packages: map[*Package]struct{}{},
			}

			for _, pkg := range c.imports {
				// Make imports to this package go to this new rule
				r.importPaths[pkg.ImportPath] = newRule

				// Move the installed package over from the old rule to the new rule
				newRule.Packages[pkg] = struct{}{}
				delete(oldRule.Packages, pkg)
			}

			return r.resolveCycles(append(modules, newRule))
		}
	}
	return modules
}

func (r *resolver) dependsOn(pkg *Package, module *Module) bool {
	if module == r.importPaths[pkg.ImportPath] {
		return true
	}
	for _, i := range pkg.Imports {
		if r.dependsOn(i, module) {
			return true
		}
	}
	return false
}

func (r *resolver) buildResolution(from *Module, to *Module, tryReverse bool) *cycleResolution {
	c := &cycleResolution{
		from: from,
		to: to,
	}

	done := map[*Package]struct{}{}
	for pkg := range c.from.Packages {
		for _, i := range pkg.Imports {
			if _, ok := done[i]; !ok && r.dependsOn(i, to) {
				if r.dependsOn(i, c.from) {
					continue
				}
				c.imports = append(c.imports, i)
				done[i] = struct{}{}
			}
		}
	}
	if len(c.imports) == 0 {
		if !tryReverse {
			panic("couldn't find a resolution")
		}
		return r.buildResolution(to, from, false)
	}
	return c
}

func (r *resolver) findCycle(start *Module, last, next *Module) *cycleResolution {
	if next == start {
		return r.buildResolution(last, next, true)
	}

	for i := range next.Imports(r.importPaths) {
		dep := r.importPaths[i.ImportPath]
		if dep == next {
			continue
		}
		if c := r.findCycle(start, next, dep); c != nil {
			return c
		}
	}

	return nil
}

// ResolveGet resolves a `go get` style wildcard into a graph of packages
func ResolveGet(ctx context.Context, knownImports map[string]struct{}, getPath string) (*ModuleGraph, error) {
	pkgs, err := host.GoListAll(ctx, getPath)
	if err != nil {
		return nil, err
	}

	r := &resolver{
		pkgs:         map[string]*Package{},
		modules:      map[string]*Module{},
		importPaths:  map[string]*Module{},
		moduleCounts: map[string]int{},
		queue:        pkgs,
		knownImports: knownImports,
	}

	err = r.resolve(ctx)
	if err != nil {
		return nil, err
	}

	for _, pkg := range r.pkgs {
		module := r.getModule(pkg.Module)
		module.Packages[pkg] = struct{}{}
		r.importPaths[pkg.ImportPath] = module

		for _, i := range pkg.Imports {
			if i.Module != module.Name {
				module.Imports(r.importPaths)[i] = struct{}{}
			}
		}
	}

	modules := make([]*Module, 0, len(r.modules))
	for _, m := range r.modules {
		modules = append(modules, m)
	}

	modules = r.resolveCycles(modules)

	ret := new(ModuleGraph)
	ret.ModuleRules = make([]*ModuleRule, 0, len(modules))

	dlRules := make(map[string]struct{})
	for _, module := range modules {
		ret.ModuleRules = append(ret.ModuleRules, r.toModuleRule(module))
		if module.External {
			if _, ok := dlRules[module.Name]; !ok {
				ret.DownloadRules = append(ret.DownloadRules, &DownloadRule{
					Name:    ruleName(module.Name) + "_dl",
					Module:  module.Name,
					Version: getVersion(module.Name),
				})
			}
			dlRules[module.Name] = struct{}{}
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
			Packages: map[*Package]struct{}{},
		}
		r.modules[path] = m
	}
	return m
}

func (r *resolver) toModuleRule(module *Module) *ModuleRule {
	rule := &ModuleRule{
		Name: ruleName(module.Name) + module.Postfix,
		Module: module.Name,
	}
	for pkg := range module.Packages {
		rule.Installs = append(rule.Installs, pkg.toInstall())
	}

	if module.External {
		rule.Download = ruleName(module.Name) + "_dl"
	} else {
		rule.Version = getVersion(module.Name)
	}

	done := map[string]struct{}{}
	for i := range module.Imports(r.importPaths) {
		dep := r.importPaths[i.ImportPath]
		name := ruleName(dep.Name)
		if dep.Postfix != "" {
			name = name + dep.Postfix
		}
		if _, ok := done[name]; ok {
			continue
		}
		done[name] = struct{}{}
		rule.Deps = append(rule.Deps, name)
	}
	return rule
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