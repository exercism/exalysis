//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

// Templates to be used in the response of suggester
var (
	Unicode         = gtpl.NewFormatTemplate("unicode.md", MustAsset)
	UnicodeLoop     = gtpl.NewFormatTemplate("unicode-loop.md", MustAsset)
	MapRune         = gtpl.NewFormatTemplate("maprune.md", MustAsset)
	FlattenMap      = gtpl.NewFormatTemplate("flatten-map.md", MustAsset)
	TypeConversion  = gtpl.NewStringTemplate("type-conversion.md", MustAsset)
	TrySwitch       = gtpl.NewStringTemplate("try-switch.md", MustAsset)
	LoopRuneNotByte = gtpl.NewStringTemplate("loop-rune-not-byte.md", MustAsset)
	IfsToSwitch     = gtpl.NewStringTemplate("ifs-to-switch.md", MustAsset)
	GoRoutines      = gtpl.NewStringTemplate("go-routines.md", MustAsset)
	MoveMap         = gtpl.NewStringTemplate("move-map.md", MustAsset)
	Regex           = gtpl.NewStringTemplate("regex.md", MustAsset)
	Challenge       = gtpl.NewStringTemplate("challenge.md", MustAsset)
	SliceRuneConv   = gtpl.NewStringTemplate("slice-rune-conversion.md", MustAsset)
)
