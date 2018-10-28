package scrabble

import (
	"go/ast"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/track/scrabble/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	testGoRoutine,
	testRegex,
	testMapInFunc,
	testFlattenMap,
	testMapRuneInt,
	testToLowerUpper("ToLower"),
	testToLowerUpper("ToUpper"),
	testTrySwitch,
	testIfElseToSwitch,
	testRuneLoop,
}

func testRegex(pkg *astrav.Package, r *extypes.Response) {
	rgx := pkg.FindFirstByName("Score").FindFirstByName("MustCompile")
	if rgx == nil {
		rgx = pkg.FindFirstByName("Score").FindFirstByName("Compile")
	}

	if rgx != nil &&
		rgx.(*astrav.SelectorExpr).X != nil &&
		rgx.(*astrav.SelectorExpr).X.(*ast.Ident).Name == "regexp" {

		lit := rgx.Parent().ChildByNodeType(astrav.NodeTypeBasicLit)
		if lit != nil && lit.IsValueType("string") {
			// is a static regex
			r.AppendTodo(tpl.Regex)
			r.AppendComment(tpl.Challenge)
			r.AppendOutro(gtpl.Regex)
			addSpeedComment(r)
		}
	}
}

func testToLowerUpper(fnName string) extypes.SuggestionFunc {
	return func(pkg *astrav.Package, r *extypes.Response) {
		if r.HasSuggestion(tpl.Regex) {
			return
		}

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

func testFlattenMap(pkg *astrav.Package, r *extypes.Response) {
	loopCount := len(pkg.FindByNodeType(astrav.NodeTypeForStmt))
	loopCount += len(pkg.FindByNodeType(astrav.NodeTypeRangeStmt))
	if 1 < loopCount {
		addSpeedComment(r)
		r.AppendImprovement(tpl.FlattenMap)
	}
}

func testMapRuneInt(pkg *astrav.Package, r *extypes.Response) {
	if r.HasSuggestion(tpl.FlattenMap) {
		return
	}
	if len(pkg.FindByValueType("map[string]int")) != 0 {
		addSpeedComment(r)
		r.AppendImprovement(tpl.MapRune.Format("map[string]int"))
		return
	}
	if len(pkg.FindByValueType("map[int]string")) != 0 {
		addSpeedComment(r)
		r.AppendImprovement(tpl.MapRune.Format("map[int]string"))
		return
	}
}

func testTrySwitch(pkg *astrav.Package, r *extypes.Response) {
	if r.HasSuggestion(tpl.FlattenMap) {
		return
	}
	if len(pkg.FindByValueType("map[rune]int")) != 0 {
		addSpeedComment(r)
		r.AppendComment(tpl.TrySwitch)
		return
	}
	if len(pkg.FindByValueType("map[string]int")) != 0 {
		addSpeedComment(r)
		r.AppendComment(tpl.TrySwitch)
		return
	}
	if len(pkg.FindByValueType("map[int]string")) != 0 {
		addSpeedComment(r)
		r.AppendComment(tpl.TrySwitch)
		return
	}
}

func testRuneLoop(pkg *astrav.Package, r *extypes.Response) {
	ranges := pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		l := rng.(*astrav.RangeStmt)
		if l.Value() == nil {
			if l.Key() != nil {
				r.AppendImprovement(tpl.LoopRuneNotByte)
			}
		} else {
			var isByte bool
			for _, ident := range rng.FindIdentByName(l.Value().(*astrav.Ident).Name) {
				if ident.IsValueType("byte") {
					isByte = true
				}
			}
			if isByte {
				r.AppendImprovement(tpl.LoopRuneNotByte)
			}
		}

		if rng.FindByName("string") != nil &&
			!r.HasSuggestion(tpl.MapRune) &&
			!r.HasSuggestion(tpl.FlattenMap) {

			r.AppendImprovement(tpl.TypeConversion)
			return
		}
	}
}

func testGoRoutine(pkg *astrav.Package, r *extypes.Response) {
	goStmts := pkg.FindByNodeType(astrav.NodeTypeGoStmt)
	if len(goStmts) != 0 {
		addSpeedComment(r)
		r.AppendTodo(tpl.GoRoutines)
	}
}

func testIfElseToSwitch(pkg *astrav.Package, r *extypes.Response) {
	ranges := pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		ifs := rng.FindByNodeType(astrav.NodeTypeIfStmt)
		if 5 < len(ifs) {
			r.AppendTodo(tpl.IfsToSwitch)
			return
		}
	}
}

func testMapInFunc(pkg *astrav.Package, r *extypes.Response) {
	fn := pkg.FindFirstByName("Score")
	maps := fn.FindMaps()

	var hasMapDef bool
	for _, m := range maps {
		if !m.IsNodeType(astrav.NodeTypeIdent) {
			hasMapDef = true
		}
	}
	if hasMapDef {
		addSpeedComment(r)
		r.AppendTodo(tpl.MoveMap)
	}
}

var speedCommentAdded bool

func addSpeedComment(r *extypes.Response) {
	if !speedCommentAdded {
		speedCommentAdded = true
		r.AppendOutro(gtpl.Benchmarking)
	}
}
