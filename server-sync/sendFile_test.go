package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
)

func TestSendFile(t *testing.T) {
	// Given
	port := findPort(t)
	testFileContent := []byte("This is a test file.")
	testFilePath := "testFile.txt"
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
	err = sendFile(conn, testFilePath)
	conn.Close()
	fileContent, err := mockClientRead(client)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	fmt.Println("FileContent: ", string(fileContent))
	if !bytes.Equal(testFileContent, fileContent) {
		t.Error("The file sent and the file received are not equal.")
	}
	err = os.Remove(testFilePath)
}

func mockClientRead(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 1024*1024) // 1 MB buffer size
	message := make([]byte, 0)
	for {
		bytesRead, err := conn.Read(buffer)
		println("bytesRead: ", bytesRead)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		message = append(message, buffer[:bytesRead]...)
	}
	conn.Close()
	return message, nil
}
