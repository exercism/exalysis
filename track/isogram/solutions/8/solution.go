package isogram

import "unicode"

// IsIsogram checks if string given as input is an isogram (contains only unique characters).
func IsIsogram(input string) bool {
	for index, letter := range input {
		if letter == ' ' || letter == '-' {
			continue
		}
		for _, anotherLetter := range input[index+1:] {
			if unicode.ToLower(letter) == unicode.ToLower(anotherLetter) {
				return false
			}
		}
	}
	return true
}
