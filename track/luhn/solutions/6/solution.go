// Package luhn provides calculation of the Luhn checksum.
package luhn

import "unicode"

// Valid checks whether a string contains a number valid per the Luhn algorithm.
func Valid(input string) bool {
	sum := 0
	digitIndex := 0
	for n := len(input) - 1; n >= 0; n-- {
		char := rune(input[n])
		if char == ' ' {
			continue
		}
		if !unicode.IsDigit(char) {
			return false
		}
		value := int(char - '0')
		if (digitIndex)%2 == 0 {
			sum += value
		} else {
			value2 := value * 2
			if value2 > 9 {
				sum += value2 - 9
			} else {
				sum += value2
			}
		}
		digitIndex++
	}
	if digitIndex <= 1 {
		return false
	}
	return sum%10 == 0
}
