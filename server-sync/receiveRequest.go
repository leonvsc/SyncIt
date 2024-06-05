package main

import (
	"encoding/base64"
	"io"
	"net"
	"os"
)

func receiveFile(conn net.Conn, filePath string, headerMap map[string]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	length := int64(totalLength(headerMap))
	bytesSent := int64(0)
	buffer := make([]byte, 1024*1024) // 1 MB buffer size
	for bytesSent < length {
		bytesRead, err := conn.Read(buffer)
		println("bytesRead: ", bytesRead)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		_, err = file.Write(buffer[:bytesRead])
		if err != nil {
			return err
		}
		bytesSent += int64(bytesRead)
	}

	return nil
}

func processAuthRequest(auth string, conn net.Conn) string {
	userName, _ := base64.StdEncoding.DecodeString(auth)
	clientUserName = string(userName)
	conn.Write([]byte("0029\nAUTH SFTP 1.0\nStatuscode: 200"))
	//remoteAddr := conn.RemoteAddr().String()
	//responseChan := make(chan *Client)
	//clientChannel <- ClientRequest{Username: auth, Response: responseChan}
	//go func() {
	//	for req := range clientChannel {
	//		mu.Lock()
	//		client, exists := clients[req.Username]
	//		if !exists {
	//			client = &Client{Username: req.Username, ClientAddress: remoteAddr}
	//			clients[req.Username] = client
	//		}
	//		mu.Unlock()
	//		req.Response <- client
	//	}
	//}()
	return clientUserName
}
