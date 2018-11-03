// Package raindrops is used to transform a number into the lovely sound of rain
package raindrops

import "fmt"

// Convert converts an integer into a raindrop-speak string
func Convert(num int) string {
	convertedString := ""

	if num%3 == 0 {
		convertedString = convertedString + "Pling"
	}

	if num%5 == 0 {
		convertedString = convertedString + "Plang"
	}

	if num%7 == 0 {
		convertedString = convertedString + "Plong"
	}

	if convertedString == "" {
		return fmt.Sprintf("%d", num)
	}

	return convertedString
}
