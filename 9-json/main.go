package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_name": "John",
			"last_name": "Wick",
			"has_dog": true
		}, 
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"has_dog": false
		}
	]
	`

	// give slice(array) type Person, many Persons
	var unmarshalled []Person

	// means put myJson into unmarshalled  & -> this sign should be used
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	// this prints out the error if there is any
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from struct:
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Doni"
	m1.LastName = "Zholmat"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Sabyr"
	m2.LastName = "Orozbek"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	// MarshalIndent gives json nice format, -->
	// [
	// 	{
	// 		"first_name": "Doni",
	// 		"last_name": "Zholmat",
	// 		"has_dog": false
	// 	},
	// 	{
	// 		"first_name": "Sabyr",
	// 		"last_name": "Orozbek",
	// 		"has_dog": true
	// 	}
	// ]

	// but Marshal does it as smaall is it could ---> [{"first_name":"Doni","last_name":"Zholmat","has_dog":false},{"first_name":"Sabyr","last_name":"Orozbek","has_dog":true}]
	newJson, err := json.MarshalIndent(mySlice, "", "   ") // first "" is prefix, the second one is indent

	if err != nil {
		log.Println("Error marshalling json", err)
	}

	// this gives it in a string format
	fmt.Println(string(newJson))

}
