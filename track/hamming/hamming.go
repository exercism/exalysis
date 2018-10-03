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
	examInvertIf,
	examRuneToByte,
	examMultipleStringConversions,
	examIncrease,
	examErrorMessage,
}

func examInvertIf(pkg *astrav.Package, r *extypes.Response) {
	for _, ifNode := range pkg.FindByNodeType(astrav.NodeTypeIfStmt) {
		loop := ifNode.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
		if loop == nil {
			loop = ifNode.FindFirstByNodeType(astrav.NodeTypeForStmt)
		}
		condition := ifNode.ChildByNodeType(astrav.NodeTypeBinaryExpr).(*astrav.BinaryExpr)
		if loop != nil && condition != nil && condition.Op.String() == "==" {
			r.AppendImprovement(tpl.InvertIf)
		}
	}
}

func examRuneToByte(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByName("byte")
	for _, node := range nodes {
		for _, n := range node.Siblings() {
			if n.ValueType().String() == "rune" {
				r.AppendComment(tpl.RuneToByte)
			}
		}
	}
}

func examMultipleStringConversions(pkg *astrav.Package, r *extypes.Response) {
	rngNode := pkg.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
	if rngNode == nil {
		return
	}

	count := 0
	for _, node := range rngNode.FindByName("string") {
		if node.Parent().IsNodeType(astrav.NodeTypeCallExpr) {
			count++
		}
	}
	if 1 < count {
		r.AppendImprovement(tpl.MultiStringConv)
	}
}

func examIncrease(pkg *astrav.Package, r *extypes.Response) {
	loop := pkg.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
	if loop == nil {
		loop = pkg.FindFirstByNodeType(astrav.NodeTypeForStmt)
	}
	for _, node := range loop.FindByNodeType(astrav.NodeTypeBinaryExpr) {
		if node.(*astrav.BinaryExpr).Op.String() == "+" {
			r.AppendComment(tpl.Increase)
		}
	}
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
