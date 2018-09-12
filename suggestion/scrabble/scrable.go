package scrabble

import (
	"fmt"
	"go/ast"

	"github.com/tehsphinx/dbg"
	"github.com/tehsphinx/exalysis/extools"
	"github.com/tehsphinx/exalysis/suggestion/types"
	"golang.org/x/tools/go/loader"
	"honnef.co/go/tools/ssa"
)

type testFunc func(sugg []string) []string

//NewScrabble creates a new suggester for the scrabble exercise
func NewScrabble(program *loader.Program, pkg *ssa.Package) types.Suggester {
	pkg.Build()
	s := &Scrabble{
		pkg:  pkg,
		prog: program,
		lPkg: program.Package("."),
	}
	s.tests = []testFunc{
		s.testToLowerUpper("ToLower"),
		s.testToLowerUpper("ToUpper"),
		s.testMapRuneInt,
	}

	return s

}

//Scrabble implements the suggester for the scrabble exercise
type Scrabble struct {
	prog  *loader.Program
	lPkg  *loader.PackageInfo
	pkg   *ssa.Package
	tests []testFunc
}

//Suggest builds suggestions for the exercise solution
func (s *Scrabble) Suggest() []string {
	//ast.Print(s.pkg, s.pkg.Files["scrabble.go"])
	//extools2.PrintAST(local)

	var sugg []string
	for _, tf := range s.tests {
		sugg = tf(sugg)
	}
	return sugg
}

const (
	speedComment = "- this exercise is a lot about speed and memory usage. If you haven't done so already " +
		"use `go test -v --bench . --benchmem` to benchmark your solution. Do that before and after every " +
		"to see if the change improved the speed or not. Also note the amount of memory allocations which " +
		"cost time. This exercise can be solved with 0 allocs/op."
	toLowerUpper = "- you could have a look at `unicode.%[1]s` inside the for loop instead of `strings.%[1]s` " +
		"before the loop to increase speed"
	toLowerUpperInLoop = "- you could have a look at `unicode.%[1]s` to replace `strings.%[1]s` " +
		"in the loop to increase speed"
	mapRune   = "- you could use a `map[rune]int` for direct lookup without type conversion instead of `map[string]int`"
	trySwitch = "- if you are up for it using a `switch` instead of a `map` will increase speed significantly"
)

var (
	speedCommentAdded bool
)

func (s *Scrabble) testToLowerUpper(fnType string) testFunc {
	return func(sugg []string) []string {
		fnID, fn := extools.GetUsageFunc(fnType, s.lPkg)
		if fn == nil {
			return sugg
		}
		if fn.Pkg().Name() == "unicode" {
			return sugg
		}

		sugg = addSpeedComment(sugg)
		if forRange := extools.EnclosingRangeStmt(fnID, s.lPkg); forRange != nil {
			return append(sugg, fmt.Sprintf(toLowerUpperInLoop, fnType))
		}
		return append(sugg, fmt.Sprintf(toLowerUpper, fnType))
	}
}

//func (s *Scrabble) testToUpper(sugg []string) []string {
//	id, fn := extools.GetUsageFunc("ToUpper", s.lPkg)
//	if fn == nil {
//		return sugg
//	}
//	if fn.Pkg().Name() == "unicode" {
//		return sugg
//	}
//
//	dbg.Green(fn.Scope())
//	dbg.Green(fn.Parent())
//
//	//s.pkg.Pkg.Scope().Contains(nil)
//	entryFn := s.pkg.Func("Score")
//	dbg.Magenta(entryFn)
//	dbg.Magenta(entryFn.Blocks)
//	for _, block := range entryFn.Blocks {
//		dbg.Blue("BLOCK", block)
//		dbg.Blue(block.Comment)
//		dbg.Blue(block.Index)
//		dbg.Blue(block.Instrs)
//		dbg.Blue(block.Preds)
//		dbg.Blue(block.Succs)
//		dbg.Blue(block.String())
//		dbg.Blue(block.Dominees())
//		dbg.Blue(block.Idom())
//		dbg.Blue(block.Phis())
//	}
//
//	dbg.Magenta(entryFn.AnonFuncs)
//	dbg.Magenta(entryFn.FreeVars)
//	dbg.Magenta(entryFn.Locals)
//	dbg.Magenta(entryFn.Params)
//	dbg.Magenta(entryFn.Signature)
//	dbg.Magenta(entryFn.Synthetic)
//	dbg.Magenta(entryFn.DomPreorder())
//	dbg.Magenta(entryFn.Name())
//	dbg.Magenta(entryFn.Object())
//	dbg.Magenta(entryFn.Parent())
//	dbg.Magenta(entryFn.Pos())
//	dbg.Magenta(entryFn.Referrers())
//	dbg.Magenta(entryFn.String())
//	dbg.Magenta(entryFn.Syntax())
//	dbg.Magenta(entryFn.Syntax().Pos())
//	dbg.Magenta(entryFn.Syntax().End())
//	dbg.Magenta(entryFn.Token())
//	dbg.Magenta(entryFn.Type())
//	//encFn := ssa.EnclosingFunction(s.pkg, []ast.Node{id})
//	//dbg.Cyan(encFn)
//	//dbg.Cyan(encFn.String())
//
//	for k, v := range s.lPkg.Scopes {
//		if v.Contains(id.Pos()) {
//			dbg.Red(v)
//			//dbg.Green(v.Names())
//			dbg.Blue(k)
//			dbg.Green(reflect.TypeOf(k))
//			switch o := k.(type) {
//			case *ast.ForStmt:
//				dbg.Yellow(o)
//			case *ast.RangeStmt:
//				dbg.Yellow(o)
//			}
//		}
//	}
//
//	//floop := s.lPkg.Scopes
//	extools.PrintScopes(s.lPkg)
//
//	dbg.Blue(id.Name)
//	dbg.Blue(id.Obj)
//	dbg.Blue(id.NamePos)
//	dbg.Blue(id.String())
//	dbg.Blue(id.Pos())
//	dbg.Blue(id.End())
//	dbg.Blue(id.IsExported())
//
//	//extools.GetDefinition("Score", local)
//	//usage := extools.GetUsage("ToUpper", local)
//	//dbg.Green(usage.Pos())
//
//	//sugg = addSpeedComment(sugg)
//	return sugg
//}

func (s *Scrabble) testMapRuneInt(sugg []string) []string {
	for _, t := range s.lPkg.Types {
		switch t.Type.String() {
		case "map[rune]int":
			sugg = addSpeedComment(sugg)
			sugg = append(sugg, trySwitch)
			return sugg
		case "map[string]int":
			sugg = addSpeedComment(sugg)
			sugg = append(sugg, mapRune)
			sugg = append(sugg, trySwitch)
			return sugg
		}
	}
	return sugg
}

func addSpeedComment(sugg []string) []string {
	if !speedCommentAdded {
		speedCommentAdded = true
		return append(sugg, speedComment)
	}
	return sugg
}

//GetFuncDecl searches for a function declaration
func GetFuncDecl(name string, parent ast.Node) ast.Node {
	var decl ast.Node
	ast.Inspect(parent, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.FuncDecl:
			if n.Name.Name == name {
				decl = node
			}
		}
		return true
	})
	return decl
}

//FindFuncCall finds a function call
func FindFuncCall(name string, parent ast.Node) ast.Node {
	var decl ast.Node
	ast.Inspect(parent, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.CallExpr:
			dbg.Green(n)
		case *ast.FuncDecl:
			if n.Name.Name == name {
				decl = node
			}
		}
		return true
	})
	return decl
}
