// Package letter provide methods to count letter frequency.
package letter

import "sync"

// FreqMap express a character map with the number of occurencies for each letter.
type FreqMap map[rune]int

// Frequency takes a string and returns a FreqMap with the total frequency of each letter.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes a list of strings and calculate letter frequency concurrently, it returns a FreqMap with the total frequency of each letter.
func ConcurrentFrequency(texts []string) FreqMap {
	return <-collect(count(texts))
}

func count(texts []string) <-chan FreqMap {
	ch := make(chan FreqMap, len(texts))
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(texts))
		for _, t := range texts {
			go func(t string) {
				ch <- Frequency(t)
				wg.Done()
			}(t)
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

func collect(ch <-chan FreqMap) <-chan FreqMap {
	res := make(chan FreqMap)
	go func() {
		m := FreqMap{}
		for r := range ch {
			merge(m, r)
		}
		res <- m
		close(res)
	}()
	return res
}

func merge(to, this FreqMap) {
	for k, v := range this {
		to[k] += v
	}
}
