package diffsquares

import "math"

// Difference between the SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareOfSum is the square of the summation of the first N numbers
func SquareOfSum(n int) int {
	sumOfFirstN := (n * (n + 1)) / 2
	return int(math.Pow(float64(sumOfFirstN), 2))
}

// SumOfSquares is the summation of each of the first N numbers squared
func SumOfSquares(n int) int {
	return ((n * (n + 1)) * (2*n + 1)) / 6
}
