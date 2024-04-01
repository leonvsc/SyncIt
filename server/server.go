package main

import (
	"bufio"
	"fmt"
	"net"
)

const HeaderByte = 5

func handleConnection(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		message, err := bufio.NewReader(conn).ReadBytes('\n')

		println(message)
		if err != nil {
			fmt.Println(err)
			conn.Write([]byte("Error reading data: 500\n"))
			return
		}

		if len(message) < HeaderByte {
			conn.Write([]byte("Error: Message too short: 400\n"))
			return
		}

		parseMessage(message)
		returnMessage := responseMessage()
		fmt.Println(returnMessage[:])
		_, err = conn.Write(returnMessage[:])

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
