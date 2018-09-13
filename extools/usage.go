package extools

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/loader"
)

func GetUsageInScope(name string, scope *types.Scope, pkg *loader.PackageInfo) (*ast.Ident, types.Object) {
	for k, v := range pkg.Uses {
		if k.Name == name {
			if scope.Contains(k.Pos()) {
				return k, v
			}
		}
	}
	return nil, nil
}

//GetUsage returns a usage by name
func GetUsage(name string, pkg *loader.PackageInfo) (*ast.Ident, types.Object) {
	for k, v := range pkg.Uses {
		if k.Name == name {
			return k, v
		}
	}
	return nil, nil
}

//GetUsageFunc returns a function usage from given package
func GetUsageFunc(funcName string, pkg *loader.PackageInfo) (*ast.Ident, *types.Func) {
	for k, v := range pkg.Uses {
		if k.Name == funcName {
			if fn, ok := v.(*types.Func); ok {
				return k, fn
			}
		}
	}
	return nil, nil
}
