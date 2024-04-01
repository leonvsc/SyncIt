package main

import (
	"fmt"
	"mime"
	"net"
	"os"
	"path/filepath"
)

func pushFolderToServer(conn net.Conn, folderPath string) {
	// Get a list of files and directories in the folder
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		return
	}

	for i, entry := range entries {
		// Construct the full path of the entry
		entryPath := filepath.Join(folderPath, entry.Name())

		if entry.IsDir() {
			// If the entry is a directory, recursively call pushFolderToServer
			pushFolderToServer(conn, entryPath)
		} else {
			// If the entry is a file, push it to the server
			pushFileToServer(conn, entryPath)

			// Add a delimiter between files (except for the last file)
			if i < len(entries)-1 {
				conn.Write([]byte("\n\n"))
			}
		}
	}

	// Add a delimiter after syncing the folder
	conn.Write([]byte("\n\n"))
	fmt.Println("Folder synced successfully.")
}

func pushFileToServer(conn net.Conn, localFilePath string) {
	fmt.Println("Syncing file:", localFilePath)

	bodyResponseResult := bodyResponse(localFilePath)

	contentLength := getContentLength(bodyResponseResult)
	fileName := getFileName(localFilePath)
	fileExtension := getFileExtension(localFilePath)

	headerResponseResult := headerResponse("image/png", contentLength, localFilePath, fileName, fileExtension)

	// Convert byte array to string
	bodyResponseString := string(bodyResponseResult)

	// Concatenate header response, newline character, and body response string
	finalResponse := []byte(headerResponseResult + "\n" + bodyResponseString)

	// Send response to the server
	_, err := conn.Write(finalResponse)
	if err != nil {
		fmt.Println("Error sending data for file", localFilePath, ":", err)
		return
	}

	fmt.Println("File", localFilePath, "sent successfully.")
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
