package grains

import (
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/twofer/tpl"
)

// Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examStringsJoin,
}

func examStringsJoin(pkg *astrav.Package, r *extypes.Response) {
	node := pkg.FindFirstByName("Join")
	if node != nil {
		r.AppendImprovement(tpl.StringsJoin)
	}
}
