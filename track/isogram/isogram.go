package isogram

import (
	"go/token"
	"strings"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/track/isogram/tpl"
	"github.com/tehsphinx/astrav"
)

// Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	addSpeedComment = getAddSpeedComment()

	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examRegexCompileInFunc,
	examToLowerUpper("strings.ToLower"),
	examToLowerUpper("strings.ToUpper"),
	examJustReturn,
	examNonExistingMapValue,
	examUniversalIsLetter,
	examIfContinueIsLetter,
	examZeroValueAssign,
	examTwoLoops,
}

func examTwoLoops(pkg *astrav.Package, r *extypes.Response) {
	loops := pkg.FindByNodeType(astrav.NodeTypeRangeStmt)
	loops = append(loops, pkg.FindByNodeType(astrav.NodeTypeForStmt)...)

	if 1 < len(loops) {
		r.AppendImprovementTpl(tpl.TwoLoops)
	}
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
				r.AppendImprovementTpl(tpl.ZeroValueAssign)
			}
		case "bool":
			if bLit.Value == `false` {
				r.AppendImprovementTpl(tpl.ZeroValueAssign)
			}
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			if bLit.Value == `0` {
				r.AppendImprovementTpl(tpl.ZeroValueAssign)
			}
		}
	}
}

func examIfContinueIsLetter(pkg *astrav.Package, r *extypes.Response) {
	node := pkg.FindFirstByName("unicode.IsLetter")
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
		r.AppendImprovementTpl(tpl.IfContinue)
	} else if len(contNodes) == 0 && len(ifNodes) != 0 {
		r.AppendImprovementTpl(tpl.IfContinue)
	}
}

func examUniversalIsLetter(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeBasicLit)
	for _, node := range nodes {
		bLit := node.(*astrav.BasicLit)
		if strings.Contains(bLit.Value, "-") {
			r.AppendImprovementTpl(tpl.UniversalIsLetter)
		}
		if strings.Contains(bLit.Value, "a") && bLit.IsValueType("rune") {
			r.AppendImprovementTpl(tpl.UniversalIsLetter)
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
			r.AppendCommentTpl(tpl.NonExistingMapValue)
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

		r.AppendImprovementTpl(tpl.JustReturn)
		break
	}
}

func examRegexCompileInFunc(pkg *astrav.Package, r *extypes.Response) {
	main := pkg.FindFirstByName("IsIsogram")
	regComp := pkg.FindFirstByName("regexp.Compile")
	if regComp != nil && main.Contains(regComp) {
		r.AppendTodoTpl(tpl.RegexInFunc)
		r.AppendTodoTpl(tpl.MustCompile)
	}
	if regComp != nil {
		r.AppendTodoTpl(tpl.IsLetter)
	}

	regComp = pkg.FindFirstByName("regexp.MustCompile")
	if regComp != nil && main.Contains(regComp) {
		r.AppendTodoTpl(tpl.RegexInFunc)
	}
	if regComp != nil {
		r.AppendTodoTpl(tpl.IsLetter)
	}
}

func examToLowerUpper(fnName string) extypes.SuggestionFunc {
	return func(pkg *astrav.Package, r *extypes.Response) {
		fns := pkg.FindByName(fnName)
		for _, fn := range fns {
			if _, ok := fn.(*astrav.SelectorExpr); !ok {
				continue
			}
			addSpeedComment(r)

			if fn.NextParentByType(astrav.NodeTypeBlockStmt).IsContainedByType(astrav.NodeTypeRangeStmt) {
				r.AppendImprovementTpl(tpl.UnicodeLoop.Format(fnName))
			} else {
				r.AppendImprovementTpl(tpl.Unicode.Format(fnName))
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
