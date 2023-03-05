package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
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
	http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(portZNumber, nil)
}
