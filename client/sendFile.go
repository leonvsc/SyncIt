package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func sendFile(fileToSync string, conn net.Conn) error {
	headers := generateHeaders(fileToSync)
	return generateBody(fileToSync, headers, conn)
}

func generateHeaders(fileToSync string) []byte {
	contentLength := getContentLength(fileToSync)
	fileName := getFileName(fileToSync)
	fileExtension := getFileExtension(fileToSync)
	contentType := getContentType(fileToSync)
	base64Username := base64.StdEncoding.EncodeToString([]byte(username))

	requestString := fmt.Sprintf(`POST SFTP 1.0
	RequestType: POST
	ContentType: %s
	ContentLength: %d
	Path: %s
	GUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4
	FileName: %s
	FileSystem: Unix
	FileExtension: %s
	Authorization: %s`, contentType, contentLength, fileToSync, fileName, fileExtension, base64Username)

	request := createHeaderLength(requestString)

	return []byte(request)

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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

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

func createHeaderLength(requestString string) string {
	byteArray := []byte(requestString)

	// Calculate the length of the byte array
	length := len(byteArray)

	// Convert the length integer to a byte array of 4 bytes
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	// Get the last value of lengthBytes
	lastByte := lengthBytes[3] // 3 is the index from the last element in the slice

	// Convert the last byte to a string
	lastValue := strconv.Itoa(int(lastByte))

	// Add zero's to the left to create a digit with 4 digits
	paddedLastValue := fmt.Sprintf("%04s", lastValue)

	// Print the result
	request := paddedLastValue + "\n" + requestString
	return request
}
