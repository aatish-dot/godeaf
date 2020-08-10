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

	respond(netcon)
}

func request(reqcon net.Conn) {
	i := 0
	scanner := bufio.NewScanner(reqcon)
	var u string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		switch {
		case i == 0:
			fields := strings.Fields(line)
			m := fields[0]
			fmt.Println("****METHOD:", m)
			u = fields[1]
		case i == 1:
			u = strings.Fields(line)[1] + u
			fmt.Println("url is ", u)

		case line == "":
			break

		}

		i++

	}
}

func respond(respcon net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(respcon, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(respcon, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(respcon, "Content-Type: text/html\r\n")
	fmt.Fprint(respcon, "\r\n")
	fmt.Fprint(respcon, body)
}
