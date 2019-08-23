package main

import (
	"html/template"
	"log"
	"net/http"
)

// Custom type to implement handler
type form string

// ServeHTTTP executes selected template and parses the form data.
func(f form) ServeHTTP(rw http.ResponseWriter, req *http.Request)	{
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(rw, "index.gohtml", req.Form)
}

// Make the template pointer global
var tpl *template.Template

// init initialise the pointer to contain all .gohtlm files from the templates dir.
func init()	{
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

//  main creates a variable that implements the handler interface and starts listening on port 8181.
func main()	{
	var d form
	http.ListenAndServe(":8181", d)	// use of 8181 over 8080 to keep other web server running
}
