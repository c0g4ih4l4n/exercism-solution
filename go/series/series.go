package series

const testVersion = 2

// All :
func All(n int, str string) (result []string) {
	len := len(str)
	if n > len {
		return
	}

	result = make([]string, len-n+1)

	for i := 0; i < len-n+1; i++ {
		result[i] = str[i : i+n]
	}

	return result
}

// UnsafeFirst :
func UnsafeFirst(n int, str string) (result string) {
	len := len(str)

	if n > len {
		return
	}
	result = str[:n]
	return result
}
