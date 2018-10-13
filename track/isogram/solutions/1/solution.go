package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks whether a string is an isogram or not
func IsIsogram(word string) bool {
	letters := make(map[rune]bool, len(word))

	for _, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) && letters[r] {
			return false
		}
		letters[r] = true
	}
	return true
}
