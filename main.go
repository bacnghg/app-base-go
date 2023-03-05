package main

import (
	"fmt"
	"net/http"
)

var portZNumber = ":8088"

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("NUmber of bytes written: %d", n))
	// })
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Starting application on port", portZNumber)
	_ = http.ListenAndServe(portZNumber, nil)
}
