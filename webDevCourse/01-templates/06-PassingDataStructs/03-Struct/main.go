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

	type Duelist struct {
		Name       string
		AceMonster string
	}

	yugi := Duelist{
		Name:       "Yugi",
		AceMonster: "Dark Magician",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", yugi)
	if err != nil {
		log.Fatal(err)
	}
}
