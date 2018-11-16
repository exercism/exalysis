//Package isogram implements IsIsogram to check input string is an isogram
package isogram

import (
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`[\s-]`)

//IsIsogram function tell us if the input string is an isogram
func IsIsogram(input string) bool {
	input = strings.ToLower(reg.ReplaceAllString(input, ""))
	charmap := make(map[rune]bool)
	for _, c := range input {
		_, present := charmap[c]
		if present {
			return false
		}
		charmap[c] = true
	}
	return true
}
