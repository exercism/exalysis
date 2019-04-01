//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/exercism/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	RegexInFunc = gtpl.NewStringTemplate("regex-in-func.md", MustAsset)
	MustCompile = gtpl.NewStringTemplate("mustcompile.md", MustAsset)
	RegexToFast = gtpl.NewStringTemplate("regex-to-fast.md", MustAsset)
	OneLoop     = gtpl.NewStringTemplate("one-loop.md", MustAsset)
)
