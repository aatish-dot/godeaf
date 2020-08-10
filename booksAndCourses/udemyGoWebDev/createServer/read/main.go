package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	for {

		connection, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(connection)

	}

}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("I magically got here\n")

}
