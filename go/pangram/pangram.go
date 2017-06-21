package pangram

import "strings"

const testVersion = 1

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for _, letter := range alphabet {
		if !strings.ContainsRune(s, letter) {
			return false
		}
	}
	return true
}
