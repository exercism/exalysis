package collatzconjecture

import (
	"errors"
)

// CollatzConjecture - Given a number n, returns the number of steps required to reach 1.
func CollatzConjecture(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("`n` can not be less than 1")
	}

	var steps int

	for number != 1 {
		steps++

		if number%2 == 0 {
			number /= 2
			continue
		}

		number *= 3
		number++
	}

	return steps, nil
}
