package letter

// FreqMap structure
type FreqMap map[rune]int

// Frequency counts letter frequencies in a text
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func frequencyWithChannel(s string, ch chan FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	ch <- m
}

func combineMaps(maps []FreqMap) FreqMap {
	res := FreqMap{}
	for _, m := range maps {
		for letter, freq := range m {
			res[letter] += freq
		}
	}
	return res
}

// ConcurrentFrequency calculates letter frequencies in each text on parallelism
// And combine all results
func ConcurrentFrequency(texts []string) FreqMap {
	maps := []FreqMap{}
	channel := make(chan FreqMap, 3)
	for _, text := range texts {
		go frequencyWithChannel(text, channel)
	}
	for range texts {
		maps = append(maps, <-channel)
	}
	return combineMaps(maps)
}
