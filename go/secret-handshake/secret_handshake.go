package secret

const testVersion = 2

// Handshake :
// given code return string compatiable with that code
func Handshake(code uint) []string {

	secretKey := []string{"wink", "double blink", "close your eyes", "jump"}

	resultSlice := make([]string, 4)
	num := 0
	resultLen := 0
	for code != 0 {
		if code%2 != 0 {
			if num > 4 {
				break
			} else if num == 4 {
				resultSlice = reverse(resultSlice, resultLen)
				break
			}
			resultSlice[resultLen] = secretKey[num]
			resultLen++
		}
		code = code / 2
		num++
	}
	return resultSlice[:resultLen]
}

// reverse : when fouth digit is 1
func reverse(arr []string, resultLen int) []string {
	for i, j := 0, resultLen-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
