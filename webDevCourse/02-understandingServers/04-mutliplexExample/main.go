package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Initialises the server and then infinitely loops listening for quests to be handled
// Putting each handler on a new go routine
// Closes connection when program stops running
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

// Handles closing the connection the gfo routine ends and passes the connection on to the request function
func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

// Limits each go routine to one request and scan in data so the mutexer can decide how to handle the request
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// Headers are done
			break
		}
		i++ //Request served
	}
}

// Split the request string into two parts, request type and uri. So it can decide which function to call
// and prints out the top of a standard HTTP request
func mux(conn net.Conn, ln string) {
	// Request line
	m := strings.Fields(ln)[0] // Method of connection i.e GET
	u := strings.Fields(ln)[1] // uri

	// Standard http request response
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	//Decided on what action to take
	if m == "GET" && u == "/" {
		index(conn)
	} else if m == "GET" && u == "/about" {
		about(conn)
	} else if m == "GET" && u == "/content" {
		content(conn)
	} else if m == "GET" {
		// If the user tries to get any page the doesn't exist
		notFound(conn)
	}

}

func index(conn net.Conn) {
	// It is best practice to sue a index.gohtml but this code is for example only
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	</head>
	<title>Index</title>
	<body>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/content">content</a><br>
	</body>
	</html>
	`

	// Standard format HTML request response
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n ")
	fmt.Fprint(conn, "\r\n")
	// After response print content
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	// It is best practice to sue a index.gohtml but this code is for example only
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	</head>
	<title>about</title>
	<body>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/content">content</a><br>
	</body>
	</html>
	`

	// Standard format HTML request response
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n ")
	fmt.Fprint(conn, "\r\n")
	// After response print content
	fmt.Fprint(conn, body)
}

func content(conn net.Conn) {
	// It is best practice to sue a index.gohtml but this code is for example only
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	</head>
	<title>content</title>
	<body>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/content">content</a><br>
	</body>
	</html>
	`
	// Standard format HTML request response
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n ")
	fmt.Fprint(conn, "\r\n")
	// After response print content
	fmt.Fprint(conn, body)
}

func notFound(conn net.Conn) {
	fmt.Fprint(conn, "HTTP/1.1 404 Page Not Found\r\n")
	fmt.Fprint(conn, "r\n")
}
