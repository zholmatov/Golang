package main

import "fmt"

// define function which returns function of int
func intSeq() func() int {
	i := 0

	// return anonymous function of int
	return func() int {
		i += 1
		return i
	}
}

func main() {
	newInt := intSeq()

	fmt.Println(newInt()) // 1
	fmt.Println(newInt()) // 2
	fmt.Println(newInt()) // 3

	newInt = intSeq()
	fmt.Println(newInt()) // 1
}
