// Package luhn checks the validity of creditcard numbers
// and canadian social security numbers by applying
// the "Luhn" algorithm.
package luhn

import (
	"regexp"
	"strings"
)

// Valid validates a given string by the means stated in the
// package description.
func Valid(s string) bool {
	// Remove spaces.
	o := strings.Replace(s, " ", "", -1)
	// Check if the remaining string is at lest 2 chars long.
	if len(o) <= 1 {
		return false
	}
	// Check for non-digits.
	if ok, _ := regexp.MatchString(`((?s)[\D[:punct:]])+`, o); ok {
		return false
	}

	// Algorithm violently torn off of
	// http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers#Go
	t := []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	odd := len(o) & 1
	var sum int
	for k, v := range o {
		if k&1 == odd {
			sum += t[v-'0']
		} else {
			sum += int(v - '0')
		}
	}

	return sum%10 == 0

}
