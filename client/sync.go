package main

import (
	"fmt"
	"path/filepath"
)

func sync() {
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
	fmt.Println(finalResponse)
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
