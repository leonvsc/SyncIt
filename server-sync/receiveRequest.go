package main

import (
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

func processAuthRequest(auth string) error {

	return nil
}
