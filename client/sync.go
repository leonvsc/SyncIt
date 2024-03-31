package main

import "net"

func sync(conn net.Conn) {
	pushToServer(conn)
	runMainMenu()
}
