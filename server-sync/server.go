package main

import (
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read the header from the connection
	headerMap := readHeader(conn)
	// TODO: Get the file path from the header
	filePath := "testfile.txt"
	if headerMap["RequestType"] == "GET" {
		err := sendFile(conn, filePath)
		if err != nil {
			panic(err)
		}
	}
	switch headerMap["RequestType"] {
	case "GET":
		err := sendFile(conn, filePath)
		if err != nil {
			panic(err)
		}
	case "POST":
		err := receiveFile(conn, filePath)
		if err != nil {
			panic(err)
		}
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
