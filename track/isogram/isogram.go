package isogram

import (
	"go/ast"
	"go/token"

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

		r.AppendComment(tpl.NonExistingMapValue)
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
		r.AppendTodo(tpl.IsLetter)
		return
	}

	regComp = pkg.FindFirstByName("MustCompile")
	if regComp != nil && main.Contains(regComp) {
		r.AppendTodo(tpl.RegexInFunc)
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
