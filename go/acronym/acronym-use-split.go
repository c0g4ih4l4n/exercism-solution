package acronym

import (
	"strings"
)

const testVersion = 3

func Abbreviate(pharse string) (acronym string) {
	words := strings.FieldsFunc(pharse, split)

	for _, word := range words {
		acronym = acronym + strings.ToUpper(string(word[0]))
	}
	return acronym
}

func split(r rune) bool {
	return r == ' ' || r == '-'
}
