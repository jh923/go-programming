package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

// Create a map containing a name reference and a function pair in a FuncMap
var fm = template.FuncMap{
	"Caps": strings.ToUpper,
	"FT":   firstThree,
}

// Rather than just making a container with the files we also pass a template with a blank name that
// Contains out functions by passing the function map to it using the Funcs function/
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
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
