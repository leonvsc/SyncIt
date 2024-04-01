package main

import (
	"bufio"
	"fmt"
	"net"
)

var HeaderLength = 0

func handleConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)
		headerMap := make(map[string]string)
		message := processMessage(scanner, headerMap)

		if headerMap["RequestType"] == "GET" {
			returnMessage := sendFile()
			_, _ = conn.Write(returnMessage[:])
		}
		if headerMap["RequestType"] == "POST" {
			addFileToServer(message, headerMap)
			okMessage := sendOkToClient()
			_, _ = conn.Write(okMessage)
		}

	}
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":50000")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)
	}
}
