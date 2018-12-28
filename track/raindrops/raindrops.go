package raindrops

import (
	"regexp"
	"strings"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/raindrops/tpl"
)

// Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examLoopMap,
	examManyLoops,
	examStringsBuilder,
	examBytesBuffer,
	examItoa,
	examExtensiveForLoop,
	examPlusEqual,
	examTooManyConcats,
	examRemoveExtraBool,
	examFmtPrintf,
	examAllCases,
	examCompareLen,
}

func examAllCases(pkg *astrav.Package, r *extypes.Response) {
	lits := pkg.FindByNodeType(astrav.NodeTypeBasicLit)
	for _, lit := range lits {
		if lit.ValueType() == nil || lit.ValueType().String() != "string" {
			continue
		}

		occs := fmtPrintfRegex.FindAllString(lit.(*astrav.BasicLit).Value, -1)
		if 1 < len(occs) {
			r.AppendTodo(tpl.AllCases)
			return
		}
	}

	rets := pkg.FindByNodeType(astrav.NodeTypeReturnStmt)
	if 6 < len(rets) {
		r.AppendTodo(tpl.AllCases)
	}
}

var fmtPrintfRegex = regexp.MustCompile(`Pl.ng`)

func examFmtPrintf(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByName("Sprintf")
	for _, node := range nodes {
		lits := node.Parent().FindByNodeType(astrav.NodeTypeBasicLit)
		for _, lit := range lits {
			if lit.ValueType().String() != "string" {
				continue
			}
			if fmtPrintfRegex.MatchString(lit.(*astrav.BasicLit).Value) {
				r.AppendImprovement(tpl.FmtPrint.Format("fmt.Sprintf"))
			}
		}
	}
}

func examRemoveExtraBool(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByValueType("bool")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeIdent) {
			continue
		}
		if ifParent := node.NextParentByType(astrav.NodeTypeIfStmt); ifParent == nil {
			continue
		} else if ifParent.Level()+2 < node.Level() {
			continue
		}

		name := node.(*astrav.Ident).Name
		r.AppendImprovement(tpl.RemoveExtraBool.Format(name))
		break
	}
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
			r.AppendTodo(tpl.LoopMap)
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
		if loop.Init() == nil {
			// Probably a condition-only loop
			return
		}
		basicLit = loop.Init().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if basicLit != nil {
			if basicLit.(*astrav.BasicLit).Value != "3" {
				r.AppendImprovement(tpl.ExtensiveFor)
				return
			}
		}

		// check if loop uses steps of 2: 3, 5, 7
		if loop.Post() == nil {
			// Probably a condition-only loop
			return
		}
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
		r.AppendImprovement(tpl.StringsBuilder.Format("strings.Builder"))
	}
}

func examBytesBuffer(pkg *astrav.Package, r *extypes.Response) {
	buffer := pkg.FindByName("Buffer")
	if buffer != nil {
		r.AppendImprovement(tpl.StringsBuilder.Format("bytes.Buffer"))
	}
}

func examCompareLen(pkg *astrav.Package, r *extypes.Response) {
	// Find all function calls
	for _, call := range pkg.FindByNodeType(astrav.NodeTypeCallExpr) {
		nodeName := call.(*astrav.CallExpr).NodeName()
		if nodeName == nil {
			// not a builtin, so not 'len()'
			continue
		}
		if nodeName.NodeName().Name != "len" {
			// not 'len()'
			continue
		}
		identName := call.(*astrav.CallExpr).Children()[1].(*astrav.Ident).Name
		if pkg.FindFirstIdentByName(identName).IsValueType("string") {
			r.AppendImprovement(tpl.CompareEmptyString)
		}
	}
}
