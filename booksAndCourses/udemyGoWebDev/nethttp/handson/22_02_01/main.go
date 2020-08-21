// # Create a basic server using TCP.

// The server should use net.Listen to listen on port 8080.

// Remember to close the listener using defer.

// Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.

// Now write a response back on the connection.

// Use io.WriteString to write the response: I see you connected.

// Remember to close the connection.

// Once you have all of that working, run your TCP server and test it from telnet (telnet localhost 8080).

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error in listening on port 8080:", err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln("error in accepting connections: ", err)
		}

		go handle(conn)
	}

}

func handle(netcon net.Conn) {
	defer netcon.Close()

	request(netcon)
	response(netcon)
}

func request(netcon net.Conn) {
	i := 0

	scanner := bufio.NewScanner(netcon)
	var url string
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
		switch {
		case i == 0:
			fields := strings.Fields(line)
			fmt.Println("Method****:", fields[0])
			url = fields[1]

		case i == 1:
			url = strings.Fields(line)[1] + url
			fmt.Println("URL is: ", url)
		case line == "":
			break
		}
		i++
	}

}

func response(c net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>This is handson exercise 22_02_01</title></head><body><strong>I see you connected</strong></body></html>`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}
