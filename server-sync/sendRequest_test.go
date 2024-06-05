package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"testing"
)

func TestSendFile(t *testing.T) {
	// Given
	port := findPort(t)
	testFileContent := []byte("This is a test file.")
	response := "0018ContentLength: 20\nThis is a test file."
	contentLength := len(testFileContent)
	testFilePath := "testFile.txt"
	headerMap := map[string]string{"ContentLength": strconv.Itoa(contentLength)}
	file, err := os.Create(testFilePath)
	_, err = file.Write(testFileContent)
	file.Close()
	listener := mockStartServer(t, port)
	client := mockStartClient(t, port)

	// When
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}
	err = sendFile(conn, testFilePath, headerMap)
	fileContent, err := mockClientRead(client, int64(contentLength))
	if err != nil {
		t.Fatal(err)
	}

	// Then
	fmt.Println("FileContent: ", string(fileContent))
	if !bytes.Equal([]byte(response), fileContent) {
		t.Error("The file sent and the file received are not equal.")
	}
	err = os.Remove(testFilePath)
}

func TestSendFileTotalLength(t *testing.T) {
	// Given
	headerMap := map[string]string{"HeaderLength": "20", "ContentLength": "21", "FilePath": "testFile.txt"}
	headerLength, _ := strconv.Atoi(headerMap["HeaderLength"])
	contentLength, _ := strconv.Atoi(headerMap["ContentLength"])
	lengthGiven := headerLength + contentLength

	// When
	length := totalLength(headerMap)

	// Then
	if !bytes.Equal([]byte(strconv.Itoa(length)), []byte(strconv.Itoa(lengthGiven))) {
		t.Error("The total length of the file sent and the file received are not equal.")
	}

}

func mockClientRead(conn net.Conn, length int64) ([]byte, error) {
	buffer := make([]byte, 1024*1024) // 1 MB buffer size
	message := make([]byte, 0)
	bytesSent := int64(0)
	for bytesSent < length {
		bytesRead, err := conn.Read(buffer)
		println("bytesRead: ", bytesRead)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		message = append(message, buffer[:bytesRead]...)
		bytesSent += int64(bytesRead)
	}
	conn.Close()
	fmt.Println("Message: ", string(message))
	return message, nil
}
