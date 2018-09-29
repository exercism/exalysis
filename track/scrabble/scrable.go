package scrabble

import (
	"go/ast"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/track/scrabble/tpl"
	"golang.org/x/tools/go/loader"
	"honnef.co/go/tools/ssa"
)

type testFunc func(r *extypes.Response)

//NewScrabble creates a new suggester for the scrabble exercise
func NewScrabble(pkg *astrav.Package) extypes.Suggester {
	s := &Scrabble{
		pkg: pkg,
	}
	s.tests = []testFunc{
		s.testGoRoutine,
		s.testMapInFunc,
		s.testToLowerUpper("ToLower"),
		s.testToLowerUpper("ToUpper"),
		s.testMapRuneInt,
		s.testIfElseToSwitch,
		s.testRuneLoop,
	}

	return s
}

//Scrabble implements the suggester for the scrabble exercise
type Scrabble struct {
	pkg    *astrav.Package
	prog   *loader.Program
	lPkg   *loader.PackageInfo
	ssaPkg *ssa.Package
	tests  []testFunc
}

//Suggest builds suggestions for the exercise solution
func (s *Scrabble) Suggest(r *extypes.Response) {
	for _, tf := range s.tests {
		tf(r)
	}
}

func (s *Scrabble) testToLowerUpper(fnName string) testFunc {
	return func(r *extypes.Response) {
		fns := s.pkg.FindByName(fnName)
		for _, fn := range fns {
			if f, ok := fn.(*astrav.SelectorExpr); !ok ||
				f.X.(*ast.Ident).Name == "unicode" {
				continue
			}
			addSpeedComment(r)

			if fn.IsContainedByType(astrav.NodeTypeRangeStmt) {
				r.AppendImprovement(tpl.UnicodeLoop.Format(fnName))
			} else {
				r.AppendImprovement(tpl.Unicode.Format(fnName))
			}
		}
	}
}

func (s *Scrabble) testMapRuneInt(r *extypes.Response) {
	if len(s.pkg.FindByValueType("map[rune]int")) != 0 {
		addSpeedComment(r)
		r.AppendImprovement(tpl.TrySwitch)
		return
	}
	if len(s.pkg.FindByValueType("map[string]int")) != 0 {
		addSpeedComment(r)
		r.AppendImprovement(tpl.MapRune.Format("map[string]int"))
		r.AppendImprovement(tpl.TrySwitch)
		return
	}
	if len(s.pkg.FindByValueType("map[int]string")) != 0 {
		addSpeedComment(r)
		r.AppendImprovement(tpl.MapRune.Format("map[int]string"))
		r.AppendImprovement(tpl.TrySwitch)
		return
	}
}

func (s *Scrabble) testRuneLoop(r *extypes.Response) {
	ranges := s.pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		l := rng.(*astrav.RangeStmt)
		if l.Value == nil {
			if l.Key != nil {
				r.AppendImprovement(tpl.LoopRuneNotByte)
			}
		} else {
			var isByte bool
			for _, ident := range rng.FindIdentByName(l.Value.(*ast.Ident).Name) {
				if ident.IsValueType("byte") {
					isByte = true
				}
			}
			if isByte {
				r.AppendImprovement(tpl.LoopRuneNotByte)
			}
		}

		if rng.FindByName("string") != nil {
			r.AppendImprovement(tpl.TypeConversion)
			return
		}
	}
}

func (s *Scrabble) testGoRoutine(r *extypes.Response) {
	goStmts := s.pkg.FindByNodeType(astrav.NodeTypeGoStmt)
	if len(goStmts) != 0 {
		addSpeedComment(r)
		r.AppendTodo(tpl.GoRoutines)
	}
}

func (s *Scrabble) testIfElseToSwitch(r *extypes.Response) {
	ranges := s.pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		ifs := rng.FindByNodeType(astrav.NodeTypeIfStmt)
		if 5 < len(ifs) {
			r.AppendTodo(tpl.IfsToSwitch)
			return
		}
	}
}

func (s *Scrabble) testMapInFunc(r *extypes.Response) {
	fn := s.pkg.FindFirstByName("Score")
	maps := fn.FindByValueType("map[rune]int")
	if len(maps) == 0 {
		maps = fn.FindByValueType("map[string]int")
	}
	if len(maps) == 0 {
		maps = fn.FindByValueType("map[int]string")
	}
	if len(maps) != 0 {
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
