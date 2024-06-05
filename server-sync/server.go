package main

import (
	"net"
)

var (
	clientUserName string
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read the header from the connection
	for {
		headerMap := readHeader(conn)
		// Handle the request
		handleRequest(headerMap, conn)
	}
}

func main() {
	// Start a TCP server
	ln, err := net.Listen("tcp", ":50000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		// Accept connections from clients
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}
