package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 3)
	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is: %d", sum)
}

func AddValues(x, y int) int {
	var sum int = x + y

	return sum
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting an application on port: %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
