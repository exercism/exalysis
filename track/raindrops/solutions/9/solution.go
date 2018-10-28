// Package raindrops convert a number to a string, the contents of which depend on the number's factors.
package raindrops

import "strconv"

// Convert converts a number to a string, the contents of which depend on the number's factors.
func Convert(i int) string {
	var result string
	for n := 3; n <= 7; n++ {
		if i%n == 0 {
			switch n {
			case 3:
				result = result + "Pling"
			case 5:
				result = result + "Plang"
			case 7:
				result = result + "Plong"
			}
		}
	}
	if result == "" {
		return strconv.Itoa(i)
	}
	return result
}
