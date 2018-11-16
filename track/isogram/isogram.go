package isogram

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/track/isogram/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	addSpeedComment = getAddSpeedComment()

	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examRegexCompileInFunc,
	examToLowerUpper("ToLower"),
	examToLowerUpper("ToUpper"),
	examJustReturn,
	examNonExistingMapValue,
	examUniversalIsLetter,
	examIfContinueIsLetter,
	examZeroValueAssign,
}

func examZeroValueAssign(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeAssignStmt)
	for _, node := range nodes {
		ident := node.ChildByNodeType(astrav.NodeTypeIdent)
		b := node.ChildByNodeType(astrav.NodeTypeBasicLit)
		if b == nil || ident == nil {
			continue
		}
		bLit := b.(*astrav.BasicLit)

		switch ident.ValueType().String() {
		case "string":
			if bLit.Value == `""` {
				r.AppendImprovement(tpl.ZeroValueAssign)
			}
		case "bool":
			if bLit.Value == `false` {
				r.AppendImprovement(tpl.ZeroValueAssign)
			}
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			if bLit.Value == `0` {
				r.AppendImprovement(tpl.ZeroValueAssign)
			}
		}
	}
}

func examIfContinueIsLetter(pkg *astrav.Package, r *extypes.Response) {
	node := pkg.FindFirstByName("IsLetter")
	if node == nil {
		return
	}

	ifNode := node.NextParentByType(astrav.NodeTypeIfStmt)
	if ifNode == nil {
		return
	}
	contNodes := ifNode.FindByToken(token.CONTINUE)
	retNodes := ifNode.FindByNodeType(astrav.NodeTypeReturnStmt)
	ifNodes := ifNode.FindByNodeType(astrav.NodeTypeIfStmt)

	if len(contNodes)+len(retNodes) == 0 {
		r.AppendImprovement(tpl.IfContinue)
	} else if len(contNodes) == 0 && len(ifNodes) != 0 {
		r.AppendImprovement(tpl.IfContinue)
	}
}

func examUniversalIsLetter(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeBasicLit)
	for _, node := range nodes {
		bLit := node.(*astrav.BasicLit)
		if strings.Contains(bLit.Value, "-") {
			r.AppendImprovement(tpl.UniversalIsLetter)
		}
		if strings.Contains(bLit.Value, "a") && bLit.IsValueType("rune") {
			r.AppendImprovement(tpl.UniversalIsLetter)
		}
	}
}

func examNonExistingMapValue(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeIndexExpr)
	for _, node := range nodes {
		parent := node.Parent()
		if !parent.IsNodeType(astrav.NodeTypeAssignStmt) {
			continue
		}

		assign := parent.ChildrenByNodeType(astrav.NodeTypeIdent)
		if len(assign) < 2 {
			continue
		}

		if assign[0].IsValueType("bool") {
			r.AppendComment(tpl.NonExistingMapValue)
		}
	}
}

func examJustReturn(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByToken(token.BREAK)
	if len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		ifStmt := node.NextParentByType(astrav.NodeTypeIfStmt)
		if ifStmt == nil {
			continue
		}
		boolVar := ifStmt.FindByValueType("bool")
		if boolVar == nil {
			continue
		}

		r.AppendImprovement(tpl.JustReturn)
		break
	}
}

func examRegexCompileInFunc(pkg *astrav.Package, r *extypes.Response) {
	main := pkg.FindFirstByName("IsIsogram")
	regComp := pkg.FindFirstByName("Compile")
	if regComp != nil && main.Contains(regComp) {
		r.AppendTodo(tpl.RegexInFunc)
		r.AppendTodo(tpl.MustCompile)
	}
	if regComp != nil {
		r.AppendTodo(tpl.IsLetter)
	}

	regComp = pkg.FindFirstByName("MustCompile")
	if regComp != nil && main.Contains(regComp) {
		r.AppendTodo(tpl.RegexInFunc)
	}
	if regComp != nil {
		r.AppendTodo(tpl.IsLetter)
	}
}

func examToLowerUpper(fnName string) extypes.SuggestionFunc {
	return func(pkg *astrav.Package, r *extypes.Response) {
		fns := pkg.FindByName(fnName)
		for _, fn := range fns {
			if f, ok := fn.(*astrav.SelectorExpr); !ok ||
				f.X.(*ast.Ident).Name == "unicode" {
				continue
			}
			addSpeedComment(r)

			if fn.NextParentByType(astrav.NodeTypeBlockStmt).IsContainedByType(astrav.NodeTypeRangeStmt) {
				r.AppendImprovement(tpl.UnicodeLoop.Format(fnName))
			} else {
				r.AppendImprovement(tpl.Unicode.Format(fnName))
			}
		}
	}
}

var addSpeedComment func(r *extypes.Response)

func getAddSpeedComment() func(r *extypes.Response) {
	var speedCommentAdded bool
	return func(r *extypes.Response) {
		if speedCommentAdded {
			return
		}
		speedCommentAdded = true
		r.AppendOutro(gtpl.Benchmarking)
	}
}
