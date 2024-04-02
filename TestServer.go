package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// Listen for incoming connections
	ln, err := net.Listen("tcp", ":50000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on port 50000")

	// Accept connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		// Handle connections concurrently
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read incoming message
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err.Error())
		return
	}

	// Trim spaces and newlines from the message
	message = strings.TrimSpace(message)

	// Prepare response
	response := `0217SFTP 1.0 200 OK
Content-Type: text/plain; charset=utf-8
ContentLength: 22
Path: ../local/test/
GUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4
FileName: test.txt
FileSize: Unix
FileExtension: txt
Authorization: null

This is line 1 of the body.
This is line 2 of the body.
And this is line 3.
`

	// Write response to the connection
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err.Error())
		return
	}

	fmt.Println("Response sent successfully")
}
