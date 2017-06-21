package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	var result, state int
	state = 1

	if len(a) != len(b) {
		return -1, errors.New("Two string doesn't have the same len!")
	}

	for i := 0; i < len(a); i++ {
		switch state {
		case 1:
			switch {
			case i == len(a):
				return result + len(b) - len(a), nil
			case i == len(b):
				state = 2
				result++
			case b[i] != a[i]:
				result++
			default:
			}
		case 2:
			result++
		}
	}

	return result, nil
}
