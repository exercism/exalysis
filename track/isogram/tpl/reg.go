//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	Unicode             = gtpl.NewFormatTemplate("unicode.md", MustAsset)
	RegexInFunc         = gtpl.NewStringTemplate("regex-in-func.md", MustAsset)
	MustCompile         = gtpl.NewStringTemplate("mustcompile.md", MustAsset)
	JustReturn          = gtpl.NewStringTemplate("just-return.md", MustAsset)
	NonExistingMapValue = gtpl.NewStringTemplate("nonexisting-map-value.md", MustAsset)
	IsLetter            = gtpl.NewStringTemplate("isletter.md", MustAsset)
)
