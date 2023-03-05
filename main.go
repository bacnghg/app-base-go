package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var portZNumber = ":8088"

func Home(w http.ResponseWriter, r *http.Request) {

	renderTemplates(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")

}

func renderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing tamplate: ", err)
		return
	}
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
	fmt.Println("Starting application on port", portZNumber)
	_ = http.ListenAndServe(portZNumber, nil)
}
