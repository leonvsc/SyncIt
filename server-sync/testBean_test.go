package main

import (
	"net"
	"testing"
)

func mockStartClient(t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", "localhost:9096")
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func mockStartServer(t *testing.T) net.Listener {
	ln, err := net.Listen("tcp", ":9096")
	if err != nil {
		t.Error(err)
	}
	return ln
}
