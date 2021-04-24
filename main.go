package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	// List on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	defer listener.Close()
	fmt.Println("Started server on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
