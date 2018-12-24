package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func mergeFreqMap(c chan FreqMap, num int) FreqMap {
	out := FreqMap{}

	for i := 0; i < num; i++ {
		select {
		case m := <-c:
			for k, v := range m {
				out[k] += v
			}
		}
	}
	return out
}

// func ConcurrentFrequency(texts []string) FreqMap {
// 	var globMap sync.Map
// 	for _, s := range texts {
// 		go func() {
// 			result := Frequency(s)
// 			for k, v := range result {
// 				globMap
// 			}
// 		}()
// 	}

// }

func ConcurrentFrequency(texts []string) FreqMap {
	// This isnt the right way to do it
	// maps := make([]FreqMap, 3)
	var wg sync.WaitGroup
	c := make(chan FreqMap)
	for _, s := range texts {
		wg.Add(1)
		go func(wg *sync.WaitGroup, str string) {
			defer wg.Done()
			result := Frequency(str)
			c <- result
		}(&wg, s)
	}
	result := mergeFreqMap(c, len(texts))
	wg.Wait()

	return result
}
