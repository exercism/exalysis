package diffsquares

import (
	"github.com/tehsphinx/astrav"
	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/track/diffsquares/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examLoops,
	examMathPow,
	examCalcRangeCondition,
	examDry,
	examBasicFloat,
}

func examBasicFloat(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByName("float64")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeCallExpr) {
			continue
		}
		if len(node.Children()) != 2 {
			continue
		}

		child := node.ChildByNodeType(astrav.NodeTypeBasicLit)
		if child == nil {
			continue
		}
		if child.IsValueType("float64") {
			r.AppendImprovement(tpl.BasicFloat64)
		}
	}
}

func examDry(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByName("Difference")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeFuncDecl) {
			continue
		}

		nssq := node.FindFirstByName("SumOfSquares")
		if nssq == nil {
			r.AppendImprovement(tpl.Dry)
		}

		nsqs := node.FindFirstByName("SquareOfSum")
		if nsqs == nil {
			r.AppendImprovement(tpl.Dry)
		}
	}
}

func examMathPow(pkg *astrav.Package, r *extypes.Response) {
	node := pkg.FindFirstByName("Pow")
	if node != nil {
		r.AppendImprovement(tpl.MathPow)
	}
}

func examLoops(pkg *astrav.Package, r *extypes.Response) {
	funcs := pkg.FindByNodeType(astrav.NodeTypeFuncDecl)
	for _, f := range funcs {
		nodes := f.FindNodeTypeInCallTree(astrav.NodeTypeForStmt)

		switch f.(*astrav.FuncDecl).Name.Name {
		case "SquareOfSum":
			if len(nodes) != 0 {
				r.AppendImprovement(tpl.SquareSumLoop)
			}
		case "SumOfSquares":
			if len(nodes) != 0 {
				r.AppendImprovement(tpl.SumSquareLoop)
			}
		}
	}
}

func examCalcRangeCondition(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeForStmt)
	for _, node := range nodes {
		cond := node.(*astrav.ForStmt).Cond()
		if cond == nil {
			continue
		}
		binExpr := cond.FindByNodeType(astrav.NodeTypeBinaryExpr)
		if len(binExpr) != 0 {
			r.AppendImprovement(tpl.CalcRangeCondition.Format(binExpr[0].GetSourceString()))
		}
	}
}
