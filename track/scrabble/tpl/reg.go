//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import (
	"github.com/tehsphinx/exalysis/tpl"
)

//Templates to be used in the response of suggester
var (
	Unicode         = tpl.NewFormatTemplate("unicode.md", MustAsset)
	UnicodeLoop     = tpl.NewFormatTemplate("unicode_loop.md", MustAsset)
	MapRune         = tpl.NewFormatTemplate("maprune.md", MustAsset)
	TypeConversion  = tpl.NewStringTemplate("type_conversion.md", MustAsset)
	TrySwitch       = tpl.NewStringTemplate("try_switch.md", MustAsset)
	LoopRuneNotByte = tpl.NewStringTemplate("loop_rune_not_byte.md", MustAsset)
	IfsToSwitch     = tpl.NewStringTemplate("ifs_to_switch.md", MustAsset)
	GoRoutines      = tpl.NewStringTemplate("go_routines.md", MustAsset)
	MoveMap         = tpl.NewStringTemplate("move_map.md", MustAsset)
)
