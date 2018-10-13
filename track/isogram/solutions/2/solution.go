package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(s string) bool {

	alphaRX, _ := regexp.Compile("[^a-zA-Z]")
	lower := alphaRX.ReplaceAllString(strings.ToLower(s), "")
	chars := make(map[rune]bool)
	isIso := true

	for _, c := range lower {
		_, ok := chars[c]

		if ok {
			isIso = false
			break
		}

		chars[c] = true
	}

	return isIso

}
