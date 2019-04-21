package luhn

import (
	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/track/luhn/tpl"
	"github.com/tehsphinx/astrav"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examRegexCompileInFunc,
	examOneLoop,
}

func examOneLoop(pkg *astrav.Package, r *extypes.Response) {
	loops := len(pkg.FindByNodeType(astrav.NodeTypeForStmt))
	loops += len(pkg.FindByNodeType(astrav.NodeTypeRangeStmt))
	if 1 < loops {
		r.AppendBlockTpl(tpl.OneLoop)
	}
}

func examRegexCompileInFunc(pkg *astrav.Package, r *extypes.Response) {
	main := pkg.FindFirstByName("Valid")
	regComp := pkg.FindFirstByName("Compile")
	if regComp != nil && main.Contains(regComp) {
		r.AppendTodoTpl(tpl.RegexInFunc)
		r.AppendTodoTpl(tpl.MustCompile)
	}
	if regComp != nil {
		r.AppendBlockTpl(tpl.RegexToFast)
	}

	regComp = pkg.FindFirstByName("MustCompile")
	if regComp != nil && main.Contains(regComp) {
		r.AppendTodoTpl(tpl.RegexInFunc)
	}
	if regComp != nil {
		r.AppendBlockTpl(tpl.RegexToFast)
	}
	regComp = pkg.FindFirstByName("MatchString")
	if regComp != nil {
		r.AppendBlockTpl(tpl.RegexToFast)
	}
}
