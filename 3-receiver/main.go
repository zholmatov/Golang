package main

import "log"

type myStruct struct {
	Surname string
	Name    string
	Age     int
}

// this m points to myStruct
func (m *myStruct) printName() string {
	return m.Surname
}

func main() {
	var user1 myStruct
	user1.Surname = "Zhol"

	user2 := myStruct{
		Surname: "Mat",
	}
	// that's how it prints surname
	log.Println(user1.printName(), user2.printName())
}
