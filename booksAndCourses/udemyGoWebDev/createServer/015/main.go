package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

		io.WriteString(connection, "\n Hellow from tcp server\n")
		fmt.Fprintln(connection, "I hope you are doing good")
		fmt.Fprintf(connection, "%v", "I am doing fantastic\n")

		connection.Close()

	}

}
