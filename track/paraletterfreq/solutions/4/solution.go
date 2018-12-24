package letter

// FreqMap is a map of each letter and it's frequency
type FreqMap map[rune]int

// Frequency is a function to count the frequency of letters in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency is a function that uses channels and goroutines to
// concurrently count the letter frequency in different strings at once.
func ConcurrentFrequency(texts []string) FreqMap {
	channel := make(chan FreqMap, len(texts))
	for _, text := range texts {
		go func(text string) {
			channel <- Frequency(text)
		}(text)
	}

	result := FreqMap{}
	for range texts {
		for letter, count := range <-channel {
			result[letter] += count
		}
	}
	return result
}