package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // Returns an error as no fav icon is set
	http.ListenAndServe(":8181", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Do my search "+v)
}

// Use http://localhost:8181/?q=String
