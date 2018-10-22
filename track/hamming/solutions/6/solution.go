package hamming

import "errors"

// Distance function calculates the hamming distance between two DNA strings
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return -1, errors.New("unequal distance")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil

}
