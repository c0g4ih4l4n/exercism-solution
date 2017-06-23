package grains

import (
	"errors"
)

const testVersion = 1

// Square :
func Square(num int) (result uint64, err error) {

	if num <= 0 || num > 64 {
		return 0, errors.New("Number not valid")
	}

	return uint64(1 << uint(num-1)), nil
}

// Total : Max grains on table
// float64 reach max need to change to uint64
// to continue calculate
func Total() uint64 {
	return uint64(1<<64 - 1)
}

// DynamicSquare : using static var result
func DynamicSquare(num int) func(num int) uint64 {
	var result uint64
	result = 1
	return func(num int) uint64 {
		result = result * 2
		return result
	}
}
