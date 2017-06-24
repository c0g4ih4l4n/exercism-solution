package summultiples

const testVersion = 2

// SumMultiples :
func SumMultiples(limit int, divisors ...int) (sum int) {

	if len(divisors) == 0 {
		return
	}

	for i := 1; i < limit; i++ {
		// loop through divisors
		// if one mod == 0
		// + that i and break
		// to make sure one i plus only one time
		for _, value := range divisors {
			if i%value == 0 {
				sum += i
				break
			}
		}
	}

	return sum
}
