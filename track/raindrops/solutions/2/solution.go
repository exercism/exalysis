// Package raindrops provides a simple int to raindrop-speak conversion function
package raindrops

import "strconv"

// Convert int to raindrop speak
func Convert(n int) string {

	var raindrop string
	if n%3 == 0 {
		raindrop += "Pling"
	}

	if n%5 == 0 {
		raindrop += "Plang"
	}

	if n%7 == 0 {
		raindrop += "Plong"
	}

	if raindrop == "" {
		raindrop += strconv.Itoa(n)
	}

	return raindrop
}
