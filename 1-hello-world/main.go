package main

import "log"

// every go file has to have a package
// and at least one function

func main() {
	// to create a variable, you should write var first
	// then type of that variable

	// variables video code
	// var a string = saySomething("Hello")
	// log.Println(a)

	var red string = "Red"

	log.Println(red)

	pointToSmth(&red)

	log.Println(red)
}

// after declaring the type of input,
// you should declare the type of return
// as well
func saySomething(s string) string {
	return s
}

func pointToSmth(s *string) {
	green := "Green"
	*s = green
}
