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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>
	`+v)
}

// The method "POST" passed data though the body, the method "GET" passes data though the url
// The key matches the name given to input's name
