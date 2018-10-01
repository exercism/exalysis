package twofer

import (
	"bytes"
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

	//TODO: add more checks
}

func examComments(pkg *astrav.Package, r *extypes.Response) {
	if bytes.Contains(pkg.GetSource(), []byte("stub")) {
		addStub(r)
	}

	//TODO: create a function for this to reduce replicated code
	cGroup := pkg.ChildByNodeType(astrav.NodeTypeFile).ChildByNodeType(astrav.NodeTypeCommentGroup)
	if cGroup == nil {
		r.AppendImprovement(tpl.MissingComment.Format("Packages"))
		addCommentFormat(r)
	} else {
		c := strings.TrimSpace(strings.Replace(cGroup.Children()[0].(*astrav.Comment).Text, "/", "", 2))

		if strings.Contains(c, "should have a package comment") {
			addStub(r)
		} else if !strings.HasPrefix(c, "Package twofer ") {
			r.AppendImprovement(tpl.WrongCommentFormat.Format("package twofer"))
			addCommentFormat(r)
		}
	}

	cGroup = pkg.FindFirstByName("ShareWith").ChildByNodeType(astrav.NodeTypeCommentGroup)
	if cGroup == nil {
		r.AppendImprovement(tpl.MissingComment.Format("Exported functions"))
		addCommentFormat(r)
	} else {
		c := strings.TrimSpace(strings.Replace(cGroup.Children()[0].(*astrav.Comment).Text, "/", "", 2))

		if strings.Contains(c, "should have a comment") {
			addStub(r)
		} else if !strings.HasPrefix(c, "ShareWith ") {
			r.AppendImprovement(tpl.WrongCommentFormat.Format("function `ShareWith`"))
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
