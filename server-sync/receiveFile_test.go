package main

import (
	"bytes"
	"os"
	"testing"
)

func TestReceiveFile(t *testing.T) {
	// Given
	port := findPort(t)
	testFileContent := []byte("This is a test file.")
	testFilePath := "testFile.txt"
	listener := mockStartServer(t, port)
	client := mockStartClient(t, port)

	// When
	client.Write(testFileContent)
	client.Close()
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}
	err = receiveFile(conn, testFilePath)
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
