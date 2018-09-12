package extools

import (
	"go/ast"
	"go/types"
)

//Contains checks if a node contains another node
func Contains(parent, child ast.Node) bool {
	return types.NewScope(nil, parent.Pos(), parent.End(), "").
		Contains(child.Pos())
}
