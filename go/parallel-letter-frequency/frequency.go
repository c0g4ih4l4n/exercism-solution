package letter

const testVersion = 1

// FreqMap : map hold rune frequency
type FreqMap map[rune]int

// Frequency :
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Solution 2: using many param to save result

// ConcurrentFrequency :
func ConcurrentFrequency(arrStr []string) FreqMap {

	result := FreqMap{}
	len := len(arrStr)
	// bufferred channel to save result
	ch := make(chan FreqMap, 3)

	for _, str := range arrStr {
		go func(s string) {
			ch <- Frequency(s)
		}(str)
	}

	for i := 0; i < len; i++ {
		for r, value := range <-ch {
			result[r] += value
		}
	}

	return result
}
