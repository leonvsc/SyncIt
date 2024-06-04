package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"testing"
)

func TestSendFile(t *testing.T) {
	// Given
	port := findPort(t)
	testFileContent := []byte("This is a test file.\n" +
		"This is the second line of the test file.\n")
	testFilePath := "testFile.txt"
	err := os.WriteFile(testFilePath, testFileContent, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Error(err)
		}
	}(testFilePath)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		t.Fatal(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			t.Error(err)
		}
	}(listener)

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				t.Error(err)
			}
		}(conn)

		// Read headers
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Error(err)
			return
		}
		headers := buf[:n]
		fmt.Println("Headers received:", string(headers))

		// Read file content
		receivedContent := make([]byte, len(testFileContent))
		_, err = conn.Read(receivedContent)
		if err != nil {
			t.Error(err)
			return
		}

		// Compare file content
		if !bytes.Equal(testFileContent, receivedContent) {
			t.Error("The file sent and the file received are not equal.")
		}
	}()

	client, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		t.Fatal(err)
	}
	defer func(client net.Conn) {
		err := client.Close()
		if err != nil {
			t.Error(err)
		}
	}(client)

	// When
	err = sendFile(testFilePath, client)
	if err != nil {
		t.Fatal(err)
	}
}
