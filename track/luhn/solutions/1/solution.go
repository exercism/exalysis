//Package luhn contains a single function Valid
//and a couple of internal functions used for intermediate
//calculations
package luhn

import (
	"regexp"
)

//Valid will take a string as input and if the string is
//of length greater than 2 and contains only digits and spaces
//it will return a boolean value based on its validity per the Luhn
//formula. Any inputs that do not meet the requirements are false
func Valid(str string) bool {
	s := regexp.MustCompile("\\s").ReplaceAllString(str, "")
	validRegex := regexp.MustCompile("^[0-9]{2,}$")
	isInputValid := validRegex.MatchString(s)
	total := 0

	if !isInputValid {
		return false
	}

	lengthOffset := (len(s) - 1) % 2
	for i, r := range s {
		total += calculateValue(i+lengthOffset, r)
	}
	return evenlyDivisbleByTen(total)
}

//calculateValue will return double the int value from
//the input rune for every second digit. If the value is greater than 9, 9 is subtracted. Just the int value is returned otherwise
func calculateValue(i int, r rune) int {
	val := int(r) - 48
	if i%2 == 0 {
		return val
	}
	val *= 2
	if val > 9 {
		return val - 9
	}
	return val
}

//evenlyDivisbleByTen returns boolean value of true
//if the int passed is evenly divisible by 10.
func evenlyDivisbleByTen(val int) bool {
	return val%10 == 0
}
