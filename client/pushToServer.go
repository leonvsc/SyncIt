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

	bodyResult := createBody(localFilePath)

	contentLength := getContentLength(bodyResult)
	fileName := getFileName(localFilePath)
	fileExtension := getFileExtension(localFilePath)
	contentType := getContentType(localFilePath)

	requestString := fmt.Sprintf(`POST SFTP 1.0
	ContentType: %s
	ContentLength: %d
	Path: %s
	GUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4
	FileName: %s
	FileSystem: Unix
	FileExtension: %s
	Authorization: null`, contentType, contentLength, localFilePath, fileName, fileExtension)

	headerResponseResult := createHeaders(requestString)

	// Convert byte array to string
	bodyResponseString := string(bodyResult)

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

func getContentType(filePath string) string {
	// Get the extension of the file
	ext := filepath.Ext(filePath)

	// Lookup the content type based on the extension
	contentType := mime.TypeByExtension(ext)

	// If the content type is not found, default to application/octet-stream
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return contentType
}
