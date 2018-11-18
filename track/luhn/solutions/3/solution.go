//Package luhn contains a single function Valid
//and a couple of internal functions used for intermediate
//calculations
package luhn

//Valid will take a string as input and if the string is
//of length greater than 2 and contains only digits and spaces
//it will return a boolean value based on its validity per the Luhn
//formula. Any inputs that do not meet the requirements are false
func Valid(s string) bool {
	counter, total := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])

		if r == ' ' {
			continue
		} else if r < '0' || r > '9' {
			return false
		} else {
			total += calculateValue(counter, r)
			counter++
		}
	}

	if counter < 2 {
		return false
	}
	return isEvenlyDivisibleByTen(total)
}

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

func isEvenlyDivisibleByTen(val int) bool {
	return val%10 == 0
}
