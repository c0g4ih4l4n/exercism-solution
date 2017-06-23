package summultiples

import "fmt"

const testVersion = 2

// SumMultiples :
func SumMultiples(limit int, divisors ...int) (sum int) {

	//ch3, ch5, ch15 := make(chan int), make(chan int), make(chan int)

	for divisor := range divisors {
		fmt.Println(divisor)
	}
	return sum
}
