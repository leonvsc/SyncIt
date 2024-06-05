package main

import (
	"net"
)

//type ClientRequest struct {
//	Username string
//	Response chan *Client
//}
//
//type Client struct {
//	Username      string
//	ClientAddress string
//}

var (
	clientUserName string
	//clientChannel  = make(chan ClientRequest)
	//clients        = make(map[string]*Client)
	//mu             sync.Mutex
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read the header from the connection
	headerMap := readHeader(conn)
	// Handle the request
	handleRequest(headerMap, conn)
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
