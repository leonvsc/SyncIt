package main

import "net"

func sync(conn net.Conn) {
	pullFromServer(conn)
	runMainMenu()
}
