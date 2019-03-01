//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

// Templates to be used in the response of suggester
var (
	BitsRotate64    = gtpl.NewStringTemplate("bits-rotate64.md", MustAsset)
	BitsRotate      = gtpl.NewStringTemplate("bits-rotate.md", MustAsset)
	StaticTotalNice = gtpl.NewStringTemplate("static-total-nice.md", MustAsset)
	MathPow         = gtpl.NewStringTemplate("math-pow.md", MustAsset)
	StaticTotal     = gtpl.NewStringTemplate("static-total.md", MustAsset)
	ErrorFormatted  = gtpl.NewStringTemplate("error-formatted.md", MustAsset)
	CheckError      = gtpl.NewStringTemplate("check-error.md", MustAsset)
)
