package exalysis

import (
	"fmt"
	"go/format"
	"log"
	"strings"

	"github.com/golang/lint"
	"github.com/logrusorgru/aurora"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/track/hamming"
	"github.com/tehsphinx/exalysis/track/raindrops"
	"github.com/tehsphinx/exalysis/track/scrabble"
	"github.com/tehsphinx/exalysis/track/twofer"
)

var (
	//LintMinConfidence sets the min confidence for linting
	LintMinConfidence float64
)

var exercisePkgs = map[string]extypes.SuggestionFunc{
	"twofer":    twofer.Suggest,
	"hamming":   hamming.Suggest,
	"raindrops": raindrops.Suggest,
	"scrabble":  scrabble.Suggest,
}

//GetSuggestions selects the package suggestion routine and returns the suggestions
func GetSuggestions() (string, string) {
	folder := astrav.NewFolder(".")
	_, err := folder.ParseFolder()
	if err != nil {
		log.Fatal(err)
	}

	var pkgName string
	pkg, suggFunc := getExercisePkg(folder)
	if suggFunc == nil {
		fmt.Println(aurora.Red("Current folder does not contain a known exercism solution!"))
		fmt.Println(aurora.Gray("Running general code checks:"))
	} else {
		pkgName = pkg.Name
		fmt.Println(aurora.Sprintf(aurora.Gray("Exercism package found: %s"), aurora.Green(pkgName)))
	}

	var r = extypes.NewResponse()
	addGreeting(r, pkgName, folder.StudentName())

	files := folder.GetRawFiles()
	fmtOk := checkFmt(files, r, pkgName)
	lintOk := checkLint(files, r, pkgName)

	if suggFunc != nil {
		suggFunc(pkg, r)
	}

	return r.GetAnswerString(), approval(r, fmtOk, lintOk)
}

func checkFmt(files map[string][]byte, r *extypes.Response, pkgName string) bool {
	resFmt := fmtCode(files)
	if resFmt == "" {
		fmt.Println(aurora.Gray("gofmt:\t\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("gofmt:\t\t"), aurora.Red("FAIL"))
	fmt.Println(resFmt)
	if pkgName == "twofer" || pkgName == "hamming" {
		r.AppendImprovement(gtpl.NotFormatted)
		return true
	} else {
		r.AppendTodo(gtpl.NotFormatted)
	}

	return false
}

func checkLint(files map[string][]byte, r *extypes.Response, pkgName string) bool {
	resLint := lintCode(files)
	if resLint == "" {
		fmt.Println(aurora.Gray("golint:\t\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("golint:\t\t"), aurora.Red("FAIL"))
	fmt.Println(resLint)
	if pkgName == "twofer" || pkgName == "hamming" {
		r.AppendImprovement(gtpl.NotLinted)
		return true
	} else {
		r.AppendTodo(gtpl.NotLinted)
	}

	return false
}

func lintCode(files map[string][]byte) string {
	l := lint.Linter{}
	ps, err := l.LintFiles(files)
	if err != nil {
		log.Fatal(err)
	}

	var lintRes string
	for _, p := range ps {
		if p.Confidence < LintMinConfidence {
			continue
		}
		lintRes += fmt.Sprintf("%s: %s\n\t%s\n\tdoc: %s\n", p.Category, p.Text, p.Position.String(), p.Link)
	}
	return lintRes
}

func fmtCode(files map[string][]byte) string {
	for _, file := range files {
		f, err := format.Source(file)
		if err != nil {
			return fmt.Sprintf("code fails to format with error: %s\n", err)
		}
		if string(f) != strings.Replace(string(file), "\r\n", "\n", -1) {
			return getDiff(file, f)
		}
	}
	return ""
}

func getDiff(current, formatted []byte) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(current)),
		B:        difflib.SplitLines(string(formatted)),
		FromFile: "Current",
		ToFile:   "Formatted",
		Context:  0,
	}
	text, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		return fmt.Sprintf("error while diffing strings: %s", err)
	}
	return text
}

func getExercisePkg(folder *astrav.Folder) (*astrav.Package, extypes.SuggestionFunc) {
	for name, pkg := range folder.Pkgs {
		if sg, ok := exercisePkgs[name]; ok {
			return pkg, sg
		}
	}
	return nil, nil
}

func addGreeting(r *extypes.Response, pkg, student string) {
	r.SetGreeting(gtpl.Greeting.Format(student))
	switch pkg {
	case "twofer":
		r.AppendGreeting(gtpl.NewcomerGreeting)
	}
}

func approval(r *extypes.Response, gofmt, golint bool) string {
	rating := aurora.Gray("Rating Suggestion\n").String()
	var approve aurora.Value
	if !gofmt && approve == nil {
		approve = aurora.Red("NO APPROVAL")
	}
	if !golint && approve == nil {
		approve = aurora.Red("NO APPROVAL")
	}

	rating += fmt.Sprintf("Todos:\t\t%d\n", aurora.Red(r.LenTodos()))
	rating += fmt.Sprintf("Suggestions:\t%d\n", aurora.Brown(r.LenImprovements()))
	rating += fmt.Sprintf("Comments:\t%d\n", aurora.Green(r.LenComments()))

	if approve == nil {
		l := r.LenSuggestions()
		switch {
		case 5 < l:
			approve = aurora.Red("NO APPROVAL")
		case 2 < l:
			approve = aurora.Magenta("MAYBE APPROVE")
		case 1 < l:
			approve = aurora.Brown("LIKELY APPROVE")
		default:
			approve = aurora.Green("APPROVE")
		}
	}
	rating += fmt.Sprintf("Suggestion:\t%s\n", approve)
	return rating
}
