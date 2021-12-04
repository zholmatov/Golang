package main

import (
	"fmt"
	"log"
)

func main() {

	mySlice := []string{"Uli", "Doni"}
	for _, x := range mySlice {
		log.Println("name is", x)
	}

	numsMap := make(map[int]int)

	numsMap[1] = 0
	numsMap[2] = 3

	findIndex := numsMap[1]
	findIndex2 := numsMap[3]
	fmt.Println(findIndex)
	fmt.Println(findIndex2)

}
