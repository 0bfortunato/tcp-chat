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
		log.Fatalf("Unable to start server: %s\n\n", err.Error())
	}

	defer listener.Close()
	fmt.Printf("Started server on port: %s\n\n", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s\n\n", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
