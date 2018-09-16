package scrabble

import (
	"fmt"
	"go/ast"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/suggestion/defs"
	"golang.org/x/tools/go/loader"
	"honnef.co/go/tools/ssa"
)

type testFunc func(sugg []string) []string

//NewScrabble creates a new suggester for the scrabble exercise
func NewScrabble(pkg *astrav.Package) defs.Suggester {
	s := &Scrabble{
		pkg: pkg,
	}
	s.tests = []testFunc{
		s.testToLowerUpper("ToLower"),
		s.testToLowerUpper("ToUpper"),
		s.testMapRuneInt,
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
func (s *Scrabble) Suggest() []string {
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
	mapRune = "- you could use a `map[rune]int` for direct lookup without type conversion instead " +
		"of `map[string]int`"
	trySwitch  = "- if you are up for it using a `switch` instead of a `map` will increase speed significantly"
	typeSwitch = "- try to avoid type switches. You could work with type `rune` instead of `string` in " +
		"the `for` loop. A rune is created with e.g. 'A'."
	loopRuneNotByte = "- Iterating over a `string` will provice `rune`s which is a complete character and can " +
		"consist of **multiple** bytes. Try using runes instead of bytes."
)

var (
	speedCommentAdded bool
)

func (s *Scrabble) testToLowerUpper(fnName string) testFunc {
	return func(sugg []string) []string {
		fns := s.pkg.FindByName(fnName)
		for _, fn := range fns {
			if f, ok := fn.(*astrav.SelectorExpr); !ok ||
				f.X.(*ast.Ident).Name == "unicode" {
				continue
			}
			sugg = addSpeedComment(sugg)

			if fn.IsContainedByType(astrav.NodeTypeRangeStmt) {
				sugg = append(sugg, fmt.Sprintf(toLowerUpperInLoop, fnName))
			} else {
				sugg = append(sugg, fmt.Sprintf(toLowerUpper, fnName))
			}
		}
		return sugg
	}
}

func (s *Scrabble) testMapRuneInt(sugg []string) []string {
	if len(s.pkg.FindByValueType("map[rune]int")) != 0 {
		sugg = addSpeedComment(sugg)
		sugg = append(sugg, trySwitch)
		return sugg
	}
	if len(s.pkg.FindByValueType("map[string]int")) != 0 {
		sugg = addSpeedComment(sugg)
		sugg = append(sugg, mapRune)
		sugg = append(sugg, trySwitch)
		return sugg
	}
	return sugg
}

func (s *Scrabble) testRuneLoop(sugg []string) []string {
	ranges := s.pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, r := range ranges {
		l := r.(*astrav.RangeStmt)
		if l.Value == nil ||
			r.FindFirstByName(l.Value.(*ast.Ident).Name).IsValueType("byte") {
			sugg = append(sugg, loopRuneNotByte)
		}

		if r.FindByName("string") != nil {
			return append(sugg, typeSwitch)
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
