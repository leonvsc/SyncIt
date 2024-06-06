package main

import (
	"fmt"
	"mime"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func handleRequest(headerMap map[string]string, conn net.Conn) {
	filePath := clientUserName + "/" + headerMap["FileName"]
	folderExists, err := exists(clientUserName + "/")
	if err != nil {
		return
	}
	if !folderExists {
		err := os.Mkdir(clientUserName, 0755)
		if err != nil {
			panic(err)
		}
	}
	switch headerMap["RequestType"] {
	case "GET":
		fileExists, err := exists(filePath)
		if err != nil {
			return
		}
		if !fileExists {
			conn.Write([]byte("0021 POST Statuscode: 404"))
			return
		}
		header := createHeaderMap(headerMap, filePath)
		err = sendFile(conn, filePath, header)
		if err != nil {
			fmt.Println("Failed to send file:", err)
			return
		}
	case "POST":
		err := receiveFile(conn, filePath, headerMap)
		if err != nil {
			fmt.Println("Failed to receive file:", err)
			return
		}
	case "AUTH":
		clientUserName = processAuthRequest(headerMap["Authorization"], conn)
	default:
	}
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createHeaderMap(header map[string]string, path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileStats, _ := file.Stat()
	fileLength := fileStats.Size()

	headerMap := make(map[string]string)
	headerMap["RequestType"] = "POST"
	headerMap["ContentLength"] = strconv.Itoa(int(fileLength))
	headerMap["GUID"] = header["GUID"]
	headerMap["Path"] = clientUserName + "/"
	headerMap["FileName"] = header["FileName"]
	headerMap["FileExtension"] = filepath.Ext(header["FileName"])
	headerMap["Authorization"] = header["Authorization"]
	headerMap["MimeType"] = mime.TypeByExtension(headerMap["FileExtension"])
	return headerMap
}
