package main

import (
	"fmt"
	"log"
)

type customErr struct {
	info string
}

func (err customErr) Error() string {
	return fmt.Sprintf("here is our error: %v", err.info)
}

func main() {
	error := customErr{
		info: "Our custom error has occurred",
	}

	foo(error)
}

func foo(err error) {
	log.Println(err)
}
