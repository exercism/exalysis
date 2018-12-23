package grains

import (
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/grains/tpl"
)

// Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examIgnoreError,
}

func examIgnoreError(pkg *astrav.Package, r *extypes.Response) {
	blank := pkg.FindFirstIdentByName("_")
	if blank != nil {
		r.AppendImprovement(tpl.CheckError)
	}
}
