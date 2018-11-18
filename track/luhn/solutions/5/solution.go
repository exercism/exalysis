package luhn

import (
	"regexp"
	"strconv"
)

func odd(num int) bool {
	return num%2 == 0
}

func doubleMod9(num int) int {
	if num == 9 {
		return 9
	}
	return (num + num) % 9
}

var whitespace = regexp.MustCompile("\\s")
var notDigits = regexp.MustCompile("\\D")

// Valid checks if a number is a valid per the Luhn firmula
func Valid(number string) bool {
	number = whitespace.ReplaceAllString(number, "")
	length := len(number)
	if notDigits.MatchString(number) || length <= 1 {
		return false
	}

	sum := 0
	for i, ch := range number {
		// At this point the string only contains numbers
		// Errors can be ignored
		digit, _ := strconv.Atoi(string(ch))
		if odd(length - i) {
			sum += doubleMod9(digit)
		} else {
			sum += digit
		}
	}
	return sum%10 == 0
}
