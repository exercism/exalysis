//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	OneLoop            = gtpl.NewStringTemplate("one-loop.md", MustAsset)
	CalcRangeCondition = gtpl.NewFormatTemplate("calc-range-condition.md", MustAsset)
)
