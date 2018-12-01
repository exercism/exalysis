package raindrops

import (
	"strconv"
)

// Convert prints a number into a raindrop-speak  by checking its factors
func Convert(number int) string {
	if number%7 == 0 && number%5 == 0 && number%3 == 0 {
		return "PlingPlangPlong"
	} else if number%5 == 0 && number%3 == 0 {
		return "PlingPlang"
	} else if number%7 == 0 && number%5 == 0 {
		return "PlangPlong"
	} else if number%7 == 0 && number%3 == 0 {
		return "PlingPlong"
	} else if number%7 == 0 {
		return "Plong"
	} else if number%5 == 0 {
		return "Plang"
	} else if number%3 == 0 {
		return "Pling"
	} else {
		return strconv.Itoa(number)
	}
}
