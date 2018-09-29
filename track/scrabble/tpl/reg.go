//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import (
	"github.com/tehsphinx/exalysis/tpl"
)

//Templates to be used in the response of suggester
var (
	Unicode         = tpl.NewFormatTemplate(MustAsset("unicode.md"))
	UnicodeLoop     = tpl.NewFormatTemplate(MustAsset("unicode_loop.md"))
	MapRune         = tpl.NewFormatTemplate(MustAsset("maprune.md"))
	TypeConversion  = tpl.NewFormatTemplate(MustAsset("type_conversion.md"))
	TrySwitch       = tpl.NewFormatTemplate(MustAsset("try_switch.md"))
	LoopRuneNotByte = tpl.NewFormatTemplate(MustAsset("loop_rune_not_byte.md"))
	IfsToSwitch     = tpl.NewFormatTemplate(MustAsset("ifs_to_switch.md"))
	GoRoutines      = tpl.NewFormatTemplate(MustAsset("go_routines.md"))
	MoveMap         = tpl.NewFormatTemplate(MustAsset("move_map.md"))
)
