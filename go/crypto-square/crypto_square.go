package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// Encode :
func Encode(plaintext string) string {

	plaintext = normalize(plaintext)
	// variable of program
	var chunk [][]rune
	len := len(plaintext)
	var r, c int
	var ciphertext []rune

	// get row and col of chunk
	r1 := math.Sqrt(float64(len))

	// if r1 is sqrt of len
	if r1 == float64(int(r1)) {
		r, c = int(r1), int(r1)
	} else {
		// else take floor of r1
		r2 := int(math.Floor(r1))

		// if r2 * (r2 + 1) > len
		// it avaiable take it
		if r2*(r2+1) > len {
			r, c = r2, r2+1
		} else {
			// take next sqrt num
			r, c = r2+1, r2+1
		}
	}

	chunk = make([][]rune, r)
	for i := 0; i < r; i++ {
		chunk[i] = make([]rune, c)
	}

	// paste string to chunk
	line, col := 0, 0
	for _, v := range plaintext {
		chunk[line][col] = v

		// change line and col to the next space in chunk
		if col == c-1 {
			line++
			col = 0
			continue
		}
		col++
	}

	// get ciphertext from chunk
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if chunk[j][i] == 0 {
				break
			}
			ciphertext = append(ciphertext, chunk[j][i])
		}
		if i != c-1 {
			ciphertext = append(ciphertext, ' ')
		}
	}

	return string(ciphertext)
}

func normalize(plaintext string) string {

	var result []rune
	plaintext = strings.ToLower(plaintext)
	r := []rune(plaintext)

	for _, v := range r {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			result = append(result, v)
		}
	}

	return string(result)
}
