package suggestion

import (
	"fmt"
	"strings"

	"github.com/tehsphinx/exalysis/suggestion/scrabble"
	"golang.org/x/tools/go/loader"
)

//GetSuggestions selects the package suggestion routine and returns the suggestions
func GetSuggestions(pkgName string, local *loader.PackageInfo, prog *loader.Program) string {
	var sugg []string

	reply := getGreeting(pkgName)
	switch pkgName {
	case "scrabble":
		sugg = scrabble.Scrabble(local, prog)
	default:
		return fmt.Sprintf("ERROR: No suggestions for package %s implemented!", pkgName)
	}

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

func getGreeting(pkg string) string {
	str := greeting
	switch pkg {
	case "twofer":
		str += newcomerGreeting
	}
	return str
}
