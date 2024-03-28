package main

import (
	"fmt"
	"os"
)

func bodyResponse(localPath string) []byte {
	// Lees de inhoud van het tekstbestand in een byte-array
	fileContent, err := os.ReadFile(localPath)
	if err != nil {
		fmt.Println("Fout bij lezen van het tekstbestand:", err)
	}

	return fileContent
}
