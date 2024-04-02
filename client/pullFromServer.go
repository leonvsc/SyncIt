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

	//println(string(response[:n]))

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
	lines := strings.Split(response, "\n")

	// Process each line of the response
	var contentType, contentLength, path, guid, fileName, fileSystem, fileExtension, authorization string

	var body []string // Store body content
	var statusCode []string

	// Skip the first line since it's part of the header
	headerSkipped := false

	for _, line := range lines {
		// Check if the line contains ": "
		if strings.Contains(line, ": ") {
			// Split each line by ": "
			parts := strings.Split(line, ": ")
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
		} else {
			// If the header has been skipped, include the line in the body content
			if headerSkipped {
				body = append(body, line)
			} else {
				headerSkipped = true
				statusCode = append(statusCode, line)
			}
		}
	}

	// Join the body content
	bodyContent := strings.Join(body, "\n")
	statusCodeContent := strings.Join(statusCode, "\n")

	// Print the values
	fmt.Println(statusCodeContent)
	fmt.Println("ContentType:", contentType)
	fmt.Println("ContentLength:", contentLength)
	fmt.Println("Path:", path)
	fmt.Println("GUID:", guid)
	fmt.Println("FileName:", fileName)
	fmt.Println("FileSystem:", fileSystem)
	fmt.Println("FileExtension:", fileExtension)
	fmt.Println("Authorization:", authorization)
	fmt.Println("Body:", bodyContent)

	// Call function to generate file
	generateFile(path, fileName, bodyContent)
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