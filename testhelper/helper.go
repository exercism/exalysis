package testhelper

import (
	"errors"

	"github.com/tehsphinx/astrav"
)

// LoadFolder loads an example from given example folder
func LoadFolder(exampleFolder string) (*astrav.Folder, map[string]*astrav.Package, error) {
	folder := astrav.NewFolder(exampleFolder)
	pkgs, err := folder.ParseFolder()
	return folder, pkgs, err
}

// LoadExample loads an example from given example folder and package name
func LoadExample(exampleFolder string, pkgName string) (*astrav.Folder, *astrav.Package, error) {
	folder, pkgs, err := LoadFolder(exampleFolder)
	if err != nil {
		return nil, nil, err
	}

	// TODO: get pkgName automatically without a parameter
	pkg, ok := pkgs[pkgName]
	if !ok {
		return nil, nil, errors.New("given package name not found in given folder")
	}

	return folder, pkg, nil
}
