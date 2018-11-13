package isogram

import (
	"strings"
	"unicode"
)

// Score returns scrabble score for given word
func IsIsogram(word string) bool {
	foundLetters := ""
	punctuation := " -"
	for _, letter := range word {

		letter := string(unicode.ToLower(letter))

		if strings.Index(punctuation, letter) == -1 && strings.Index(foundLetters, letter) > -1 {
			return false
		}
		foundLetters += letter

	}

	return true
}
