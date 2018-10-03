package hamming

import (
	"strings"
	"unicode"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/hamming/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	addErrorMsgFormat = getAddErrorMsgFormat()

	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examErrorMessage,
}

func examErrorMessage(pkg *astrav.Package, r *extypes.Response) {
	checkErrMessage(pkg.FindByName("New"), r)
	checkErrMessage(pkg.FindByName("Errorf"), r)
}

func checkErrMessage(nodes []astrav.Node, r *extypes.Response) {
	for _, node := range nodes {
		if node.NodeType() == astrav.NodeTypeIdent {
			continue
		}
		errMsgNode := node.(*astrav.SelectorExpr).Parent().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if errMsgNode == nil {
			continue
		}

		errText := errMsgNode.(*astrav.BasicLit).Value

		// check punctuation
		if strings.HasSuffix(errText, ".") {
			addErrorMsgFormat(r)
			continue
		}

		// check if first letter is capitalized and second not.
		var isUpper bool
		for i, rn := range strings.Split(errText, " ")[0] {
			// first letter is " or `
			if i == 2 {
				if isUpper && !unicode.IsUpper(rn) {
					addErrorMsgFormat(r)
				}
				break
			}
			isUpper = unicode.IsUpper(rn)
		}
	}
}

var addErrorMsgFormat func(r *extypes.Response)

func getAddErrorMsgFormat() func(r *extypes.Response) {
	var added bool
	return func(r *extypes.Response) {
		if added {
			return
		}
		added = true
		r.AppendImprovement(tpl.ErrorMessageFormat)
	}
}
