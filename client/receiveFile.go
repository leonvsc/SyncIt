package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func receiveFile(conn net.Conn, filePath string, headerMap map[string]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	contentLength := headerMap["ContentLength"]
	content, _ := strconv.Atoi(contentLength)
	length := int64(content)
	bytesSent := int64(0)
	fmt.Println("Length: ", length)
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

func totalLength(headerMap map[string]string) int {
	headerLength, _ := strconv.Atoi(headerMap["HeaderLength"])
	fmt.Println("Header Length: ", headerLength)
	contentLength, _ := strconv.Atoi(headerMap["ContentLength"])
	fmt.Println("Content Length: ", contentLength)
	return headerLength + contentLength
}
