package exam

import "github.com/tehsphinx/exalysis/extypes"

//Result contains the result if running all examinations
type Result struct {
	GoLint bool
	GoFmt  bool
}

//All runs all examinations contained in this package and returns the result
func All(files map[string][]byte, r *extypes.Response, pkgName string) Result {
	return Result{
		GoLint: GoLint(files, r, pkgName),
		GoFmt:  GoFmt(files, r, pkgName),
	}
}
