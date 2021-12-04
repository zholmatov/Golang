package main

import "fmt"

func fact(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * fact(num-1)
	}
}

func main() {
	fact7 := fact(7)

	// declare variable function fib which returns int
	// if you declare functions this way, you can change it
	var fib func(num int) int

	// write that function down here
	fib = func(num int) int {
		if num < 2 {
			return num
		} else {
			return fib(num-1) + fib(num-2)
		}
	}
	fib7 := fib(7)
	fmt.Println(fact7)
	fmt.Println(fib7)

	fib = func(num int) int {
		return num + 1
	}

	fib8 := fib(7)
	fmt.Println(fib8)
}
