package suggestion

import (
	"log"
	"strings"

	"github.com/tehsphinx/exalysis/suggestion/scrabble"
	"github.com/tehsphinx/exalysis/suggestion/types"
	"golang.org/x/tools/go/loader"
	"honnef.co/go/tools/ssa"
	"honnef.co/go/tools/ssa/ssautil"
)

var exercisePkgs = map[string]types.SuggesterCreator{
	"scrabble": scrabble.NewScrabble,
}

//GetSuggestions selects the package suggestion routine and returns the suggestions
func GetSuggestions() string {
	program := ssautil.CreateProgram(loadProg(), 0)

	pkg, creator := getExercisePkg(program)
	if pkg == nil {
		log.Fatal("no known exercise package found or not implemented")
	}

	reply := getGreeting(pkg.String())

	sg := creator(program, pkg)
	sugg := sg.Suggest()
	if len(sugg) == 0 {
		reply += perfectSolution
	} else if len(sugg) < 2 {
		reply += veryGoodSolution
	} else if len(sugg) < 6 {
		reply += goodSolution
	} else {
		reply += interestingSolution
	}
	reply += strings.Join(sugg, "\n")

	return reply
}
func loadProg() *loader.Program {
	var conf loader.Config
	conf.Import(".")
	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}
	return prog
}

func getExercisePkg(program *ssa.Program) (*ssa.Package, types.SuggesterCreator) {
	for _, pkg := range program.AllPackages() {
		if sg, ok := exercisePkgs[pkg.Pkg.Name()]; ok {
			return pkg, sg
		}
	}
	return nil, nil
}

func getGreeting(pkg string) string {
	str := greeting
	switch pkg {
	case "twofer":
		str += newcomerGreeting
	}
	return str
}
