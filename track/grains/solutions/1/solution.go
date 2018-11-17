/*
Package grains calculate the number of grains of wheat on a chessboard given that the number on each square doubles.
*/
package grains

import (
	"errors"
	"math"
	"math/bits"
)

// Square returns number of grains on the specified square
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid number of squares")
	}
	return uint64(bits.RotateLeft(1, number-1)), nil
}

// Total returns total number of grains from 64 squares (2^64 - 1)
func Total() uint64 {
	return math.MaxUint64
}
