package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	serv, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer serv.Close()

	for {
		connection, err := serv.Accept()
		if err != nil {
			log.Println("Accepting connections failed", err)
		}
		go handle(connection)

	}

}

func handle(netcon net.Conn) {
	defer netcon.Close()

	request(netcon)

}

func request(reqcon net.Conn) {
	i := 0
	scanner := bufio.NewScanner(reqcon)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		switch {
		case i == 0:
			fields := strings.Fields(line)
			m := fields[0]
			fmt.Println("****METHOD:", m)
			respondmux(reqcon, line)
		case line == "":
			break

		}

		i++

	}
}

func respondmux(respcon net.Conn, url string) {
	var body string
	switch {
	case strings.Contains(url, "blah"):
		body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World ! blah blah blah</strong></body></html>`
	case strings.Contains(url, "foo/bar"):
		body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World !Its all FOOBAR</strong></body></html>`
	default:
		body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World !Life is a circus</strong></body></html>`

	}

	fmt.Fprint(respcon, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(respcon, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(respcon, "Content-Type: text/html\r\n")
	fmt.Fprint(respcon, "\r\n")
	fmt.Fprint(respcon, body)
}
