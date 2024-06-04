package main

import (
	"fmt"
	"net"
	"testing"
)

func mockStartClient(t *testing.T, port int) net.Conn {
	portString := fmt.Sprint(port)
	conn, err := net.Dial("tcp", ":"+portString)
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func mockStartServer(t *testing.T, port int) net.Listener {
	portString := fmt.Sprint(port)
	ln, err := net.Listen("tcp", ":"+portString)
	if err != nil {
		t.Error(err)
	}
	return ln
}

func findPort(t *testing.T) int {
	addr, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	port := addr.Addr().(*net.TCPAddr).Port
	err = addr.Close()
	if err != nil {
		return 0
	}
	return port
}
