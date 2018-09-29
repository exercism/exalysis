//Package scrabble implements a scoring function for taking scrabble words and finding the associated points
package scrabble

import "strings"

//Score returns an int based on the Scrabble score of the string that is passed in as a parameter
func Score(input string) (scrabbleScore int) {
	scoreMap := map[string]int{
		"A": 1,
		"B": 3,
		"C": 3,
		"D": 2,
		"E": 1,
		"F": 4,
		"G": 2,
		"H": 4,
		"I": 1,
		"J": 8,
		"K": 5,
		"L": 1,
		"M": 3,
		"N": 1,
		"O": 1,
		"P": 3,
		"Q": 10,
		"R": 1,
		"S": 1,
		"T": 1,
		"U": 1,
		"V": 4,
		"W": 4,
		"X": 8,
		"Y": 4,
		"Z": 10,
	}

	for _, scrabbleChar := range input {
		scrabbleScore += scoreMap[strings.ToUpper(string(scrabbleChar))]
	}

	return scrabbleScore
}
