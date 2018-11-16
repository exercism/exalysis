// Package diffsquares provides functions for computing aggregate values on slices
package diffsquares

// SquareOfSum iteratively computes the square of sum of numbers from 1 to n
func SquareOfSum(n int) (res int) {
	res = 0
	for i := 1; i < n+1; i++ {
		res += i
	}
	return res * res
}

// SumOfSquares iteratively computes the sum of squares of numbers from 1 to n
func SumOfSquares(n int) (res int) {
	res = 0
	for i := 1; i < n+1; i++ {
		res += i * i
	}
	return
}

// Difference computes the difference between the square of sums and the sum of squares
func Difference(n int) int{
	return SquareOfSum(n) - SumOfSquares(n)
}
