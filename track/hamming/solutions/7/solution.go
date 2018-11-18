package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("")
	}
	dif := 0
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			dif++
		}
	}
	return dif, nil
}
