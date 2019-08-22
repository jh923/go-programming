package main

import (
	"html/template"
	"log"
	"math"
	"os"
	"strconv"
)

var tpl *template.Template

// Create a map containing a name reference and a function pair in a FuncMap
var fm = template.FuncMap{
	"Sqrt": math.Sqrt,
	"Ceil": math.Ceil,
}

// Rather than just making a container with the files we also pass a template with a blank name that
// Contains out functions by passing the function map to it using the Funcs function/
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	//Take in data
	data := os.Args[1]

	// Convert from string to f64
	i, err := strconv.ParseFloat(data, 64)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "example.gohtml", i)
	if err != nil {
		log.Fatal(err)
	}
}
