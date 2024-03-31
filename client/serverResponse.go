package main

import (
	"bufio"
	"fmt"
	"net"
)

func serverResponse(conn net.Conn) {
	// Create a reader to read from the connection
	reader := bufio.NewReader(conn)

	// Continuously read from the connection
	for {
		// Read a line from the connection
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// Print the received message
		fmt.Print("Message from server:", message)
	}
}
