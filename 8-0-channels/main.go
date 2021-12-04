package main

// with this import we can use another functions in other packages
import (
	"log"

	"github.com/zholmatov/Golang/8-0-channels/helpers"
)

const numPool = 10

// if func is Camel case, it can be used in other files
func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber // ---> passed value to the channel
}

func main() {
	// creates a channel that can send
	// information. (only integer)
	intChan := make(chan int)
	defer close(intChan)
	// defer close() means to close the channel as soon as it is finished running

	// goroutines all run at the same time
	// and to run that, we can do:
	go CalculateValue(intChan)

	// we are listening to that channel down here
	num := <-intChan // ---> assign num with whatever comes from intChan

	log.Println(num)

}
