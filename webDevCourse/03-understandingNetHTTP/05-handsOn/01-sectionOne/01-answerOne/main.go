package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request)	{
	io.WriteString(res, "Welcome to the index page")
}

func dog(res http.ResponseWriter, req *http.Request)	{
	io.WriteString(res, "Welcome to the page about dogs")
}

func me(res http.ResponseWriter, req *http.Request)	{
	io.WriteString(res, "Welcome Josh")
}

func main()	{
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8181", nil)
}
