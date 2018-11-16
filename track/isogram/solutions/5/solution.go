package isogram

import "unicode"

// IsIsogram returns true if the given word contains at most one instance of
// any letter. Spaces and hyphens are ignored.
func IsIsogram(word string) bool {
	seen := make(map[rune]struct{})
	for _, letter := range word {
		if letter == '-' || letter == ' ' {
			continue
		}
		letter = unicode.ToLower(letter)
		if _, found := seen[letter]; found {
			return false
		}
		seen[letter] = struct{}{}
	}
	return true
}
