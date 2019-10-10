package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func index(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
