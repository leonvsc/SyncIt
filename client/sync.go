package main

import "net"

func sync(conn net.Conn) {
	pushToServer(conn)
	pullFromServer(conn)
	runMainMenu()
}
