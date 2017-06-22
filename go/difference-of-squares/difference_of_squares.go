package diffsquares

const testVersion = 1

// SquareOfSums : using formula
func SquareOfSums(num int) (result int) {
	return num * num * (num + 1) * (num + 1) / 4
}

// SumOfSquares : using formula
func SumOfSquares(num int) (result int) {
	return num * (num + 1) * (2*num + 1) / 6
}

// Difference :
func Difference(num int) (result int) {
	return SquareOfSums(num) - SumOfSquares(num)
}
