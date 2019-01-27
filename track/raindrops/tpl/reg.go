//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

// Templates to be used in the response of suggester
var (
	ManyLoops       = gtpl.NewStringTemplate("many-loops.md", MustAsset)
	StringsBuilder  = gtpl.NewFormatTemplate("strings-builder.md", MustAsset)
	PlusEqual       = gtpl.NewStringTemplate("plus-equal.md", MustAsset)
	Itoa            = gtpl.NewStringTemplate("itoa.md", MustAsset)
	ExtensiveFor    = gtpl.NewStringTemplate("extensive-for-loop.md", MustAsset)
	LoopMap         = gtpl.NewStringTemplate("loop-map.md", MustAsset)
	RemoveExtraBool = gtpl.NewFormatTemplate("remove-extra-bool.md", MustAsset)
	FmtPrint        = gtpl.NewFormatTemplate("fmt-print.md", MustAsset)
	AllCases        = gtpl.NewStringTemplate("all-cases.md", MustAsset)
)
