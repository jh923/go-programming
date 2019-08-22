package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]
	fmt.Println(os.Args[0]) // Prints directory
	fmt.Println(os.Args[1]) // Print first manual enter arg

	str := `
	<!DOCTYPE html>
	<html>
	<title>Page Layout</title>
	<body>

	<!-- Header Section -->
	<header>
	</header>

	<!-- Body section -->
	<div class = "body_sec">
	<section id="Content">
	<h1> 
	Hello	` + name + `
	</h1>
	</section>
	</div>

	<!-- Footer Section -->
	</body>
	</html>
	`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error making file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
