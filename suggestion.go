package exalysis

import (
	"fmt"
	"go/format"
	"go/token"
	"io/ioutil"
	"log"
	"strings"

	"github.com/golang/lint"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/dbg"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"github.com/tehsphinx/exalysis/track/scrabble"
)

var (
	//LintMinConfidence sets the min confidence for linting
	LintMinConfidence float64
)

var exercisePkgs = map[string]extypes.SuggestionFunc{
	"scrabble": scrabble.Suggest,
}

//GetSuggestions selects the package suggestion routine and returns the suggestions
func GetSuggestions() (string, string) {
	folder := astrav.NewFolder(".")
	_, err := folder.ParseFolder()
	if err != nil {
		log.Fatal(err)
	}

	pkg, suggFunc := getExercisePkg(folder)
	if suggFunc == nil {
		log.Fatal("no known exercise package found or not implemented")
	}

	var r = extypes.NewResponse()
	addGreeting(r, pkg.Name, folder.StudentName())

	files := getFiles(folder)
	resFmt := fmtCode(files)
	resLint := lintCode(files)

	if resFmt != "" {
		dbg.Cyan("#### gofmt")
		fmt.Println(resFmt)
		r.AppendTodo(gtpl.NotFormatted)
	}
	if resLint != "" {
		dbg.Cyan("#### golint")
		fmt.Println(resLint)
		r.AppendTodo(gtpl.NotLinted)
	}

	suggFunc(pkg, r)

	return r.GetAnswerString(), approval(r, resFmt == "", resLint == "")
}

func getFiles(prog *astrav.Folder) map[string][]byte {
	var files = map[string][]byte{}
	prog.FSet.Iterate(func(file *token.File) bool {
		if !strings.HasSuffix(file.Name(), "_test.go") {

			b, err := ioutil.ReadFile(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			files[file.Name()] = b
		}
		return true
	})

	return files
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

func getExercisePkg(program *astrav.Folder) (*astrav.Package, extypes.SuggestionFunc) {
	for name, pkg := range program.Pkgs {
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
	rating := "\n\n" + dbg.Sprint(dbg.ColorCyan, "#### Rating Suggestion")
	var approve string
	if !gofmt {
		rating += dbg.Sprint(dbg.ColorRed, "go code not formatted")
		if approve == "" {
			approve = dbg.Sprint(dbg.ColorRed, "NO APPROVAL")
		}
	}
	if !golint {
		rating += dbg.Sprint(dbg.ColorRed, "go code not linted")
		if approve == "" {
			approve = dbg.Sprint(dbg.ColorRed, "NO APPROVAL")
		}
	}

	l := r.LenSuggestions()
	var suggsAdded = fmt.Sprintf("Suggestions added: %d", l)
	if 5 < l {
		rating += dbg.Sprint(dbg.ColorRed, suggsAdded)
		if approve == "" {
			approve = dbg.Sprint(dbg.ColorRed, "NO APPROVAL")
		}
	} else if 2 < l {
		rating += dbg.Sprint(dbg.ColorMagenta, suggsAdded)
		if approve == "" {
			approve = dbg.Sprint(dbg.ColorMagenta, "MAYBE APPROVE")
		}
	} else if 1 < l {
		rating += dbg.Sprint(dbg.ColorYellow, suggsAdded)
		if approve == "" {
			approve = dbg.Sprint(dbg.ColorYellow, "LIKELY APPROVE")
		}
	}
	if approve == "" {
		approve = dbg.Sprint(dbg.ColorGreen, "APPROVE")
	}
	rating += "Suggestion: " + approve
	return rating
}
