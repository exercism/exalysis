package diffsquares

//SquareOfSum returns the Square of sum of the natural numbers from 1 to input
func SquareOfSum(input int) int {
	out := 0
	for i := 1; i <= input; i++ {
		out += i
	}
	return out * out
}

//SquareOfSum returns the sum of the squares of natural numbers from 1 to input
func SumOfSquares(input int) int {
	var out int
	for i := 1; i <= input; i++ {
		out += i * i
	}
	return out
}

//Difference returns the difference between sum of squares and Square of sum of input
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
