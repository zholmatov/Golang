package main

import (
	"fmt"
	"time"
)

// select in go makes you wait multiple channel operations
func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "First message"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Second message"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}
}
