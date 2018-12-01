package exam

import (
	"os"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
)

// Result contains the result if running all examinations
type Result struct {
	GoLint       bool
	GoFmt        bool
	GoTest       bool
	GoBench      bool
	GoVet        bool
	GolangCILint bool
}

// All runs all examinations contained in this package and returns the result
func All(folder *astrav.Folder, r *extypes.Response, pkgName string) (res *Result, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if err := os.Chdir(folder.GetPath()); err != nil {
		return nil, err
	}
	defer func() {
		if r := os.Chdir(dir); r != nil {
			err = r
		}
	}()

	return &Result{
		GoLint:       GoLint(folder, r, pkgName),
		GoFmt:        GoFmt(folder, r, pkgName),
		GoTest:       GoTest(folder, r, pkgName),
		GoVet:        GoVet(folder, r, pkgName),
		GolangCILint: GolangCILint(folder, r, pkgName),
		GoBench:      GoBench(folder, r, pkgName),
	}, err
}
