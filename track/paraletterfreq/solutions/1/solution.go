package letter

import (
	"sync"
)

type FreqMap map[rune]int

type SyncFreqMap struct {
	sync.Mutex
	FreqMap
}

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) FreqMap {
	m := FreqMap{}
	var wg sync.WaitGroup
	wg.Add(len(input))
	var ch = make(chan rune, len(input))
	for _, in := range input {
		go func(chunk string, c chan rune) {
			for _, r := range chunk {
				c <- r
			}
			wg.Done()
		}(in, ch)
	}
	go func() {
		for {
			select {
			case ltr := <-ch:
				m[ltr]++
			}
		}
	}()
	wg.Wait()
	return m
}
