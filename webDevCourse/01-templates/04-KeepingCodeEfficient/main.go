package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

// Must takes the template and an error and just returns a template
// This use of parseglob is used in init so it is only run once
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// Once a template is made we can parse more files into it
	tpl.ParseFiles("exampleP2.gohtml")

	// When more files are stored in the templates we can execute them by name
	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n----------------------------------------------------------")

	err = tpl.ExecuteTemplate(os.Stdout, "exampleP2.gohtml", nil)
	if err != nil {

	}
}
