package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestReceiveFile(t *testing.T) {
	// Given
	port := findPort(t)
	testFileContent := []byte("This is a test file.\n" +
		"This is the second line of the test file.\n")
	contentLength := len(testFileContent)
	fmt.Println("Content Length: ", contentLength)
	testFilePath := "testFile.txt"
	listener := mockStartServer(t, port)
	client := mockStartClient(t, port)

	// When
	client.Write(testFileContent)
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}
	err = receiveFile(conn, testFilePath, map[string]string{"HeaderLength": "0", "ContentLength": strconv.Itoa(contentLength)})
	fileContent, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if !bytes.Equal(testFileContent, fileContent) {
		t.Error("The file sent and the file received are not equal.")
	}

	// Clean up
	conn.Close()
	err = os.Remove(testFilePath)
}
