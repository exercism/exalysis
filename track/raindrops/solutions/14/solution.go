// Package raindrops converts ints into onomatopoeia depending on factors of that int.
package raindrops

import "fmt"

// Convert takes an int and returns a string containing either the original int, or a series of sounds based on factors of that int.
func Convert(n int) string {
	res := ""
	if hasFactor(n, 3) {
		res = "Pling"
	}
	if hasFactor(n, 5) {
		res = fmt.Sprintf("%sPlang", res)
	}
	if hasFactor(n, 7) {
		res = fmt.Sprintf("%sPlong", res)
	}
	if len(res) == 0 {
		res = fmt.Sprintf("%d", n)
	}

	return res
}

func hasFactor(n, f int) bool {
	return n%f == 0
}
