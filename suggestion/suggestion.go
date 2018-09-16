package suggestion

import (
	"fmt"
	"go/format"
	"go/token"
	"io/ioutil"
	"log"
	"strings"

	"github.com/golang/lint"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/dbg"
	"github.com/tehsphinx/exalysis/suggestion/defs"
	"github.com/tehsphinx/exalysis/suggestion/scrabble"
)

var (
	//LintMinConfidence sets the min confidence for linting
	LintMinConfidence float64
)

var exercisePkgs = map[string]defs.SuggesterCreator{
	"scrabble": scrabble.NewScrabble,
}

//GetSuggestions selects the package suggestion routine and returns the suggestions
func GetSuggestions() (string, string) {
	folder := astrav.NewFolder(".")
	_, err := folder.ParseFolder()
	if err != nil {
		log.Fatal(err)
	}

	pkg, creator := getExercisePkg(folder)
	if creator == nil {
		log.Fatal("no known exercise package found or not implemented")
	}

	files := getFiles(folder)
	resFmt := fmtCode(files)
	resLint := lintCode(files)

	var sugg []string
	if resFmt != "" {
		dbg.Cyan("#### gofmt")
		fmt.Println(resFmt)
		sugg = append(sugg, notFormatted)
	}
	if resLint != "" {
		dbg.Cyan("#### golint")
		fmt.Println(resLint)
		sugg = append(sugg, notLinted)
	}

	sg := creator(pkg)
	sugg = append(sugg, sg.Suggest()...)

	reply := getGreeting(pkg.Name)
	reply += praise(sugg)
	reply += strings.Join(sugg, "\n")

	return reply, approval(sugg, resFmt == "", resLint == "")
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
		if string(f) != string(file) {
			return "code is not formatted!\n"
		}
	}
	return ""
}

func getExercisePkg(program *astrav.Folder) (*astrav.Package, defs.SuggesterCreator) {
	for name, pkg := range program.Pkgs {
		if sg, ok := exercisePkgs[name]; ok {
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

func praise(sugg []string) string {
	if len(sugg) == 0 {
		return perfectSolution
	} else if len(sugg) < 2 {
		return veryGoodSolution
	} else if len(sugg) < 6 {
		return goodSolution
	}
	return interestingSolution
}

func approval(sugg []string, gofmt, golint bool) string {
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

	var suggsAdded = fmt.Sprintf("Suggestions added: %d", len(sugg))
	if 5 < len(sugg) {
		rating += dbg.Sprint(dbg.ColorRed, suggsAdded)
		if approve == "" {
			approve = dbg.Sprint(dbg.ColorRed, "NO APPROVAL")
		}
	} else if 2 < len(sugg) {
		rating += dbg.Sprint(dbg.ColorMagenta, suggsAdded)
		if approve == "" {
			approve = dbg.Sprint(dbg.ColorMagenta, "MAYBE APPROVE")
		}
	} else if 1 < len(sugg) {
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
