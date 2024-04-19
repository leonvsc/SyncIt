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
	response := `0029
AUTH SFTP 1.0
Statuscode: 200
`

	// Write response to the connection
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err.Error())
		return
	}

	fmt.Println("Response sent successfully")
}
