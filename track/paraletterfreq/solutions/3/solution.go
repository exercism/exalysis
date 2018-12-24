package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {

	freqCh := make(chan FreqMap)

	for _, sentence := range s {
		go func(s string) {
			freqCh <- Frequency(s)
		}(sentence)
	}

	finalMap := FreqMap{}
	for i := 0; i < len(s); i++ {
		tmpFreqMap := <-freqCh
		for word, freq := range tmpFreqMap {
			finalMap[word] += freq
		}
	}
	return finalMap
}