package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to the index page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to the page about dogs")
}

func me(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", "Josh")
	if err != nil {
		log.Fatal(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", "Josh")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
}
