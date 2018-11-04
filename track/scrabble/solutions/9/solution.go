package scrabble

import "strings"

var scoreMap = map[string]int{
	"AEIOULNRSTaeioulnrst": 1,
	"DGdg":                 2,
	"BCMPbcmp":             3,
	"FHVWYfhvwy":           4,
	"Kk":                   5,
	"JXjx":                 8,
	"QZqz":                 10,
}

// Score Compute the scrabble score for that input string.
func Score(input string) int {
	score := 0

	for _, r := range input {
		score += getVal(r)
	}

	return score
}

func getVal(r rune) int {
	for key, value := range scoreMap {
		if strings.Contains(key, string(r)) {
			return value
		}
	}

	return 0
}
