package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	tpl.ParseFiles("exampleP2.gohtml")

	// Using the {{.}} method only 1 bit of data can be passed
	// So we pass a data struct
	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", "big fella.")
	if err != nil {
		log.Fatal(err)
	}
}
