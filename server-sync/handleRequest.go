package main

import "net"

func handleRequest(headerMap map[string]string, conn net.Conn) {
	filePath := headerMap["FilePath"]
	switch headerMap["RequestType"] {
	case "GET":
		err := sendFile(conn, filePath, nil)
		if err != nil {
			panic(err)
		}
	case "POST":
		err := receiveFile(conn, filePath, headerMap)
		if err != nil {
			panic(err)
		}
	case "AUTH":
		err := processAuthRequest(headerMap["Authorization"])
		if err != nil {
			panic(err)
		}
	default:
	}
}
