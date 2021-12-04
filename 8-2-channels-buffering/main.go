package main

import "fmt"

func main() {

	// if you add 2 at the end of making channel, you can pass 2 strings,
	// and not only in goroutine function
	message := make(chan string, 2)

	// this also works
	// go func() {
	// 	message <- "First message"
	// 	message <- "Second message"
	// }()

	// this one works too
	message <- "First message"
	message <- "Second message"

	fmt.Println(<-message)
	fmt.Println(<-message)

}
