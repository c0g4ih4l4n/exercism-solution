package pascal

const testVersion = 1

// Triangle : Pascal Triable
// two loop : O(n^2)
func Triangle(n int) [][]int {

	result := make([][]int, n)

	// number element more than 1
	if n > 1 {
		// create based
		result[0], result[1] = make([]int, 1), make([]int, 2)
		result[0][0], result[1][0], result[1][1] = 1, 1, 1

		// loop to get higher triangle
		for i := 2; i < n; i++ {
			//create based line
			result[i] = make([]int, i+1)
			result[i][0], result[i][i] = 1, 1

			// get element of line
			for j := 1; j < i; j++ {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	} else if n == 1 {
		// number element is 1
		result[0] = make([]int, 1)
		result[0][0] = 1
	} else {
		// if number = 0 return nil
		return make([][]int, 0)
	}
	return result
}
