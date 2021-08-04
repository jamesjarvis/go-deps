package module

import (
	"fmt"

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

func (d *Directory) Print() {
	for path, vd := range d.modules {
		fmt.Printf("MODULE: %s\n", path)
		for version, mod := range vd.versions {
			fmt.Printf("\tVERSION: %s\n", version)
			fmt.Printf("\t\t%s\n", mod.String())
		}
	}
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
