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

	xs := []string{"The dark magician", "Blue eyes white dragon", "BLS"}
	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", xs)
	if err != nil {
		log.Fatal(err)
	}
}
