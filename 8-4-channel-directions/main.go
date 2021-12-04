package main

import "fmt"

func main() {

	ping := make(chan string)
	pong := make(chan string)

	// it takes this Message received and sends it to another channel, i.e pong
	go receiveMessage(ping, "Message received")
	go receiveAndSendMessage(ping, pong)

	fmt.Println(<-pong)
}

// if chan<-  this channel only receives a message
func receiveMessage(ping chan<- string, message string) {
	ping <- message
}

// here ping can only send message and pong can only receive
func receiveAndSendMessage(ping <-chan string, pong chan<- string) {
	// here message is passed from ping channel
	message := <-ping
	// here pong receives message
	pong <- message
}
