package queenattack

import (
	"errors"
)

const testVersion = 2

// CanQueenAttack :
func CanQueenAttack(white, black string) (result bool, err error) {

	if len(white) != 2 || len(black) != 2 || white == black {
		return result, errors.New("Input same or too long")
	}

	whiteRow, whiteCol, blackRow, blackCol := white[0],
		white[1], black[0], black[1]

	if checkValid(whiteRow, whiteCol) || checkValid(blackRow, blackCol) {
		return false, errors.New("Position not valid")
	}

	if whiteRow == blackRow || blackCol == whiteCol ||
		(whiteCol-whiteRow) == (blackCol-blackRow) ||
		(whiteCol+whiteRow) == (blackCol+blackRow) {
		return true, nil
	}
	return false, nil
}

func checkValid(row byte, col byte) bool {
	if row < 'a' || row > 'h' || col > '8' || col < '1' {
		return true
	}
	return false
}
