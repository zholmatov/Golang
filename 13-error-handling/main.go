package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is about page")
}

func divide(w http.ResponseWriter, r *http.Request) {
	f, err := DivideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 0.0, f)
}

func DivideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", divide)

	fmt.Printf("Starting an application on port: %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
