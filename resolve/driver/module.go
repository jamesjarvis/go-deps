package driver

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/tools/go/packages"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func findGoImportMetaTag(body io.ReadCloser) string {
	page, _ := io.ReadAll(body)
	node, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		return ""
	}

	node = node.FirstChild.FirstChild
	var head *html.Node
	for {
		if node.Data == "head" {
			head = node
			break
		}
		node = node.NextSibling
		if node == nil {
			return ""
		}
	}
	if head == nil {
		return ""
	}

	node = head.FirstChild
	for {
		if node.Data == "meta" {
			content := ""
			isGoImport := false
			for _, attr := range node.Attr {
				if attr.Key == "name" && attr.Val == "go-import" {
					isGoImport = true
				} else if attr.Key == "content" {
					content = attr.Val
				}
			}
			if isGoImport {
				parts := strings.Split(content, " git ")
				if len(parts) == 2 {
					return parts[1]
				}
				return ""
			}
		}
		node = node.NextSibling
		if node == nil {
			return ""
		}
	}
}

func (driver *pleaseDriver) ensureDownloaded(mod *packages.Module) (srcRoot string, err error) {
	dir := filepath.Join(modcacheDir, mod.Path + "@" + mod.Version)
	if _, err := os.Stat(dir); err == nil {
		return dir, nil
	}

	if err := os.MkdirAll(dir, dirPerms); err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			os.RemoveAll(dir)
		}
	}()

	url := fmt.Sprintf("%s/%s/@v/%s.zip", driver.moduleProxy, mod.Path, mod.Version)
	dest := modcacheDir
	stripPrefix := ""
	// Download from github if we can when we're using a psudoversion
	if !strings.HasPrefix(mod.Version, "v") {
		githubRepo := fmt.Sprintf("https://%s", mod.Path)
		if !strings.HasPrefix(githubRepo, "https://github.com") {
			resp, err := client.Get(githubRepo)
			if err != nil {
				return "", err
			}

			if goImport := findGoImportMetaTag(resp.Body); goImport != "" {
				githubRepo = goImport
			} else {
				// follow any redirects from the server
				githubRepo = resp.Request.URL.String()
			}
		}

		if strings.HasPrefix(githubRepo, "https://github.com") {
			url = fmt.Sprintf("%s/archive/%s.zip", githubRepo, mod.Version)
			dest = filepath.Join(modcacheDir, mod.Path + "@" + mod.Version)
			stripPrefix = fmt.Sprintf("%s-%s", filepath.Base(githubRepo), mod.Version)
		}
	}

	resp, err := client.Get(strings.ToLower(url))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status %v downloading module %v@%v at %v\n%v", resp.StatusCode, mod.Path, mod.Version, url, string(body))
	}

	zipFilePath := filepath.Join(modcacheDir, "downloads", fmt.Sprintf("%s@%s/mod.zip", mod.Path, mod.Version))
	if err := os.MkdirAll(filepath.Dir(zipFilePath), dirPerms);  err != nil && !os.IsExist(err) {
		return "", err
	}
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(zipFile, resp.Body); err != nil {
		return "", err
	}
	zipFile.Close()

	zipReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return "", err
	}
	defer zipReader.Close()

	if err := unpackZip(zipReader, stripPrefix, dest); err != nil {
		return "", fmt.Errorf("failed to unpack zip: %v", err)
	}

	return dir, nil
}

func unpackZip(r *zip.ReadCloser, stripPrefix string, dest string) error {
	for _, f := range r.File {
		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, strings.TrimPrefix(f.Name, stripPrefix + "/"))

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// Make File
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func (driver *pleaseDriver) getKnownMods() error {
	if driver.knownModules != nil {
		return nil
	}
	driver.knownModules = []*packages.Module{}

	out := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cmd := exec.Command(driver.pleasePath, "query", "print", "-l", "go_module:", fmt.Sprintf("//%s/...", driver.thirdPartyFolder))
	cmd.Stdout = out
	cmd.Stderr = stdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to query known modules: %v\n%v\n%v", err, out, stdErr)
	}

	for _, line := range strings.Split(strings.TrimSpace(out.String()), "\n") {
		parts := strings.Split(line, "@")
		if len(parts) != 2 {
			return fmt.Errorf("invalid response listing known modules: %v", line)
		}

		driver.knownModules = append(driver.knownModules, &packages.Module{Path: parts[0], Version: strings.TrimSpace(parts[1])})
	}
	return nil
}

func (driver *pleaseDriver) findKnownModule(pattern string) *packages.Module {
	for _, mod := range driver.knownModules {
		if strings.HasPrefix(pattern, mod.Path) {
			return mod
		}
	}
	return nil
}

func (driver *pleaseDriver) resolveModuleForPackage(pattern string) (*packages.Module, error) {
	if err := driver.getKnownMods(); err != nil {
		return nil, err
	}
	mod := driver.findKnownModule(pattern)
	if mod != nil {
		return mod, nil
	}
	modulePath := strings.ToLower(strings.TrimSuffix(pattern,"/..."))

	for modulePath != "." {
		// TODO(jpoole): we should be aware of a few common module formats
		resp, err := client.Get(fmt.Sprintf("%s/%s/@latest", driver.moduleProxy, modulePath))
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			modulePath = filepath.Dir(modulePath)
			continue
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		version := struct {
			Version string
		}{}
		if err := json.Unmarshal(b, &version); err != nil {
			return nil, err
		}
		mod = &packages.Module{Path: modulePath, Version: version.Version}
		driver.knownModules = append(driver.knownModules, mod)
		return mod, nil
	}
	return nil, fmt.Errorf("couldn't find module for package %v", pattern)
}
