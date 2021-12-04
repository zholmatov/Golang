package main

import "log"

type User struct {
	Surname string
	Name    string
	age     int
}

func main() {
	user := User{
		Surname: "Zholmatov",
		Name:    "Daniiar",
	}

	log.Println(user.age)
}

// if function or var starts with a capital letter
// it is available outside of the package
func Whatever() {}
