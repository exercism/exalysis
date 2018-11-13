package isogram

import "unicode"

/*
IsIsogram  module checks whether string has repeated characters
*/
func IsIsogram(input string) bool {
	isoMap := map[rune]bool{}

	for _, ch := range input {
		ch = unicode.ToLower(ch)
		if ch >= 'a' && ch <= 'z' {
			if _, ok := isoMap[ch]; ok {
				return false
			}
			isoMap[ch] = true
		}
	}
	return true
}
