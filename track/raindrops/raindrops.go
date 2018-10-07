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
	examStringsBuilder,
	examTooManyConcats,
}

func examManyLoops(pkg *astrav.Package, r *extypes.Response) {
	var count int
	count += len(pkg.FindByNodeType(astrav.NodeTypeForStmt))
	count += len(pkg.FindByNodeType(astrav.NodeTypeRangeStmt))

	if 1 < count {
		r.AppendTodo(tpl.ManyLoops)
	}
}

func examTooManyConcats(pkg *astrav.Package, r *extypes.Response) {
	var count int
	ifs := pkg.FindByNodeType(astrav.NodeTypeIfStmt)
	if len(ifs) < 4 {
		return
	}
	for _, node := range ifs {
		assign := node.FindFirstByNodeType(astrav.NodeTypeAssignStmt)
		if assign == nil {
			continue
		}
		if assign.(*astrav.AssignStmt).Tok.String() == "+=" {
			count++
		}
	}
	if 2 < count {
		r.AppendImprovement(tpl.ConcatNotNeeded)
	}
}

func examStringsBuilder(pkg *astrav.Package, r *extypes.Response) {
	builder := pkg.FindByName("Builder")
	if builder != nil {
		r.AppendImprovement(tpl.StringsBuilder)
	}
}
