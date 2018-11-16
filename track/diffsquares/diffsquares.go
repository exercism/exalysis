package diffsquares

import (
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/diffsquares/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examOneLoop,
	examCalcRangeCondition,
}

func examOneLoop(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeForStmt)
	if 1 < len(nodes) {
		r.AppendImprovement(tpl.OneLoop)
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
