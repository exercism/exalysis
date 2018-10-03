package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return -1, errors.New("Invalid sizes")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count, nil
	}

}
