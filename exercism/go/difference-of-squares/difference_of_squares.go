package diffsquares

import "math"

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	
	return ((2 * n + 1) * (n + 1 ) * n) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
