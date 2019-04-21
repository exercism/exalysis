package exam

import (
	"os"

	"github.com/exercism/exalysis/extypes"
	"github.com/tehsphinx/astrav"
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
func All(folder *astrav.Folder, r *extypes.Response, pkgName string, analyze bool) (res *Result, err error) {
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
		GoTest: test,
	}

	resp := r
	if analyze {
		resp = &extypes.Response{}
	}
	res.GoLint = GoLint(folder, resp, pkgName)
	res.GoFmt = GoFmt(folder, resp, pkgName)
	res.GoVet = GoVet(folder, r, pkgName, errTest != nil)
	res.GolangCILint = GolangCILint(folder, r, pkgName)
	res.GoBench = GoBench(folder, r, pkgName, errTest != nil)

	return res, err
}
