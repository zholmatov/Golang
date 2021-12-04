package main

import "fmt"

func main() {

	// we can pass something from one gorouting into another
	message := make(chan string)

	go func() { message <- "Hello there" }()

	msg := <-message

	fmt.Println(msg)
}
