package main

import (
	"io"
	"net/http"
)

// Rather than making a new type which implements the handler interface we create a function which implements the
// handler interface and just pass that function

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "It's time to duel!")
}

func pm(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Gotta catch 'em all!")
}

func main() {
	// Handle func takes a URI path for what it is meant to handle and a function which implements the handler interface
	http.HandleFunc("/duelist", d)     // Only responds to website/duelist
	http.HandleFunc("/pkmmaster/", pm) // Responds to any uri in the form website/pkmmaster/whatever/.../

	http.ListenAndServe(":8181", nil) // nil passed to use default serve mux
}
