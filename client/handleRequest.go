package main

import (
	"fmt"
	"net"
)

func handleRequest(headerMap map[string]string, conn net.Conn) {
	filePath := "../local" + "/" + headerMap["FileName"]
	switch headerMap["RequestType"] {
	case "GET":
		err := sendFile(filePath, conn)
		if err != nil {
			fmt.Println("Failed to send file:", err)
			return
		}
	case "POST":
		err := receiveFile(conn, filePath, headerMap)
		if err != nil {
			fmt.Println("Failed to receive file:", err)
			return
		}
	default:
	}
}
