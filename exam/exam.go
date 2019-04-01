package exam

import (
	"os"

	"github.com/tehsphinx/astrav"
	"github.com/exercism/exalysis/extypes"
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

	test, errTest := GoTest(folder, r, pkgName)
	res = &Result{
		GoLint:       GoLint(folder, r, pkgName),
		GoFmt:        GoFmt(folder, r, pkgName),
		GoTest:       test,
		GoVet:        GoVet(folder, r, pkgName, errTest != nil),
		GolangCILint: GolangCILint(folder, r, pkgName),
		GoBench:      GoBench(folder, r, pkgName, errTest != nil),
	}
	return res, err
}
