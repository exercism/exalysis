package scrabble

import (
	"strings"
)

var letterScores [26]int

func setScores(scoreStr string, score int) {
	for _, letter := range scoreStr {
		letterScores[letter-'A'] = score
	}
}

func init() {
	setScores("AEIOULNRST", 1)
	setScores("DG", 2)
	setScores("BCMP", 3)
	setScores("FHWVY", 4)
	setScores("K", 5)
	setScores("JX", 8)
	setScores("QZ", 10)
}

// Score computes a scrabble score for the given word.
func Score(word string) int {
	result := 0
	for _, letter := range strings.ToUpper(word) {
		result += letterScores[letter-'A']
	}
	return result
}
