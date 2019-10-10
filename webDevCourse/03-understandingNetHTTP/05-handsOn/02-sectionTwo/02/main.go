package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}

		io.WriteString(conn, "Connected")
		conn.Close()
	}
}
