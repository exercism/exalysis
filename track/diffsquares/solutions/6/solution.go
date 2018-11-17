package diffsquares

import "math"

//SquareOfSum counts second power of sum of first n natural numbers (1 + 2 + ... + n)²
func SquareOfSum(number int) int {
	numberSum := sum(number, false)
	return int(math.Pow(float64(numberSum), float64(2)))
}

//SumOfSquares counts sum of second power of first n natural numbers (1² + 2² + ... + n²)
func SumOfSquares(number int) int {
	powSum := sum(number, true)
	return powSum
}

//Difference counts difference between SquareOfSum and SumOfSquares
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}

func sum(number int, usePow bool) int {
	result := 0
	for i := 1; i <= number; i++ {
		if usePow {
			result += int(math.Pow(float64(i), float64(2)))
		} else {
			result += i
		}
	}
	return result
}
