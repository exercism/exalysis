package scrabble

import "strings"

// Score Compute the scrabble score for that input string.
func Score(input string) int {
	score := 0
	scoreMap := map[string]int{
		"AEIOULNRSTaeioulnrst": 1,
		"DGdg":                 2,
		"BCMPbcmp":             3,
		"FHVWYfhvwy":           4,
		"Kk":                   5,
		"JXjx":                 8,
		"QZqz":                 10,
	}

	for _, v := range input {
		for key, value := range scoreMap {
			if strings.Contains(key, string(v)) {
				score += value
				break
			}
		}
	}

	return score
}
