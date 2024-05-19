package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"os"
	"path/filepath"
)

func sendFile(localFilePath string, conn net.Conn) error {
	headers := generateHeaders(localFilePath)
	return generateBody(localFilePath, headers, conn)
}

func generateHeaders(localFilePath string) []byte {
	contentLength := getContentLength(localFilePath)
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

	return []byte(requestString)

}

func generateBody(localFilePath string, headers []byte, conn net.Conn) error {
	_, err := conn.Write(append(headers, '\n'))
	if err != nil {
		return fmt.Errorf("error sending headers: %w", err)
	}

	// Fill with content length
	length := int64(100)

	// Var to count amount of bytes send
	bytesSent := int64(0)
	bytesRead := int64(0)

	file, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	chunkSize := 1024 * 1024 * 100

	buffer := make([]byte, chunkSize)
	for bytesSent < length {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading file: %w", err)
		}
		if n == 0 {
			break
		}
		_, err = conn.Write(buffer[:n])
		if err != nil {
			return fmt.Errorf("error sending file chunk: %w", err)
		}
		bytesSent += int64(bytesRead)
	}

	fmt.Println("File sent successfully")
	return nil
}

func getContentLength(localFilePath string) int64 {
	fileInfo, err := os.Stat(localFilePath)
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()
	return fileSize
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
