package main

import (
	"fmt"
	"net/http"
)

// My type
type Duelist struct {
	Name string
	AceMonster string
}

// Give my type the function which handlers implement
// ServerHTTP w http.ResponseWriter, Request *http.Request
func (d Duelist) ServeHTTP(w http.ResponseWriter, r *http.Request)	{
	fmt.Fprintf(w, "%s's ace monster is %s.", d.Name, d.AceMonster)
}

// Make's a duelist and passes it to our server
func main()	{
	yugi := Duelist{"Yugi", "The Dark Magician"}
	http.ListenAndServe(":8181", yugi)	// use of 8181 over 8080 to keep other web server running
}
