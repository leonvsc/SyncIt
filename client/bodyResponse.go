package main

import (
	"fmt"
	"os"
)

func bodyResponse(localPath string) []byte {
	// Read the textfile in a byte array
	fileContent, err := os.ReadFile(localPath)
	if err != nil {
		fmt.Println("Error when reading the textfile:", err)
	}

	return fileContent
}
