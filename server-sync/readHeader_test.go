package main

import (
	"bytes"
	"testing"
)

var headerMap = make(map[string]string)

func TestReadHeaderRequestType(t *testing.T) {
	// Given
	header := "0003GET"
	requestType := "GET"
	listener := mockStartServer(t)
	client := mockStartClient(t)

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

func TestReadHeaderLength(t *testing.T) {
	// Given
	header := "0003GET"
	length := "3"
	listener := mockStartServer(t)
	client := mockStartClient(t)

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
