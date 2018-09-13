package extools

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/loader"
)

//NewScope creates a scope for any ast.Node
func NewScope(node ast.Node) *types.Scope {
	return types.NewScope(nil, node.Pos(), node.End(), "")
}

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

//RangeStmtsInScope finds all range statements in given scope
func RangeStmtsInScope(scope *types.Scope, pkg *loader.PackageInfo) map[*ast.RangeStmt]*types.Scope {
	var rangeStmts = map[*ast.RangeStmt]*types.Scope{}
	for k, v := range pkg.Scopes {
		if scope.Contains(k.Pos()) {
			switch o := k.(type) {
			case *ast.RangeStmt:
				rangeStmts[o] = v
			}
		}
	}
	return rangeStmts
}
