package diffsquares

import (
	"math"
)

func SumOfSquares(n int) int {
	var sumOfSq int

	for i := 1; i <= n; i++ {
		sumOfSq += int(math.Pow(float64(i), 2.0))
	}

	return sumOfSq
}

func SquareOfSum(n int) int {
	sum := float64((n * (n + 1)) / 2)
	sqOfSum := int(math.Pow(sum, 2.0))

	return sqOfSum
}

func Difference(n int) int {
	difference := float64(SumOfSquares(n) - SquareOfSum(n))

	return int(math.Abs(difference))
}
