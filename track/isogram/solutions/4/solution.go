// Package isogram tests for whether or not a word is an isogram
package isogram

import (
	"unicode"
)

// IsIsogram returns true iff input is an isogram (word without repeating letters).
func IsIsogram(candidate string) bool {
	var used = make([]bool, 26)
	for _, c := range candidate {
		if unicode.IsLetter(c) {
			cUpper := unicode.ToUpper(c) - 'A'
			if used[cUpper] {
				return false
			}
			used[cUpper] = true
		}
	}
	return true
}
