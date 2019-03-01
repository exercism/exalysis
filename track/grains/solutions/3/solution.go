package grains

import (
	"errors"
	"math"
)

//Square returns the number of grains on the input square
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("Invalid input")
	}
	result := uint64(math.Exp2(float64(input - 1)))
	return result, nil
}

//Total returns all the grains on the board.
func Total() uint64 {
	sum := uint64(0)
	for i := 1; i <= 64; i++ {
		sum1, _ := Square(i)
		sum += sum1
	}
	return sum
}
