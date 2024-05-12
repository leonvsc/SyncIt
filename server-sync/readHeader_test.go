package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadHeaderRequestType(t *testing.T) {
	// Loop through all the request types to ensure that the header is read correctly
	// Given
	headerMap := make(map[string]string)
	requestTypes := []string{"GET", "POST", "PUT", "DELETE", ""}
	for _, requestType := range requestTypes {
		port := findPort(t)
		listener := mockStartServer(t, port)
		client := mockStartClient(t, port)
		header := "0003" + requestType

		// When
		client.Write([]byte(header))
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		headerMap = readHeader(conn)

		// Then
		if !bytes.Equal([]byte(requestType), []byte(headerMap["RequestType"])) {
			t.Error("The header sent and the header received are not equal.")
		}

		// Clean up
		client.Close()
		conn.Close()
	}
}

func TestReadHeaderLength(t *testing.T) {
	// Given
	headerMap := make(map[string]string)
	port := findPort(t)
	header := "0003GET"
	length := "3"
	listener := mockStartServer(t, port)
	client := mockStartClient(t, port)

	// When
	client.Write([]byte(header))
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}
	headerMap = readHeader(conn)

	// Then
	if !bytes.Equal([]byte(length), []byte(headerMap["HeaderLength"])) {
		t.Error("The header sent and the header received are not equal.")
	}

	// Clean up
	client.Close()
	conn.Close()
}

func TestReadHeaderLines(t *testing.T) {
	// Given
	headerMap := make(map[string]string)
	port := findPort(t)
	header := "0003GET \n" +
		"RequestType: GET\n" +
		"ContentLength: 100\n" +
		"GUID: 1234\n" +
		"Path: /test\n" +
		"FileSystem: /tmp\n" +
		"FileName: test.txt\n" +
		"FileExtension: txt\n" +
		"Authorisation: Basic\n" +
		"MimeType: text/plain\n"

	content := make(map[string]string)
	content["RequestType"] = "GET"
	content["ContentLength"] = "100"
	content["GUID"] = "1234"
	content["Path"] = "/test"
	content["FileSystem"] = "/tmp"
	content["FileName"] = "test.txt"
	content["FileExtension"] = "txt"
	content["Authorisation"] = "Basic"
	content["MimeType"] = "text/plain"
	listener := mockStartServer(t, port)
	client := mockStartClient(t, port)

	// When
	client.Write([]byte(header))
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}
	headerMap = readHeader(conn)

	// Then
	for key, value := range headerMap {
		fmt.Println("Value: ", value)
		fmt.Println("Content key: ", content[key])
		if key == "HeaderLength" {
			continue
		}
		if !bytes.Equal([]byte(content[key]), []byte(value)) {
			t.Error("The header sent and the header received are not equal.")
		}
	}

	// Clean up
	client.Close()
	conn.Close()
}
