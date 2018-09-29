package testhelper

import (
	"errors"

	"github.com/tehsphinx/astrav"
)

//LoadExample loads a example from given example folder and package name
func LoadExample(exampleFolder string, pkgName string) (*astrav.Folder, *astrav.Package, error) {
	folder := astrav.NewFolder(exampleFolder)
	pkgs, err := folder.ParseFolder()
	if err != nil {
		return nil, nil, err
	}

	//TODO: get pkgName automatically without a parameter
	pkg, ok := pkgs[pkgName]
	if !ok {
		return nil, nil, errors.New("given package name not found in given folder")
	}

	return folder, pkg, nil
}
