package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool, 1)

	go work(done)

	// we can put <- done here to notify another goroutine
	// to wait until this function is done
	<-done
}

func work(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}
