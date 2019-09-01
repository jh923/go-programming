package main

import (
	"errors"
	"log"
)

func main() {
	err := errors.New("big spooky new error")
	log.Fatalln(err)

}
