//Package raindrops converts number to string and outputs info depending on numbers factors
package raindrops

import (
	"strconv"
)

//Convert converts number and returns Pling, Plang, Plong or number
func Convert(number int) (output string) {
	var factors []int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
		if i == 7 {
			break
		}
	}

	ppp := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	for f := 0; f < len(factors); f++ {
		for k, v := range ppp {
			if factors[f] == k {
				output = output + v
			}
		}
	}
	if output == "" {
		output = strconv.Itoa(number)
	}

	return output
}
