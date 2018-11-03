//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	ErrorMessageFormat = gtpl.NewStringTemplate("error-message-format.md", MustAsset)
	RuneToByte         = gtpl.NewStringTemplate("rune-to-byte.md", MustAsset)
	MultiStringConv    = gtpl.NewStringTemplate("multiple-string-conversions.md", MustAsset)
	Increase           = gtpl.NewStringTemplate("increase.md", MustAsset)
	DeclareNeeded      = gtpl.NewFormatTemplate("declare-when-needed.md", MustAsset)
	InvertIf           = gtpl.NewStringTemplate("invert-if.md", MustAsset)
)
