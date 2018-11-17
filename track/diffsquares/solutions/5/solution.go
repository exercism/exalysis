package diffsquares

// Difference finds the difference between the square of the sum and the sum of the squares of the first N natural
// numbers
func Difference(n int) int {
	var square = 0
	var squareOfSum = 0
	for i := 1; i <= n; i++ {
		squareOfSum += i
		square += i * i
	}
	squareOfSum = squareOfSum * squareOfSum
	return squareOfSum - square
}

// SumOfSquares calculates the sum of square of the given n natural numbers
func SumOfSquares(n int) int {
	var square = 0
	for i := 1; i <= n; i++ {
		square += i * i
	}
	return square
}

// SquareOfSum calculates the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	var squareOfTheSum = 0
	for i := 1; i <= n; i++ {
		squareOfTheSum += i
	}
	return squareOfTheSum * squareOfTheSum
}
