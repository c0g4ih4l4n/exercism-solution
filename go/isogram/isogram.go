package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram :
func IsIsogram(in string) bool {
	var listChar = make(map[rune]int)

	in = strings.ToLower(in)
	for _, r := range in {
		// if r is not a letter
		// continue
		if !unicode.IsLetter(r) {
			continue
		}

		// check in map
		// if exists false
		// if not add it to map
		_, ok := listChar[r]
		if ok {
			return false
		}
		listChar[r]++
	}
	return true
}
