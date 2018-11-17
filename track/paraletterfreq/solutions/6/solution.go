package letter

// FreqMap is a map that maps runes to int
// It is a map that gives the count of each
// letter in a string, with the letters as keys.
type FreqMap map[rune]int

// Frequency takes in a string and returns a FreqMap
// map where the keys are the unique letters in the input
// string and the values are the counts of each of those
// letters in the input string. This is a synchronous
// function.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes in a slice of strings and returns
// a FreqMap where the keys are the unique letters in all the
// input strings in the input slice combined, and the values
// are the counts of those letters.
func ConcurrentFrequency(stringList []string) FreqMap {

	c := make(chan FreqMap, len(stringList))
	finalMap := FreqMap{}

	// Spawn goroutines for each string in the input slice
	for _, str := range stringList {
		go func(aString string) {
			c <- Frequency(aString)
		}(str)
	}

	// Loop over the channel and return final FreqMap
	for i := 1; i <= len(stringList); i++ {
		aMap := <-c
		for key, value := range aMap {
			finalMap[key] += value
		}
	}
	return finalMap
}
