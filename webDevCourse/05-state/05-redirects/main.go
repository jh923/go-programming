package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8181", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at index: ", req.Method, "\n\n")
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

// Example of error 301 and shows the method used at each URL
