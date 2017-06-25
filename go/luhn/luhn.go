package luhn

import (
	"unicode"
)

const testVersion = 2

// use one loop to calculate all

// Valid :
func Valid(code string) bool {

	// change code to rune
	text := []rune(code)

	// get len to loop
	// sum  to calculate
	// plusTime to check number of numbers in text
	// dbl to check whenever we double before plus
	len := len(text)
	sum := 0
	plusTime := 0
	dbl := false

	// loop
	for i := len - 1; i >= 0; i-- {

		// some check
		if text[i] == ' ' {
			continue
		}
		if unicode.IsLetter(text[i]) {
			return false
		}

		// calculate
		num := int(text[i] - '0')
		if dbl {
			if num*2 > 9 {
				sum += num*2 - 9
			} else {
				sum += num * 2
			}
		} else {
			sum += num
		}
		dbl = !dbl
		plusTime++
	}

	if plusTime > 1 && sum%10 == 0 {
		return true
	}

	return false
}
