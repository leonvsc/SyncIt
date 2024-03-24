package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		message := strings.TrimSpace(netData)
		testString := getHeaderBytes(message, 4)
		message = parseMessage(message)
		if message == "STOP" {
			break
		}
		_, err = conn.Write([]byte(string(message) + "Bytes = " + string(testString) + "\n"))

	}
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp4", ":8080")
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

func getHeaderBytes(message string, number int) []byte {
	header := []byte(message)
	return header[:number]
}
