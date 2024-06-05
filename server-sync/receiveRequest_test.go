package main

import (
	"bytes"
	"encoding/base64"
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

func TestProcessAuthRequest(t *testing.T) {
	// Given
	port := findPort(t)
	username := "testUser"
	base64Username := base64.StdEncoding.EncodeToString([]byte(username))
	requestString := fmt.Sprintf("AUTH SFTP 1.0 \n" +
		"Authorization: " + base64Username)
	length := len(requestString)
	headerLength := createHeaderLength(length)
	request := headerLength + requestString
	listener := mockStartServer(t, port)
	client := mockStartClient(t, port)
	// When
	client.Write([]byte(request))
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}
	headerMap := readHeader(conn)
	handleRequest(headerMap, conn)
	fmt.Println("Client User Name: ", clientUserName)
	// Then
	if !bytes.Equal([]byte(clientUserName), []byte(username)) {
		t.Error("The auth username send is not the send username")
	}

	buffer := make([]byte, 1024*1024)
	bytesRead, _ := client.Read(buffer)
	message := buffer[:bytesRead]

	if !bytes.Equal(message, []byte("0029\nAUTH SFTP 1.0\nStatuscode: 200")) {
		t.Error("The request send from the server is correct.")
	}

	// Clean up
	conn.Close()
}
