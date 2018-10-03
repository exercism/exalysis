package hamming

import "errors"

// Distance returns the Hamming distance of 2 strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings are different lengths")
	}

	diff := 0
	for index, char := range a {
		if byte(char) != b[index] {
			diff++
		}
	}

	return diff, nil
}
