package main

import (
	"fmt"
	"os"
)

func createBody(localPath string) []byte {
	// Read the textfile in a byte array
	fileContent, err := os.ReadFile(localPath)
	if err != nil {
		fmt.Println("Error when reading the textfile:", err)
	}

	return fileContent
}

func createBody2(localPath string) ([]byte, error) {
	file, err := os.Open(localPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	var fileBytes []byte

	const maxChunkSize = 10 // 1MB chunk size (adjust as needed)
	for {
		chunk := make([]byte, maxChunkSize)
		n, err := file.Read(chunk)
		if err != nil && err.Error() != "EOF" {
			return nil, err
		}
		if n == 0 {
			break
		}
		fileBytes = append(fileBytes, chunk[:n]...)
	}
	return fileBytes, nil
}
