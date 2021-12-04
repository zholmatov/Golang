package main

import "log"

func main() {
	var firstName string = "Uli"
	var secondName string = "Doni"

	if firstName == "Uli" {
		log.Println(secondName)
	} else {
		log.Println(firstName)
	}
}
