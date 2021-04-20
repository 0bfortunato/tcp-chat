package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Unable to start server: %s", err.Error())
	}
}
