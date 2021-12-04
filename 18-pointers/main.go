package main

import "fmt"

// func zeroVal(num int) {
// 	num = 0
// }

func zeroPtr(num *int) {
	*num = 0 // deferences the pointer of num to 0
}

func main() {
	var i int = 1 // it is the pointee

	// zeroVal(i)
	// fmt.Println("zeroVal", i)
	fmt.Println("memory address", &i)

	zeroPtr(&i)               // &i is the memory address of i
	fmt.Println("zeroPtr", i) // the memory address which was pointing to 1, now points to 0
	// thus, it prints 0 now

	fmt.Println("memory address", &i) // pointer of i, i.e memory address
}
