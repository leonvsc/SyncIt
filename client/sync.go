package main

import "net"

var folderPath = "../local"
var filePath = "../local/input.txt"

func sync(conn net.Conn) {
	//pushFolderToServer(conn, folderPath)
	//pushFileToServer(conn, filePath)
	pullFromServer(conn)
	//runMainMenu()
}
