package driver

import (
	"github.com/tatskaari/go-deps/resolve/knownimports"
	"go/build"
	"golang.org/x/tools/go/packages"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const modcacheDir = "plz-out/godeps/modcache"
const dirPerms = os.ModeDir | 0775

var client = http.DefaultClient

type pleaseDriver struct {
	moduleProxy      string
	thirdPartyFolder string
	pleasePath       string
	knownModules     []*packages.Module

	packages map[string]*packages.Package
}

type packageInfo struct {
	id              string
	srcRoot, pkgDir string
	mod             *packages.Module
	isSDKPackage    bool
}

func NewPleaseDriver(please, thirdPartyFolder string) *pleaseDriver {
	proxy := os.Getenv("GOPROXY") //TODO(jpoole): split this on , and get rid of direct
	if proxy == "" {
		proxy = "https://proxy.golang.org"
	}

	return &pleaseDriver{pleasePath: please, thirdPartyFolder: thirdPartyFolder, moduleProxy: proxy}
}

func (driver *pleaseDriver) pkgInfo(id string) (*packageInfo, error) {
	if knownimports.IsKnown(id) {
		srcDir := filepath.Join(build.Default.GOROOT, "src")
		return &packageInfo{isSDKPackage: true, id: id, srcRoot: srcDir, pkgDir: filepath.Join(srcDir, id)}, nil
	}

	mod, err := driver.resolveModuleForPackage(id)
	if err != nil {
		return nil, err
	}

	srcRoot, err := driver.ensureDownloaded(mod)
	if err != nil {
		return nil, err
	}

	return &packageInfo{
		id:      id,
		srcRoot: srcRoot,
		pkgDir:  filepath.Join(srcRoot, strings.TrimPrefix(id, mod.Path)),
		mod:     mod,
	}, nil
}

func (driver *pleaseDriver) importPattern(pattern string) ([]string, error) {
	walk := strings.HasSuffix(pattern, "...")

	info, err := driver.pkgInfo(strings.TrimSuffix(pattern, "/..."))
	if err != nil {
		return nil, err
	}

	if walk {
		var roots []string

		err := filepath.Walk(info.pkgDir, func(path string, i fs.FileInfo, err error) error {
			if !i.IsDir() {
				return nil
			}

			id := filepath.Join(info.mod.Path, strings.TrimPrefix(path, info.srcRoot))
			info, err := driver.pkgInfo(strings.TrimSuffix(id, "/..."))
			if err != nil {
				return err
			}

			if err :=  driver.importPackage(info); err != nil {
				if _, ok := err.(*build.NoGoError); ok {
					return nil
				}
				return err
			}
			roots = append(roots, id)
			return nil
		})
		return roots, err
	} else {
		return []string{info.id}, driver.importPackage(info)
	}
}

func (driver *pleaseDriver) importPackage(info *packageInfo) error {
	if _, ok := driver.packages[info.id]; ok {
		return nil
	}

	pkg, err := build.ImportDir(info.pkgDir, build.ImportComment)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	imports := map[string]*packages.Package{}
	for _, i := range pkg.Imports {
		if i == "C" {
			return nil
		}
		imports[i] = &packages.Package{ID: i}
		info, err := driver.pkgInfo(i)
		if err != nil {
			return err
		}
		if err := driver.importPackage(info); err != nil {
			return err
		}
	}

	goFiles := make([]string, 0, len(pkg.GoFiles))
	for _, f := range pkg.GoFiles {
		goFiles = append(goFiles, filepath.Join(wd, info.pkgDir, f))
	}

	driver.packages[info.id] = &packages.Package{
		ID:      info.id,
		Name:    pkg.Name,
		PkgPath: info.id,
		GoFiles: goFiles,
		Imports: imports,
		Module:  info.mod,
	}
	return nil
}

func (driver *pleaseDriver) Resolve(cfg *packages.Config, patterns ...string) (*packages.DriverResponse, error) {
	driver.packages = map[string]*packages.Package{}
	os.RemoveAll("plz-out/godeps")

	resp := new(packages.DriverResponse)
	for _, p := range patterns {
		parts := strings.Split(p, "@")
		if len(parts) > 1 {
			mod, err := driver.resolveModuleForPackage(parts[0])
			if err != nil {
				return nil, err
			}
			mod.Version = parts[1]
		}
		pkgs, err := driver.importPattern(parts[0])
		if err != nil {
			return nil, err
		}
		resp.Roots = append(resp.Roots, pkgs...)
	}

	resp.Packages = make([]*packages.Package, 0, len(driver.packages))
	for _, pkg := range driver.packages {
		resp.Packages = append(resp.Packages, pkg)
	}
	return resp, nil
}
