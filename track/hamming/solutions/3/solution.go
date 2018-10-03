package hamming

import "errors"

// Distance calculates the Hamming distance between two strands of DNA represented as Sring.
func Distance(a string, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Unable to calculate distance on strings of unequal length")
	}

	d := 0
	bs := []rune(b)
	for pos, char := range a {
		if char != bs[pos] {
			d = d + 1
		}
	}
	return d, nil
}
