package hamming

import (
	"errors"
)

func Distance(a, b string) (distance int, err error) {
	r1 := []rune(a)
	r2 := []rune(b)

	if len(r1) != len(r2) {
		err = errors.New("Different size of strings.")
		return
	}

	for i, v := range r1 {
		if r2[i] != v {
			distance++
		}
	}
	return
}
