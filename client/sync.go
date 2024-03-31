package main

import (
	"fmt"
	"net"
	"path/filepath"
)

func sync() {

	// Connect to TCP server
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	localFilePath := "../local/input.txt"
	bodyResponseResult := bodyResponse(localFilePath)

	contentLength := getContentLength(bodyResponseResult)
	fileName := getFileName(localFilePath)
	fileExtension := getFileExtension(localFilePath)

	headerResponseResult := headerResponse("text/plain", contentLength, localFilePath, fileName, fileExtension)

	// Convert byte array to string
	bodyResponseString := string(bodyResponseResult)

	// Concatenate header response, newline character, and body response string
	finalResponse := headerResponseResult + "\n" + bodyResponseString

	// Print or do something with finalResponse
	//fmt.Println(finalResponse)

	// Send response to the server
	_, err = conn.Write([]byte(finalResponse))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("Data sent successfully.")

	runMainMenu()
}

func getContentLength(bodyResponse []byte) int {
	length := len(bodyResponse)
	return length
}

func getFileName(localFilePath string) string {
	fileName := filepath.Base(localFilePath)
	return fileName
}

func getFileExtension(localFilePath string) string {
	extension := filepath.Ext(localFilePath)
	// Remove the dot from the extension
	extension = extension[1:]
	return extension
}
