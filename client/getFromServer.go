package main

import (
	"encoding/base64"
	"fmt"
	"net"
)

func getFromServer(conn net.Conn, fileToDownload string) {
	base64Username := base64.StdEncoding.EncodeToString([]byte(username))

	requestString := fmt.Sprintf(`GET SFTP 1.0
	RequestType: GET
	ContentType: text/plain
	ContentLength: 21
	Path: %s/test2.txt
	GUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4
	FileName: test2.txt
	FileSystem: Unix
	FileExtension: .txt
	Authorization: %s`, username, base64Username)

	request := createHeaders(requestString)
	if _, err := conn.Write([]byte(request)); err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
}
