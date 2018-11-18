// Package luhn calculates Luhn checksum
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

//Valid check the string is valid number
func Valid(N string) bool {
	var a string
	for _, v := range N {
		if unicode.IsLetter(v) || strings.Contains(N, "-") || unicode.IsSymbol(v) {
			return false
		}
		if !unicode.IsNumber(v) {
			continue
		}
		a += string(v)
	}
	for i := len(a) - 1; 0 < i; i -= 2 {
		b, _ := strconv.Atoi(a[i-1 : i])
		b *= 2
		if b > 9 {
			b -= 9
		}
		a = a[:i-1] + strconv.Itoa(b) + a[i:]
	}
	var sum int
	for _, v := range a {
		c, _ := strconv.Atoi(string(v))
		sum += c
	}
	valid := true
	if len(a) < 2 || sum%10 != 0 {
		valid = false
	}
	return valid
}
