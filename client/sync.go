package main

import "net"

var folderPath = "../local"

func sync(conn net.Conn) {
	//pushFolderToServer(conn, folderPath)
	pullFromServer(conn)
	//runMainMenu()
}
