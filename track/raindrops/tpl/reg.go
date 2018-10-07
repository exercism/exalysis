//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	ManyLoops       = gtpl.NewStringTemplate("many-loops.md", MustAsset)
	ConcatNotNeeded = gtpl.NewStringTemplate("concat-not-needed.md", MustAsset)
)
