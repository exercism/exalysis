package scrabble

import (
	"go/ast"

	"github.com/tehsphinx/dbg"
	"github.com/tehsphinx/exalysis/suggestion/types"
	"honnef.co/go/tools/ssa"
)

type testFunc func(sugg []string) []string

func NewScrabble(program *ssa.Program, pkg *ssa.Package) types.Suggester {
	s := &Scrabble{
		program: program,
		pkg:     pkg,
	}
	s.tests = []testFunc{
		s.testToLower,
		//s.testToUpper,
		//s.testMapRuneInt,
	}

	return s

}

//Scrabble implements the suggester for the scrabble exercise
type Scrabble struct {
	program *ssa.Program
	pkg     *ssa.Package
	tests   []testFunc
}

//Suggest builds suggestions for the exercise solution
func (s *Scrabble) Suggest() []string {
	//ast.Print(s.program, s.pkg.Files["scrabble.go"])
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
	toLower = "- you could have a look at `unicode.ToLower` inside the for loop instead of `strings.ToUpper` " +
		"before the loop to increase speed"
	toUpper = "- you could have a look at `unicode.ToUpper` inside the for loop instead of `strings.ToUpper` " +
		"before the loop to increase speed"
	mapRune   = "- you could use a `map[rune]int` for direct lookup without type conversion instead of `map[string]int`"
	trySwitch = "- if you are up for it using a `switch` instead of a `map` will increase speed significantly"
)

var (
	speedCommentAdded bool
)

func (s *Scrabble) testToLower(sugg []string) []string {
	//entryFunc := GetFuncDecl("Score", s.pkg)
	//dbg.Red(entryFunc)

	//ast.Inspect(s.pkg, func(node ast.Node) bool {
	//	switch n := node.(type) {
	//	case *ast.FuncDecl:
	//		dbg.Green("funcDecl", n)
	//	case *ast.FuncLit:
	//		dbg.Green("funcLit", n)
	//	case *ast.CallExpr:
	//		dbg.Blue("call", n, n.Fun)
	//	case *ast.Ident:
	//		dbg.Red("ident", n, n.Name)
	//	}
	//	return true
	//})

	//fn, ok := extools2.GetUsageFunc("ToLower", local)
	//if !ok {
	//	return sugg
	//}
	//if fn.Pkg().Name() == "unicode" {
	//	return sugg
	//}

	sugg = addSpeedComment(sugg)
	return append(sugg, toLower)
}

//func (s *Scrabble) testToUpper(fset *token.FileSet, pkg *ast.Package, sugg []string) []string {
//	fn, ok := extools2.GetUsageFunc("ToUpper", local)
//	if !ok {
//		return sugg
//	}
//	if fn.Pkg().Name() == "unicode" {
//		return sugg
//	}
//
//	extools2.GetDefinition("Score", local)
//	//usage := extools.GetUsage("ToUpper", local)
//	//dbg.Green(usage.Pos())
//
//	sugg = addSpeedComment(sugg)
//	return append(sugg, toUpper)
//}
//
//func (s *Scrabble) testMapRuneInt(fset *token.FileSet, pkg *ast.Package, sugg []string) []string {
//	for _, t := range local.Types {
//		switch t.Type.String() {
//		case "map[rune]int":
//			sugg = addSpeedComment(sugg)
//			sugg = append(sugg, trySwitch)
//			return sugg
//		case "map[string]int":
//			sugg = addSpeedComment(sugg)
//			sugg = append(sugg, mapRune)
//			sugg = append(sugg, trySwitch)
//			return sugg
//		}
//	}
//	return sugg
//}

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
