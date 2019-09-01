package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logf, err := os.Create("error_log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logf.Close()
	log.SetOutput(logf)

	f, err := os.Open("Not_a_real.file")
	if err != nil {
		log.Println(err)
		defer fmt.Println("Check err log file. ")
	}
	defer f.Close()
}
