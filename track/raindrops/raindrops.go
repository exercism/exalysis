package raindrops

import (
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/raindrops/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examManyLoops,
}


func examManyLoops(pkg *astrav.Package, r *extypes.Response) {
	var count int
	count += len(pkg.FindByNodeType(astrav.NodeTypeForStmt))
	count += len(pkg.FindByNodeType(astrav.NodeTypeRangeStmt))

	if 1 < count {
		r.AppendTodo(tpl.ManyLoops)
	}

}
