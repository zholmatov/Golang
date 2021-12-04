package main

import (
	"fmt"
	"net/http"
)

func main() {

	// all the http has request and response, and here, request is pointer
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world!")

		// if there is an error
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Sprintf() combines int and string or
		// any other types and returns string
		fmt.Printf("Number of bytes written: %d", n)
	})

	// _ = means if there's an error just throw it to the trash
	// and now, if you go to localhost:8080, you can see Hello world!
	_ = http.ListenAndServe(":8080", nil)
}
