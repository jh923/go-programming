package main

import (
	"fmt"
	"net/http"
)

type data int

func (m data) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Key", "Value")
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(rw, "<h1>Anything</h1>") // Prints formatting is set to text/html
	// Would print <h1>...</h1> if using text/plain
}

func main() {
	var d data
	http.ListenAndServe(":8181", d)
}
