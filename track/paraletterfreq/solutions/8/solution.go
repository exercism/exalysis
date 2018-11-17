package letter

import (
	"sync"
)

// FreqMap custom map type
type FreqMap map[rune]int

// Frequency calculates chars in string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates chars in string with concurrency
func ConcurrentFrequency(sl []string) FreqMap {
	resMap := FreqMap{}
	wg := &sync.WaitGroup{}
	ch := make(chan rune)
	for _, s := range sl {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for _, r := range s {
				ch <- r
			}
		}(s)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		resMap[r]++
	}

	return resMap
}
