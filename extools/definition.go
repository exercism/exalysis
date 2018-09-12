package extools

import (
	"go/ast"
	"go/types"

	"github.com/tehsphinx/dbg"
	"golang.org/x/tools/go/loader"
)

func GetFunc(funcName string, pkg *loader.PackageInfo) (*ast.Ident, *types.Func) {
	for k, v := range pkg.Defs {
		if k.String() == funcName {
			if fn, ok := v.(*types.Func); ok {
				return k, fn
			}
		}
	}
	return nil, nil
}

func GetDefinition(name string, pkg *loader.PackageInfo) (*ast.Ident, types.Object) {
	for k, v := range pkg.Defs {
		if k.String() == name {
			dbg.Green(v.Name())
			dbg.Magenta(v.Type())
			dbg.Red(v.Parent())
			return k, v
		}
	}
	return nil, nil
}
