package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	FirstName string
	LastName  string
	Hobby     string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8181", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	h := req.FormValue("hobby")

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, h})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
