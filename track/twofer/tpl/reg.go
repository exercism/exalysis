//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/exercism/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	PlusUsed           = gtpl.NewStringTemplate("plus-used.md", MustAsset)
	Stub               = gtpl.NewStringTemplate("stub.md", MustAsset)
	CommentFormat      = gtpl.NewStringTemplate("comment-format.md", MustAsset)
	MissingComment     = gtpl.NewFormatTemplate("missing-comment.md", MustAsset)
	WrongCommentFormat = gtpl.NewFormatTemplate("wrong-comment-format.md", MustAsset)
	MinimalConditional = gtpl.NewStringTemplate("minimal-conditional.md", MustAsset)
	UseStringPH        = gtpl.NewStringTemplate("use-string-placeholder.md", MustAsset)
	GeneralizeName     = gtpl.NewStringTemplate("generalize-name.md", MustAsset)
	StringsJoin        = gtpl.NewStringTemplate("strings-join.md", MustAsset)
)
