package main

import (
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"path/filepath"
)

func getFileList() string {
	// Get list of files in folder
	file, err := os.Open(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the directory contents
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	// Collect file names in a slice
	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	// Create a prompt for selecting a file
	prompt := promptui.Select{
		Label: "Select a file to sync",
		Items: fileNames,
	}

	// Run the prompt
	_, fileToSync, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	fullPath := filepath.Join(folderPath, fileToSync)

	return fullPath
}
