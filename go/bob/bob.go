package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Hey : function response
func Hey(text string) (response string) {

	text = strings.Trim(text, " ")
	form := make(map[string]bool)
	var switchYell bool
	form["question"] = false
	form["yell"] = false
	form["silience"] = true
	form["somethingelse"] = false

	if len(text) == 0 {
		return "Fine. Be that way!"
	}

	if text[len(text)-1] == '?' {
		form["question"] = true
	}
	for _, value := range text {
		if !switchYell && !form["yell"] && unicode.IsUpper(value) {
			if unicode.IsNumber(value) {
				continue
			}
			form["yell"] = true
		}
		if form["yell"] && unicode.IsLower(value) {
			if unicode.IsNumber(value) {
				continue
			}
			switchYell = true
			form["yell"] = false
		}

		if form["silience"] && !unicode.IsSpace(value) {
			form["silience"] = false
		}

	}

	if form["question"] && form["yell"] {
		form["question"] = false
	}

	if !(form["question"] || form["yell"] || form["silience"]) {
		form["somethingelse"] = true
	}

	for i, value := range form {
		if value {
			switch i {
			case "yell":
				return "Whoa, chill out!"
			case "question":
				return "Sure."
			case "silience":
				return "Fine. Be that way!"
			case "somethingelse":
				return "Whatever."
			default:
			}
		}
	}
	return ""
}
