package main

import "log"

func main() {

	// map[firstType]secondType
	// this is how you make objects in go
	myMap := make(map[string]string)

	myMap["dog"] = "Reks"
	myMap["cat"] = "Leo"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])
}
