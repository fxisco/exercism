package diffsquares

func SumOfSquares(n int) (sumOfSq int) {
	for i := 1; i <= n; i++ {
		sumOfSq += (i * i)
	}

	return
}

func SquareOfSum(n int) int {
	sum := float64((n * (n + 1)) / 2)
	sqOfSum := int(sum * sum)

	return sqOfSum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
