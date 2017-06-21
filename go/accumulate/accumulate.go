package accumulate

const testVersion = 1

func Accumulate(strs []string, f func(string) string) (result []string) {
	for _, str := range strs {
		result = append(result, f(str))
	}
	return result
}
