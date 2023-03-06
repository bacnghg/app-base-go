package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplates render a template
func RenderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing tamplate: ", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplatesTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we ready have the template in our cache
	_, inMap := tc[t]

	if !inMap {
		// need to create the template
		err = createTemplateCache(t)
		if err != nil {

		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl

	return nil
}
