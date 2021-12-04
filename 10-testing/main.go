package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(100.0, 0.0)

	if err != nil {
		log.Println("Error while dividing,", err)
		return
	}

	log.Println("result of division is", result)
}

// takes two parameters of type float32, and returns float32 and error, if there is any
func divide(x, y float32) (float32, error) {

	var result float32 // -> if its not set, then it returns zero

	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y

	return result, nil

}
