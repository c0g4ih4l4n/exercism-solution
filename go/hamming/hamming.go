package hamming

import (
	"errors"
)

const testVersion = 6

// Distance : Function get difference
func Distance(a, b string) (int, error) {
	var result int

	if len(a) != len(b) {
		return -1, errors.New("two string doesn't have the same len")
	}

	for i, value := range a {
		if byte(value) != b[i] {
			result++
		}
	}

	return result, nil
}
