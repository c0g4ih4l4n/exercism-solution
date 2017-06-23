package letter

/*
import (
	"sync"
)

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

// SafeFreqMap :
type SafeFreqMap struct {
	m   FreqMap
	mux sync.Mutex
}

// Solution 1: using one param to save result

// ConcurrentFrequency :
func ConcurrentFrequency(arrStr []string) FreqMap {

	//c := make(chan FreqMap)
	//var wg sync.WaitGroup
	len := len(arrStr)

	safeMap := SafeFreqMap{m: FreqMap{}}
	quit := make(chan int)
	for _, str := range arrStr {
		//wg.Add(1)
		//go frequency(str, &safeMap)
		go func(str string) {
			//defer wg.Done()
			for _, r := range str {
				safeMap.Inc(r)
			}
			quit <- 0
		}(str)
	}

	for i := 0; i < len; i++ {
		<-quit
	}

	//wg.Wait()
	return safeMap.GetMap()
}

// Inc : increase rune frequency
func (f *SafeFreqMap) Inc(r rune) {
	f.mux.Lock()
	f.m[r]++
	f.mux.Unlock()
}

// GetMap :
func (f *SafeFreqMap) GetMap() FreqMap {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.m
}
*/
