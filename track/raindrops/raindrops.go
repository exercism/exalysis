package raindrops

import (
	"strings"

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
	examLoopMap,
	examManyLoops,
	examStringsBuilder,
	examItoa,
	examExtensiveForLoop,
	examPlusEqual,
	examTooManyConcats,
}

func examLoopMap(pkg *astrav.Package, r *extypes.Response) {
	loops := pkg.FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, l := range loops {
		loop := l.(*astrav.RangeStmt)

		ident, ok := loop.X().(*astrav.Ident)
		if !ok {
			continue
		}
		if strings.HasPrefix(ident.ValueType().String(), "map") {
			r.AppendImprovement(tpl.LoopMap)
			return
		}
	}
}

func examExtensiveForLoop(pkg *astrav.Package, r *extypes.Response) {
	f := pkg.FindFirstByName("Convert")
	params := f.(*astrav.FuncDecl).Params()
	if len(params.List) == 0 || len(params.List[0].Names) == 0 {
		return
	}
	paramName := params.List[0].Names[0].Name

	loops := pkg.FindByNodeType(astrav.NodeTypeForStmt)
	for _, l := range loops {
		loop := l.(*astrav.ForStmt)

		// check if loop goes up to input var
		if loop.Cond().FindFirstByName(paramName) != nil {
			r.AppendImprovement(tpl.ExtensiveFor)
			return
		}

		// if using a basiclit it should be 7 or 8 depending on the operator
		basicLit := loop.Cond().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if basicLit != nil {
			val := basicLit.(*astrav.BasicLit).Value
			if val != "7" && val != "8" {
				r.AppendImprovement(tpl.ExtensiveFor)
				return
			}
		}

		// check if loop starts with 3
		basicLit = loop.Init().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if basicLit != nil {
			if basicLit.(*astrav.BasicLit).Value != "3" {
				r.AppendImprovement(tpl.ExtensiveFor)
				return
			}
		}

		// check if loop uses steps of 2: 3, 5, 7
		if loop.Post().IsNodeType(astrav.NodeTypeIncDecStmt) {
			r.AppendImprovement(tpl.ExtensiveFor)
			return
		}
		if loop.Post().IsNodeType(astrav.NodeTypeAssignStmt) {
			basicLit := loop.Post().FindFirstByNodeType(astrav.NodeTypeBasicLit)
			if basicLit != nil {
				if basicLit.(*astrav.BasicLit).Value != "2" {
					r.AppendImprovement(tpl.ExtensiveFor)
					return
				}
			}
		}
	}
}

func examItoa(pkg *astrav.Package, r *extypes.Response) {
	itoa := pkg.FindFirstByName("Itoa")
	if itoa != nil {
		if itoa.IsNodeType(astrav.NodeTypeSelectorExpr) &&
			itoa.(*astrav.SelectorExpr).PackageName().Name == "strconv" {
			return
		}
	}

	r.AppendImprovement(tpl.Itoa)
}

func examPlusEqual(pkg *astrav.Package, r *extypes.Response) {
	assigns := pkg.FindByNodeType(astrav.NodeTypeAssignStmt)
	for _, assign := range assigns {
		token := assign.(*astrav.AssignStmt).Tok.String()
		if token != "=" {
			continue
		}

		binExpr := assign.FindFirstByNodeType(astrav.NodeTypeBinaryExpr)
		if binExpr == nil {
			continue
		}

		if binExpr.(*astrav.BinaryExpr).Op.String() == "+" {
			r.AppendComment(tpl.PlusEqual)
			return
		}
	}
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
