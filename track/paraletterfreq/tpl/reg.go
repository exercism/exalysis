//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	WaitGroup            = gtpl.NewStringTemplate("waitgroup.md", MustAsset)
	ConcurrencyNotFaster = gtpl.NewStringTemplate("concurrency-not-faster.md", MustAsset)
	WaitGroupAddOne      = gtpl.NewStringTemplate("waitgroup-add-one.md", MustAsset)
	WaitGroupNotNeeded   = gtpl.NewStringTemplate("waitgroup-not-needed.md", MustAsset)
	BufferSizeLen        = gtpl.NewStringTemplate("buffer-size-len.md", MustAsset)
)
