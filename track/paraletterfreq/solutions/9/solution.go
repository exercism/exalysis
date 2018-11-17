package letter

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
	ch := make(chan FreqMap)

	for _, s := range sl {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	for i := 0; i < len(sl); i++ {
		for k, v := range <-ch {
			resMap[k] += v
		}
	}

	return resMap
}
