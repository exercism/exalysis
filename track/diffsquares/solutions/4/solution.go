package diffsquares

import "math"

// SumOfSquares takes in an int and returns an
// integer number equal to the Sum of all the Squares
// upto that number
func SumOfSquares(num int) int {
	var sum float64
	for i := 1; i <= num; i++ {
		sum += math.Pow(float64(i), 2)
	}
	return int(sum)
}

// SquareOfSum takes in an integer number and returns the
// square of the sum of all numbers upto the provided no.
func SquareOfSum(num int) int {
	var sum int
	for i := 1; i <= num; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

// Difference = SumOfSquares(num) - SquareOfSum(num)
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
