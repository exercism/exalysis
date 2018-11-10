package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(num int) string {
	// Store the requisite Strings as constant variables
	const pling = "Pling"
	const plang = "Plang"
	const plong = "Plong"

	// I opted to use a Byte Buffer to build my string, there are other ways
	var buffer bytes.Buffer

	// This flag handles whether or not we found a factor were looking for
	// This is done to handle the requirement that if the number does not have a factor, we return the string version of that number
	foundName := false

	// We iterate from 1 to the number, checking each number to see if it is a factor
	for i := 1; i <= num; i++ {

		if num%i == 0 {
			switch i {
			case 3:
				buffer.WriteString(pling)
				foundName = true
			case 5:
				buffer.WriteString(plang)
				foundName = true
			case 7:
				buffer.WriteString(plong)
				foundName = true
			}
		}
	}
	//If we did not find the factors, then we just return the string version of the number
	if !foundName {
		return strconv.Itoa(num)
	}

	// Otherwise, we convert the buffer to a string and return it
	return buffer.String()

}
