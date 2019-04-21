package exam

import (
	"bytes"
	"fmt"
	"go/format"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/logrusorgru/aurora"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/tehsphinx/astrav"
)

// GoFmt runs gofmt on provided go files and adds suggestions to the response
func GoFmt(folder *astrav.Folder, r *extypes.Response, pkgName string) bool {
	files := folder.GetRawFiles()

	resFmt := fmtCode(files)
	if resFmt == "" {
		fmt.Println(aurora.Gray("gofmt:\t\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("gofmt:\t\t"), aurora.Red("FAIL"))
	fmt.Println(resFmt)
	if pkgName == "twofer" || pkgName == "hamming" {
		r.AppendImprovementTpl(gtpl.NotFormatted)
	} else {
		r.AppendTodoTpl(gtpl.NotFormatted)
	}

	return false
}

func fmtCode(files map[string][]byte) string {
	for _, file := range files {
		file = bytes.Replace(file, []byte{'\r', '\n'}, []byte{'\n'}, -1)
		f, err := format.Source(file)
		if err != nil {
			return fmt.Sprintf("code fails to format with error: %s\n", err)
		}
		if string(f) != string(file) && string(f) != string(append(file, '\n')) {
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
