package luhn

/*
import (
	"unicode"
)

const testVersion = 2

// do 3 step:
// 1: trim code
// 2: increase value of every second digit
// 3: calculate sum and return result check

// Valid :
func Valid(code string) bool {

	// Trim code and get array nums
	nums, ok := getArrNum(code)

	if !ok {
		return false
	}

	// increase code
	len := len(nums)
	dbl := false
	for i := len - 1; i >= 0; i, dbl = i-1, !dbl {
		if dbl {
			if nums[i]*2 > 9 {
				nums[i] = nums[i]*2 - 9
			} else {
				nums[i] = nums[i] * 2
			}
		}
	}

	// get sum
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

func getArrNum(code string) (result []int, ok bool) {

	text := []rune(code)

	for _, c := range text {
		if c == ' ' {
			continue
		}
		if unicode.IsLetter(c) {
			return nil, false
		}
		result = append(result, int(c-'0'))
	}

	if len(result) == 1 {
		return nil, false
	}

	return result, true
}
*/
