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

		//message := strings.TrimSpace(netData)
		if len(message) < HeaderByte {
			conn.Write([]byte("Error: Message too short: 400\n"))
			continue
		}

		parseMessage(message)
		//headerLength := netData[:HeaderByte]
		//print(headerLength)
		//message = parseMessage(message, []byte(netData[:HeaderByte]))
		//
		//if message == "STOP" {
		//	break
		//}

		_, err = conn.Write([]byte(string(message) + "\n"))

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
