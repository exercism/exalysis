//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	Unicode         = gtpl.NewFormatTemplate("unicode.md", MustAsset)
	UnicodeLoop     = gtpl.NewFormatTemplate("unicode_loop.md", MustAsset)
	MapRune         = gtpl.NewFormatTemplate("maprune.md", MustAsset)
	TypeConversion  = gtpl.NewStringTemplate("type_conversion.md", MustAsset)
	TrySwitch       = gtpl.NewStringTemplate("try_switch.md", MustAsset)
	LoopRuneNotByte = gtpl.NewStringTemplate("loop_rune_not_byte.md", MustAsset)
	IfsToSwitch     = gtpl.NewStringTemplate("ifs_to_switch.md", MustAsset)
	GoRoutines      = gtpl.NewStringTemplate("go_routines.md", MustAsset)
	MoveMap         = gtpl.NewStringTemplate("move_map.md", MustAsset)
)
