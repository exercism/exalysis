//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

//Templates to be used in the response of suggester
var (
	Greeting         = NewFormatTemplate(MustAsset("greeting.md"))
	NewcomerGreeting = NewFormatTemplate(MustAsset("newcomer_greeting.md"))
	Improvement      = NewFormatTemplate(MustAsset("improvement.md"))
	Todo             = NewFormatTemplate(MustAsset("todo.md"))
	Praise           = NewFormatTemplate(MustAsset("praise.md"))
	NotLinted        = NewFormatTemplate(MustAsset("not_linted.md"))
	NotFormatted     = NewFormatTemplate(MustAsset("not_formatted.md"))
)
