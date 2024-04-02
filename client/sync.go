package main

import "net"

func sync(conn net.Conn) {
	//pushFolderToServer(conn, folderPath)
	//pushFileToServer(conn, filePath)
	pullFromServer(conn)
	//runMainMenu()
}
