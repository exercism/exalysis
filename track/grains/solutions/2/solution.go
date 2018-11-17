package grains

import (
	"fmt"
	"math"
)

// Square returns the number of  grains on a square on a chess board where the
// first square has 1 and every subsequent square doubles the number.
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, fmt.Errorf("Num [%d] is invalid.", num)
	}
	return uint64(math.Pow(2, float64(num-1))), nil
}

// Total returns the total number of grains of the chessboard.
func Total() uint64 {
	// uint64's range: 0-18446744073709551615.
	return uint64(math.Pow(2, 63))*2 - 1
}
