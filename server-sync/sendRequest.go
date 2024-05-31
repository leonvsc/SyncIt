package main

import (
	"io"
	"net"
	"os"
	"strconv"
)

func sendFile(conn net.Conn, filePath string, headerMap map[string]string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	length := int64(totalLength(headerMap))
	bytesSent := int64(0)
	buffer := make([]byte, 1024*1024) // 1 MB buffer size
	//buffer := make([]byte, 1)
	for bytesSent < length {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		_, err = conn.Write(buffer[:bytesRead])
		if err != nil {
			return err
		}
		bytesSent += int64(bytesRead)
	}

	return nil
}

func totalLength(headerMap map[string]string) int {
	headerLength, _ := strconv.Atoi(headerMap["HeaderLength"])
	contentLength, _ := strconv.Atoi(headerMap["ContentLength"])
	return headerLength + contentLength
}
