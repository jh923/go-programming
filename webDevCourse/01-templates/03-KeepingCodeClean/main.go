package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	// Pull in data from a given file
	tpl, err := template.ParseFiles("example.gohtml") //RUn from terminal otherwise file not found
	if err != nil {
		log.Fatal(err)
	}

	// Note ParseGlob could also be used here to pass all files which match a regex
	// And example being:
	// tpl.ParseGlob("dir/*.gohtml")
	// this parses all files with the gohrml extension in the dir.

	// Create a file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	// Parse data into the new file. Can also parse to os.stdout
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Once a template is made we can parse more files into it
	tpl.ParseFiles("exampleP2.gohtml")

	// When more files are stored in the templates we can execute them by name
	err = tpl.ExecuteTemplate(os.Stdout, "example.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n----------------------------------------------------------")

	err = tpl.ExecuteTemplate(os.Stdout, "exampleP2.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}
