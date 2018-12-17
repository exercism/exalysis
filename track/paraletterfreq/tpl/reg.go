//go:generate go-bindata -ignore=\.go -pkg=tpl -o=bindata.go ./...

package tpl

import "github.com/tehsphinx/exalysis/gtpl"

//Templates to be used in the response of suggester
var (
	ConcurrencyNotFaster    = gtpl.NewStringTemplate("concurrency-not-faster.md", MustAsset)
	WaitGroup               = gtpl.NewStringTemplate("waitgroup.md", MustAsset)
	WaitGroupAddOne         = gtpl.NewStringTemplate("waitgroup-add-one.md", MustAsset)
	WaitGroupNotNeeded      = gtpl.NewStringTemplate("waitgroup-not-needed.md", MustAsset)
	RangeChan               = gtpl.NewStringTemplate("range-chan.md", MustAsset)
	BufferSizeLen           = gtpl.NewStringTemplate("buffer-size-len.md", MustAsset)
	CombineMapsWhileWaiting = gtpl.NewStringTemplate("combine-maps-while-waiting.md", MustAsset)
	ForRangeNoVars          = gtpl.NewStringTemplate("for-range-novars.md", MustAsset)
	SelectNotNeeded         = gtpl.NewStringTemplate("select-not-needed.md", MustAsset)
	GoroutineLeak           = gtpl.NewStringTemplate("goroutine-leak.md", MustAsset)
	Mutex                   = gtpl.NewStringTemplate("mutex.md", MustAsset)
)
