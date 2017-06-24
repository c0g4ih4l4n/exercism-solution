package main

import "fmt"

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	fmt.Println(sample)
	fmt.Println()

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%+q\n", sample[i])
	}
}
