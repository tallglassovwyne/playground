package differenceofsquares

// SumOfSquares computes the sum of the squares of numbers up to n.
func SumOfSquares(n int) int {
	var total int = 0

	for i := 1; i <= n; i++ {
		total += i * i
	}

	return total
}

// SquareOfSum computes the square of numbers up to n.
func SquareOfSum(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += i
	}

	return total * total
}

// Difference computes the difference between the square of numbers up to n and the sum of the squares of numbers up to n.
func Difference(n int) int {
	sumOfSquare, squareOfSum:= SumOfSquares(n), SquareOfSum(n)

	return squareOfSum - sumOfSquare
}