package raindrops

import (
	"strconv"
)

// Convert turns integers into string based on factor
func Convert(a int) string {

	factor := []int{3, 5, 7}

	three := "Pling"
	five := "Plang"
	seven := "Plong"

	if a%factor[0] == 0 && a%factor[1] == 0 && a%factor[2] == 0 {
		return three + five + seven
	}
	if a%factor[0] == 0 && a%factor[1] == 0 {
		return three + five
	}
	if a%factor[0] == 0 && a%factor[2] == 0 {
		return three + seven
	}
	if a%factor[0] == 0 {
		return three
	}
	if a%factor[1] == 0 && a%factor[2] == 0 {
		return five + seven
	}

	if a%factor[1] == 0 {
		return five
	}
	if a%factor[2] == 0 {
		return seven
	}

	return strconv.Itoa(a)
}
