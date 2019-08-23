package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Port 8181 used over 8080 as I already have a web server running on 8080
	li, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nWelcome to my first GO TCP server\n")
		fmt.Fprintln(conn, "How's your day going?")
		fmt.Fprintf(conn, "%v", "wrll, I hope!")
	}
}
