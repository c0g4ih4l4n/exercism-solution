package etl

import (
	"strings"
)

const testVersion = 1

// Transform :
func Transform(input map[int][]string) (result map[string]int) {

	result = make(map[string]int)
	for i, values := range input {
		for _, value := range values {
			result[strings.ToLower(value)] = i
		}
	}
	return result
}
