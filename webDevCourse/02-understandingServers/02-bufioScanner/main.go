package main

import (
	"bufio"
	"fmt"
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}
