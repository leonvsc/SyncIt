package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func pullFromServer(conn net.Conn) {
	request := generateGetHeader()
	if _, err := conn.Write([]byte(request)); err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}

	// Read the response from the server
	response := make([]byte, 1024) // Adjust buffer size according to your needs
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Process the response
	processResponse(string(response[:n]))

	// generate get request

	// receive response

	// generate file from response
}

func generateGetHeader() string {
	// Construct the request string
	requestString := fmt.Sprintf(`GET SFTP 1.0`)

	// Convert the request string to a byte array
	byteArray := []byte(requestString)

	// Calculate the length of the byte array
	length := len(byteArray)

	// Convert the length integer to a byte array of 4 bytes
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	// Get the last value of lengthBytes
	lastByte := lengthBytes[3] // 3 is de index van het laatste element in de slice

	// Converteer de laatste byte naar een string
	lastValue := strconv.Itoa(int(lastByte))

	// Voeg nullen toe aan de linkerkant om een getal van 4 cijfers te krijgen
	paddedLastValue := fmt.Sprintf("%04s", lastValue)

	// Print het resultaat
	result := paddedLastValue + "\n" + requestString
	return result
}

func processResponse(response string) {
	// Split the response into lines
	lines := strings.Split(response, "\\n")

	// Process each line of the response
	var contentType, contentLength, path, guid, fileName, fileSystem, fileExtension, authorization, body string

	for _, line := range lines {
		// Split each line by ": "
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Assign values to corresponding variables
		switch key {
		case "ContentType":
			contentType = value
		case "ContentLength":
			contentLength = value
		case "Path":
			path = value
		case "GUID":
			guid = value
		case "FileName":
			fileName = value
		case "FileSystem":
			fileSystem = value
		case "FileExtension":
			fileExtension = value
		case "Authorization":
			authorization = value
		}
	}

	lastIndex := len(lines) - 1
	body = lines[lastIndex]

	// Print the values
	fmt.Println("ContentType:", contentType)
	fmt.Println("ContentLength:", contentLength)
	fmt.Println("Path:", path)
	fmt.Println("GUID:", guid)
	fmt.Println("FileName:", fileName)
	fmt.Println("FileSystem:", fileSystem)
	fmt.Println("FileExtension:", fileExtension)
	fmt.Println("Authorization:", authorization)
	fmt.Println("Body:", body)

	generateFile(path, fileName, body)
}

func generateFile(path, fileName, body string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	// Construct the full file path
	fullPath := filepath.Join(path, fileName)

	// Write the content to the file
	err = os.WriteFile(fullPath, []byte(body), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("File generated at: %s\n", fullPath)
	return nil
}
