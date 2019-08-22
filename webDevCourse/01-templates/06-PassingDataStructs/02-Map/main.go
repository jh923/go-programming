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

	xs := map[string]string{
		"Yugi":  "Dark Magician",
		"Kaiba": "Blue eyes white dragon",
		"Joe":   "Red eyes black dragon",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", xs)
	if err != nil {
		log.Fatal(err)
	}
}
