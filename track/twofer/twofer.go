package twofer

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/twofer/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	addStub = getAddStub()
	addCommentFormat = getAddCommentFormat()

	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	examFmt,
	examComments,
	examConditional,
}

func examFmt(pkg *astrav.Package, r *extypes.Response) {
	var plusUsed bool
	nodes := pkg.FindFirstByName("ShareWith").FindByNodeType(astrav.NodeTypeBinaryExpr)
	for _, node := range nodes {
		if node.(*astrav.BinaryExpr).Op.String() == "+" {
			plusUsed = true
		}
	}
	if plusUsed {
		r.AppendComment(tpl.PlusUsed)
		return
	}

	fmtSprintf := pkg.FindFirstByName("Sprintf")
	if fmtSprintf != nil {
		if bytes.Contains(fmtSprintf.Parent().GetSource(), []byte("%v")) {
			r.AppendImprovement(tpl.UseStringPH)
		}
	}
}

func examComments(pkg *astrav.Package, r *extypes.Response) {
	if bytes.Contains(pkg.GetSource(), []byte("stub")) {
		addStub(r)
	}

	cGroup := pkg.ChildByNodeType(astrav.NodeTypeFile).ChildByNodeType(astrav.NodeTypeCommentGroup)
	checkComment(cGroup, r, "package", "twofer")

	cGroup = pkg.FindFirstByName("ShareWith").ChildByNodeType(astrav.NodeTypeCommentGroup)
	checkComment(cGroup, r, "function", "ShareWith")
}

var outputPart = regexp.MustCompile(`, one for me\.`)

func examConditional(pkg *astrav.Package, r *extypes.Response) {
	matches := outputPart.FindAllIndex(pkg.GetSource(), -1)
	if 1 < len(matches) {
		r.AppendImprovement(tpl.MinimalConditional)
	}
}

var commentStrings = map[string]struct {
	typeString       string
	stubString       string
	prefixString     string
	wrongCommentName string
}{
	"package": {
		typeString:       "Packages",
		stubString:       "should have a package comment",
		prefixString:     "Package %s ",
		wrongCommentName: "package `%s`",
	},
	"function": {
		typeString:       "Exported functions",
		stubString:       "should have a comment",
		prefixString:     "%s ",
		wrongCommentName: "function `%s`",
	},
}

// we only do this on the first exercise. Later we ask them to use golint.
func checkComment(cGroup astrav.Node, r *extypes.Response, commentType, name string) {
	strPack := commentStrings[commentType]
	if cGroup == nil {
		r.AppendImprovement(tpl.MissingComment.Format(strPack.typeString))
		addCommentFormat(r)
	} else {
		c := strings.TrimSpace(strings.Replace(cGroup.Children()[0].(*astrav.Comment).Text, "/", "", 2))

		if strings.Contains(c, strPack.stubString) {
			addStub(r)
		} else if !strings.HasPrefix(c, fmt.Sprintf(strPack.prefixString, name)) {
			r.AppendImprovement(tpl.WrongCommentFormat.Format(fmt.Sprintf(strPack.wrongCommentName, name)))
			addCommentFormat(r)
		}
	}
}

var (
	addStub          func(r *extypes.Response)
	addCommentFormat func(r *extypes.Response)
)

func getAddStub() func(r *extypes.Response) {
	var stubAdded bool
	return func(r *extypes.Response) {
		if stubAdded {
			return
		}
		stubAdded = true
		r.AppendImprovement(tpl.Stub)
	}
}
func getAddCommentFormat() func(r *extypes.Response) {
	var commentAdded bool
	return func(r *extypes.Response) {
		if commentAdded {
			return
		}
		commentAdded = true
		r.AppendOutro(tpl.CommentFormat)
	}
}
