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
	examGeneralizeNames,
	examFmt,
	examComments,
	examConditional,
	examStringsJoin,
}

func examStringsJoin(pkg *astrav.Package, r *extypes.Response) {
	node := pkg.FindFirstByName("Join")
	if node != nil {
		r.AppendImprovement(tpl.StringsJoin)
	}
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

	var spfCount int
	nodes = pkg.FindByName("Sprintf")
	for _, fmtSprintf := range nodes {
		if !fmtSprintf.IsNodeType(astrav.NodeTypeSelectorExpr) {
			continue
		}

		spfCount++
		if 1 < spfCount {
			r.AppendImprovement(tpl.MinimalConditional)
		}

		bLit := fmtSprintf.Parent().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if bLit != nil && bytes.Contains(bLit.GetSource(), []byte("%v")) {
			r.AppendImprovement(tpl.UseStringPH)
			break
		} else if bLit != nil {
			continue
		}

		if bytes.Contains(pkg.FindFirstByName("ShareWith").GetSource(), []byte("%v")) {
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
	matches := outputPart.FindAllIndex(pkg.FindFirstByName("ShareWith").GetSource(), -1)
	if 1 < len(matches) {
		r.AppendImprovement(tpl.MinimalConditional)
	}
}

func examGeneralizeNames(pkg *astrav.Package, r *extypes.Response) {
	contains := bytes.Contains(pkg.FindFirstByName("ShareWith").GetSource(), []byte("Alice"))
	if !contains {
		contains = bytes.Contains(pkg.FindFirstByName("ShareWith").GetSource(), []byte("Bob"))
	}
	if contains {
		r.AppendImprovement(tpl.GeneralizeName)
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
		text := cGroup.Children()[0].(*astrav.Comment).Text
		c := strings.TrimSpace(strings.Replace(strings.Replace(text, "/*", "", 1), "//", "", 1))

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
	var added bool
	return func(r *extypes.Response) {
		if added {
			return
		}
		added = true
		r.AppendImprovement(tpl.Stub)
	}
}
func getAddCommentFormat() func(r *extypes.Response) {
	var added bool
	return func(r *extypes.Response) {
		if added {
			return
		}
		added = true
		r.AppendOutro(tpl.CommentFormat)
	}
}
