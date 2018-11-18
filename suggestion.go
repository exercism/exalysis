package exalysis

import (
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/exam"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/track/diffsquares"
	"github.com/tehsphinx/exalysis/track/hamming"
	"github.com/tehsphinx/exalysis/track/isogram"
	"github.com/tehsphinx/exalysis/track/luhn"
	"github.com/tehsphinx/exalysis/track/raindrops"
	"github.com/tehsphinx/exalysis/track/scrabble"
	"github.com/tehsphinx/exalysis/track/twofer"
)

var exercisePkgs = map[string]extypes.SuggestionFunc{
	"twofer":      twofer.Suggest,
	"hamming":     hamming.Suggest,
	"raindrops":   raindrops.Suggest,
	"scrabble":    scrabble.Suggest,
	"isogram":     isogram.Suggest,
	"diffsquares": diffsquares.Suggest,
	"luhn":        luhn.Suggest,
}

//GetSuggestions selects the package suggestion routine and returns the suggestions
func GetSuggestions(path string) (string, string) {
	var r = extypes.NewResponse()
	folder := astrav.NewFolder(path)
	_, err := folder.ParseFolder()
	if err != nil {
		addGreeting(r, "", "there")
		r.AppendTodo(gtpl.Compile)
		return r.GetAnswerString(), rating(r, nil, "")
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

	addGreeting(r, pkgName, folder.StudentName())
	examRes, err := exam.All(folder, r, pkgName)
	if err != nil {
		log.Fatal(err)
	}

	if suggFunc != nil {
		suggFunc(pkg, r)
	}

	return r.GetAnswerString(), rating(r, examRes, pkgName)
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

func rating(r *extypes.Response, examRes *exam.Result, pkgName string) string {
	rating := aurora.Gray("Rating Suggestion\n").String()
	rating += fmt.Sprintf("Todos:\t\t%d\n", aurora.Red(r.LenTodos()))
	rating += fmt.Sprintf("Suggestions:\t%d\n", aurora.Brown(r.LenImprovements()))
	rating += fmt.Sprintf("Comments:\t%d\n", aurora.Green(r.LenComments()))

	approve := approval(r, examRes, pkgName)
	rating += fmt.Sprintf("Suggestion:\t%s\n", approve)
	return rating
}

func approval(r *extypes.Response, examRes *exam.Result, pkgName string) aurora.Value {
	var gofmt, golint bool
	if examRes != nil {
		gofmt = examRes.GoFmt
		golint = examRes.GoLint
	}

	// don't be so strict on the first exercises
	switch pkgName {
	case "twofer":
		gofmt = true
		golint = true
	case "hamming":
		golint = true
	}

	if !gofmt {
		return aurora.Red("NO APPROVAL")
	}
	if !golint {
		return aurora.Red("NO APPROVAL")
	}
	if r.LenTodos() != 0 {
		return aurora.Red("NO APPROVAL")
	}

	l := r.LenImprovements()
	switch {
	case 5 < l:
		return aurora.Red("NO APPROVAL")
	case 2 < l:
		return aurora.Magenta("MAYBE APPROVE")
	case 1 < l:
		return aurora.Brown("LIKELY APPROVE")
	}

	return aurora.Green("APPROVE")
}
