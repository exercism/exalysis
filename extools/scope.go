package extools

import (
	"go/ast"

	"golang.org/x/tools/go/loader"
)

//EnclosingRangeStmt returns and enclosing for..range loop
func EnclosingRangeStmt(node ast.Node, pkg *loader.PackageInfo) *ast.RangeStmt {
	for k, v := range pkg.Scopes {
		if v.Contains(node.Pos()) {
			switch o := k.(type) {
			case *ast.RangeStmt:
				return o
			}
		}
	}
	return nil
}
