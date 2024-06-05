package main

import "net"

func handleRequest(headerMap map[string]string, conn net.Conn) {
	filePath := headerMap["FilePath"]
	switch headerMap["RequestType"] {
	case "GET":
		err := sendFile(filePath, conn)
		if err != nil {
			panic(err)
		}
	case "POST":
		err := receiveFile(conn, filePath, headerMap)
		if err != nil {
			panic(err)
		}
	default:
	}
}
