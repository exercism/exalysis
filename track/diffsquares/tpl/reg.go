//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	SquareSumLoop      = gtpl.NewStringTemplate("square-sum-loop.md", MustAsset)
	SumSquareLoop      = gtpl.NewStringTemplate("sum-square-loop.md", MustAsset)
	CalcRangeCondition = gtpl.NewFormatTemplate("calc-range-condition.md", MustAsset)
	MathPow            = gtpl.NewFormatTemplate("math-pow.md", MustAsset)
	Dry                = gtpl.NewFormatTemplate("dry.md", MustAsset)
	BasicFloat64       = gtpl.NewFormatTemplate("basic-float64.md", MustAsset)
)
