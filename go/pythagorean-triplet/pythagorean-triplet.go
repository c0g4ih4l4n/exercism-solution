package pythagorean

import (
	"math"
)

const testVersion = 1

// Triplet :
type Triplet [3]int

// Range :
func Range(min, max int) (result []Triplet) {

	num := 0

	// using some condition to decrease loop
	for a := min; a < max; a++ {
		for b := a + 1; b <= int(math.Floor(math.Sqrt(float64(max*max-a*a)))); b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if c-math.Floor(c) == 0 {
				result = append(result, Triplet{a, b, int(c)})
				num++
			}
		}
	}
	return
}

// Sum :
func Sum(sum int) (result []Triplet) {

	// using fomula to speed calculate
	for a := 1; a <= sum/2; a++ {
		if b := (float64(sum) + float64(a*sum)/float64(a-sum)) / 2.0; a < int(b) && b == float64(int(b)) {
			result = append(result, Triplet{a, int(b), int(math.Sqrt(float64(a*a) + b*b))})
		}
	}
	return
}
