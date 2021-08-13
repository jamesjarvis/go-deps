package module

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jamesjarvis/go-deps/host"
	"golang.org/x/mod/semver"
)

var GlobalCache *Directory

func init() {
	// I know this is kinda dirty code but I couldn't be bothered to
	// think about this too hard for the initial implementation....
	GlobalCache = NewDirectory()
}

// Directory will be the global cache of seen modules, it will be responsible
// for storing modules and ultimately resolving module version clashes.
type Directory struct {
	modules map[string]*VersionDirectory
}

func NewDirectory() *Directory {
	return &Directory{
		modules: map[string]*VersionDirectory{},
	}
}

// Sync is a lazy implementation to refresh all of the module dependencies to the closest semver.
func (d *Directory) Sync() {
	for _, vd := range d.modules {
		for _, mod := range vd.versions {
			// Flag this module as requiring to specify the version as we have multiple versions.
			mod.nameWithVersion = len(vd.versions) > 1
			for i, dep := range mod.Deps {
				closestMod := d.GetClosestModule(dep.Path, dep.Version)
				if dep != closestMod {
					log.Printf("Synced %s --> %s\n", dep.String(), closestMod.String())
				}
				mod.Deps[i] = closestMod
			}
		}
	}
}

type Writer interface {
	Write(modules map[string]*VersionDirectory)
}

func (d *Directory) Write(writer Writer) {
	writer.Write(d.modules)
}


func (d *Directory) Print() {
	// Sort the paths to deterministically print output.
	paths := make([]string, 0, len(d.modules))
	for path := range d.modules {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		vd := d.modules[path]
		fmt.Printf("MODULE: %s\n", path)
		// Sort the versions to deterministically print output.
		versions := make([]string, 0, len(vd.versions))
		for version := range vd.versions {
			versions = append(versions, version)
		}
		sort.Strings(versions)
		for _, version := range versions {
			mod := vd.versions[version]
			fmt.Printf("\tVERSION: %s\n", version)
			fmt.Printf("\t\t%s\n", mod.String())
			if len(mod.Deps) > 0 {
				fmt.Printf("\t\t|\n")
			}
			for _, dep := range mod.Deps {
				fmt.Printf("\t\t|---- %s\n", dep.String())
			}
		}
	}
}

func (d *Directory) ExportBuildRules() error {
	// Delete all existing third party build files.
	err := host.RemoveAllThirdPartyFiles()
	if err != nil {
		return err
	}
	// Sort the paths to deterministically write build files.
	paths := make([]string, 0, len(d.modules))
	for path := range d.modules {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		vd := d.modules[path]
		// Sort the versions to deterministically write build files.
		versions := make([]string, 0, len(vd.versions))
		for version := range vd.versions {
			versions = append(versions, version)
		}
		sort.Strings(versions)
		for _, version := range versions {
			mod := vd.versions[version]
			buildFilePath := mod.GetBuildPath()
			if _, err := os.Stat(buildFilePath); os.IsNotExist(err) { 
				err = os.MkdirAll(strings.TrimSuffix(buildFilePath, "/BUILD"), 0700) // Create the nested directory
				if err != nil {
					return fmt.Errorf("failed to create directory: %w", err)
				}
			}
			f, err := os.OpenFile(buildFilePath,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			defer f.Close()
			err = mod.WriteGoModuleRule(f)
			if err != nil {
				return fmt.Errorf("failed to append go_module to file: %w", err)
			}
		}
	}
	return nil
}

func (d *Directory) Get(path string) *VersionDirectory {
	vd, ok := d.modules[path]
	if !ok {
		return nil
	}
	return vd
}

func (d *Directory) GetModule(path, version string) *Module {
	vd := d.Get(path)
	if vd == nil {
		return nil
	}
	mod := vd.GetVersion(version)
	if mod == nil {
		return nil
	}
	return mod
}

func (d *Directory) GetClosestModule(path, version string) *Module {
	vd := d.Get(path)
	if vd == nil {
		return nil
	}
	closestVersion := vd.GetClosestVersion(version)
	mod := vd.GetVersion(closestVersion)
	if mod == nil {
		return nil
	}
	return mod
}

func (d *Directory) SetModule(mod *Module) *Module {
	vd := d.Get(mod.Path)
	if vd == nil {
		vd = NewVersionDirectory()
	}

	fixedMod := vd.SetVersion(mod.Version, mod)
	d.Set(mod.Path, vd)
	return fixedMod
}

func (d *Directory) Set(path string, vd *VersionDirectory) {
	d.modules[path] = vd
}

type VersionDirectory struct {
	versions map[string]*Module
}

func NewVersionDirectory() *VersionDirectory {
	return &VersionDirectory{
		versions: map[string]*Module{},
	}
}

func (vd *VersionDirectory) GetVersion(version string) *Module {
	mod, ok := vd.versions[version]
	if !ok {
		return nil
	}
	return mod
}

// GetClosestVersion returns the highest semver within the same major version,
// or itself if no matches found.
func (vd *VersionDirectory) GetClosestVersion(version string) string {
	major := semver.Major(version)
	for existingVers := range vd.versions {
		existingMajor := semver.Major(existingVers)
		if major == existingMajor {
			comparison := semver.Compare(version, existingVers)
			// If incoming is less than existing, return existing.
			if comparison < 0 {
				return existingVers
			}
		}
	}
	return version
}

func (vd *VersionDirectory) SetVersion(version string, mod *Module) *Module {
	// If this version already exists, overwrite with the new one and return.
	if existing := vd.GetVersion(version); existing != nil {
		vd.versions[version] = mod
		return mod
	}

	// Get existing versions.
	// If there is an existing version with the same major, but higher minor,
	// then return the existing version as it is "better".
	// If there is an existing version with the same major, but lower minor,
	// then delete that version, and store this "better" version and return.
	// If there is an existing version with different major, or none at all,
	// then store this "different" version and return.
	// 
	// This should eventually lead to there only being one entry
	// for each major version.
	// A different implementation may choose to have stricter/more relaxed
	// version resolving logic.
	major := semver.Major(version)
	for existingVers, existingMod := range vd.versions {
		existingMajor := semver.Major(existingVers)
		if major == existingMajor {
			comparison := semver.Compare(version, existingVers)
			// If incoming is greater than existing, replace and return.
			if comparison > 0 {
				delete(vd.versions, existingVers)
				vd.versions[version] = mod
				return mod
			}
			// If incoming is less than existing, return existing.
			return existingMod
		}
	}

	vd.versions[version] = mod
	return mod
}
