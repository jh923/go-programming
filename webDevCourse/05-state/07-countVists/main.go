package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8181", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	// If the user has no cookie give them one with a value of 0
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	// Get the current value of the cookie, then increment it
	count, _ := strconv.Atoi(c.Value)
	count++
	c.Value = strconv.Itoa(count)

	// Give the user the cookie with the updated value
	http.SetCookie(w, c)
	io.WriteString(w, c.Value)
}
