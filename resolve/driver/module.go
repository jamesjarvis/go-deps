package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
	"golang.org/x/tools/go/packages"

	"github.com/tatskaari/go-deps/progress"
)

// goModDownloadRule represents a `go_mod_download()` rule from Please BUILD files
type goModDownloadRule struct {
	label   string
	built   bool
	srcRoot string
}

// ensureDownloaded ensures the a module has been downloaded and returns the filepath to its source root
func (driver *pleaseDriver) ensureDownloaded(mod *packages.Module) (srcRoot string, err error) {
	key := fmt.Sprintf("%v@%v", mod.Path, mod.Version)
	if path, ok := driver.downloaded[key]; ok {
		return path, nil
	}

	if target, ok := driver.pleaseModules[mod.Path]; ok {
		if target.built {
			return target.srcRoot, nil
		}
		cmd := exec.Command(driver.pleasePath, "build", target.label)
		progress.PrintUpdate("Building %s...", target.label)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("failed to build %v: %v\n%v", target.label, err, string(out))
		}

		target.built = true
		return target.srcRoot, nil
	}

	oldWd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if _, err := os.Lstat("plz-out/godeps/go.mod"); err != nil {
		if os.IsNotExist(err) {
			cmd := exec.Command("go", "mod", "init", "dummy")
			cmd.Dir = "plz-out/godeps"
			out, err := cmd.CombinedOutput()
			if err != nil {
				return "", fmt.Errorf("failed to create dummy mod: %v\n%v", err, string(out))
			}
		} else {
			return "", err
		}
	}

	var resp = struct {
		Path string
		GoMod string
		Version string
		Dir string
		Error string
	}{}

	cmd := exec.Command("go", "mod", "download", "--json", key)
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOPATH=%s", filepath.Join(oldWd, "plz-out/godeps/go")))
	cmd.Dir = "plz-out/godeps"
	progress.PrintUpdate("Downloading %s...", key)
	out, err := cmd.CombinedOutput()

	if err != nil {
		json.Unmarshal(out, &resp)
		errorString := string(out)
		if resp.Error != "" {
			s, e := strconv.Unquote(resp.Error)
			if e == nil {
				errorString = s
			}
		}
		return "", fmt.Errorf("failed to download module %v: %v\n%v", key, err, errorString)
	}
	if err := json.Unmarshal(out, &resp); err != nil {
		return "", err
	}

	driver.downloaded[key] = resp.Dir

	return resp.Dir, nil
}

// getGoMod returns the go mod for a modules from the proxy
func (driver *pleaseDriver) getGoMod(mod, ver string) (*modfile.File, error) {
	file := fmt.Sprintf("%s/%s/@v/%s.mod", driver.moduleProxy, strings.ToLower(mod), ver)
	resp, err := client.Get(file)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%v %v: \n%v", file, resp.StatusCode, string(body))
	}

	return modfile.Parse(file, body, nil)

}

// determineVersionRequirements loads the version requirements from the go.mod files for each module, and applies
// the minimum valid version algorithm.
func (driver *pleaseDriver) determineVersionRequirements(mod, ver string) error {
	if oldVer, ok := driver.moduleRequirements[mod]; ok {
		// if we already require at this version or higher, we don't need to do anything
		if semver.Compare(ver, oldVer.Version) <= 0 {
			return nil
		}
	}

	if mod == "" {
		panic(mod)
	}

	progress.PrintUpdate("Resolving %v@%v", mod, ver)

	modFile, err := driver.getGoMod(mod, ver)
	if err != nil {
		ver := fmt.Sprintf("%v-incompatible", ver)
		modFile, err = driver.getGoMod(mod, ver)
		if err != nil {
			return err
		}
	}

	driver.moduleRequirements[mod] = &packages.Module{Path: mod, Version: ver}
	for _, req := range modFile.Require {
		if err := driver.determineVersionRequirements(req.Mod.Path, req.Mod.Version); err != nil {
			return err
		}
	}
	return nil
}

// resolveGetModules resolves the get wildcards with versions, and loads them into the driver. It returns the package
// parts of the get patterns e.g. github.com/example/module/...@v1.0.0 -> github.com/example/module/...
func (driver *pleaseDriver) resolveGetModules(patterns []string) ([]string, error) {
	pkgWildCards := make([]string, 0, len(patterns))
	for _, p := range patterns {
		parts := strings.Split(p, "@")
		pkgPart := parts[0]
		pkgWildCards = append(pkgWildCards, pkgPart)

		mod, err := driver.resolveModuleForPackage(pkgPart)
		if err != nil {
			return nil, err
		}
		if len(parts) > 1 && strings.HasPrefix(parts[1], "v") {
			if err := driver.determineVersionRequirements(mod, parts[1]); err != nil {
				return nil, err
			}
		} else {
			ver, err := driver.getLatestVersion(mod)
			if err != nil {
				return nil, err
			}
			if err := driver.determineVersionRequirements(mod, ver); err != nil {
				return nil, err
			}
		}

	}
	return pkgWildCards, nil
}

// loadPleaseModules queries the Please build graph and loads in any modules defined there. It applies the minimum valid
// version algorithm.
func (driver *pleaseDriver) loadPleaseModules() error {
	out := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cmd := exec.Command(driver.pleasePath, "query", "print", "-i", "go_module", "--json", fmt.Sprintf("//%s/...", driver.thirdPartyFolder))
	cmd.Stdout = out
	cmd.Stderr = stdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to query known modules: %v\n%v\n%v", err, out, stdErr)
	}

	res := map[string]struct{
		Outs []string
		Labels []string
	}{}

	if err := json.Unmarshal(out.Bytes(), &res); err != nil {
		return err
	}

	for label, target := range res {
		rule := &goModDownloadRule{
			label:   label,
			srcRoot: filepath.Join("plz-out/gen", target.Outs[0]),
		}
		for _, l := range target.Labels {
			if strings.HasPrefix(l, "go_module:") {
				parts := strings.Split(strings.TrimPrefix(l, "go_module:"), "@")
				if len(parts) != 2 {
					return fmt.Errorf("invalid go_module label: %v", l)
				}

				mod := &packages.Module{Path: parts[0], Version: strings.TrimSpace(parts[1])}
				oldMod, ok := driver.moduleRequirements[mod.Path]

				// Only add the Please version of this module if it's greater than or equal to the version requirement
				if !ok || semver.Compare(oldMod.Version, mod.Version) <= 0 {
					driver.moduleRequirements[mod.Path] = mod
					driver.pleaseModules[mod.Path] = rule
				}
			}
		}

	}
	return nil
}

// findKnownModule checks a list of discovered modules to see if the package pattern exists there
func (driver *pleaseDriver) findKnownModule(pattern string) string {
	longestMatch := ""
	for _, mod := range driver.knownModules {
		if pattern == mod {
			return mod
		}
		if strings.HasPrefix(pattern, mod + "/") {
			if len(mod) > len(longestMatch) {
				longestMatch = mod
			}
		}
	}
	return longestMatch
}


// getLatestVersion returns the latest versin for a mdoule from the proxy
func (driver *pleaseDriver) getLatestVersion(modulePath string) (string, error) {
	if modulePath == "" {
		panic(modulePath)
	}
	resp, err := client.Get(fmt.Sprintf("%s/%s/@latest", driver.moduleProxy, modulePath))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	version := struct {
		Version string
	}{}
	if err := json.Unmarshal(b, &version); err != nil {
		return "", err
	}
	return version.Version, nil
}

// resolveModuleForPackage tries to determine the module name for a given package pattern
func (driver *pleaseDriver) resolveModuleForPackage(pattern string) (string, error) {
	mod := driver.findKnownModule(pattern)
	if mod != "" {
		return mod, nil
	}
	modulePath := strings.ToLower(strings.TrimSuffix(pattern,"/..."))

	for modulePath != "." {
		if strings.HasPrefix(pattern, "github.com") {
			parts := strings.Split(pattern, "/")

			if len(parts) < 3 {
				return "", fmt.Errorf("can't determine module for package %v", pattern)
			}
			modPart := 3
			if len(parts) >= 4 && strings.HasPrefix(parts[3], "v") {
				modPart++
			}
			mod := filepath.Join(parts[:modPart]...)
			driver.knownModules  = append(driver.knownModules, mod)
			return mod, nil
		}

		// Try and get the latest version to determine if we've found the module part yet
		version, err := driver.getLatestVersion(modulePath)
		if err != nil {
			return "", err
		}
		if version == "" {
			modulePath = filepath.Dir(modulePath)
			continue
		}

		driver.knownModules = append(driver.knownModules, modulePath)
		return modulePath, nil
	}
	return "", fmt.Errorf("couldn't find module for package %v", pattern)
}
