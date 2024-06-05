package main

import (
	"encoding/binary"
	"fmt"
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
	headerLength := sendHeader(conn, headerMap)
	headerMap["HeaderLength"] = strconv.Itoa(headerLength)
	length := int64(totalLength(headerMap))
	bytesSent := int64(0)
	buffer := make([]byte, 1024*1024) // 1 MB buffer size
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

func sendHeader(conn net.Conn, headerMap map[string]string) int {
	var header string
	for key, value := range headerMap {
		header += key + ": " + value + "\n"
	}
	length := len(header)

	headerLength := createHeaderLength(length)

	header = headerLength + header
	_, err := conn.Write([]byte(header))
	if err != nil {
		panic(err)
	}
	return length
}

func createHeaderLength(length int) string {
	// Convert the length integer to a byte array of 4 bytes
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	// Get the last value of lengthBytes
	lastByte := lengthBytes[3] // 3 is the index from the last element in the slice

	// Convert the last byte to a string
	lastValue := strconv.Itoa(int(lastByte))

	// Add zero's to the left to create a digit with 4 digits
	return fmt.Sprintf("%04s", lastValue)
}

func totalLength(headerMap map[string]string) int {
	headerLength, _ := strconv.Atoi(headerMap["HeaderLength"])
	contentLength, _ := strconv.Atoi(headerMap["ContentLength"])
	return headerLength + contentLength
}
