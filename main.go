package main

import (
	"fmt"
	"net/http"
)

var portZNumber = ":8088"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepage")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(10, 20)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 10 + 20 is %d", sum))
}

func addValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

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

	_ = http.ListenAndServe(portZNumber, nil)
}
