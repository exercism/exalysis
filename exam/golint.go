package exam

import (
	"fmt"
	"log"

	"github.com/golang/lint"
	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
)

var (
	//LintMinConfidence sets the min confidence for linting
	LintMinConfidence float64
)

//GoLint runs golint on provided go files and adds suggestions to the response
func GoLint(files map[string][]byte, r *extypes.Response, pkgName string) bool {
	resLint := lintCode(files)
	if resLint == "" {
		fmt.Println(aurora.Gray("golint:\t\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("golint:\t\t"), aurora.Red("FAIL"))
	fmt.Println(resLint)
	if pkgName == "twofer" || pkgName == "hamming" {
		r.AppendImprovement(gtpl.NotLinted)
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
