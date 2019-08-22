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
	kaiba := Duelist{
		Name:       "Kaiba",
		AceMonster: "Blue eyes white dragon",
	}
	joe := Duelist{
		Name:       "Joe",
		AceMonster: "Red eyes black dragon",
	}

	ds := []Duelist{yugi, kaiba, joe}

	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", ds)
	if err != nil {
		log.Fatal(err)
	}
}
