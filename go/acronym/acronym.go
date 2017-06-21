package acronym

/*
import (
	"strings"
	"unicode"
)

const testVersion = 3

// Abbreviate : abbr of a word
func Abbreviate(pharse string) (acronym string) {
	var state bool

	for _, char := range pharse {
		if state {
			if !unicode.IsLetter(char) {
				state = false
			}
		} else {
			if unicode.IsLetter(char) {
				state = true
				acronym += strings.ToUpper(string(char))
			}
		}
	}
	return acronym
}*/
