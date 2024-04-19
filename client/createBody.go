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
