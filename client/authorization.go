package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var (
	username string
)

func authorization() {
	askUsername()
	authRequest()

}

func askUsername() {
	fmt.Println("Enter username:")
	_, err := fmt.Scanln(&username)
	if err != nil {
		return
	}
	fmt.Println("Username set to:", username)
}

func authRequest() {
	// Encode Username to Base64
	base64Username := base64.StdEncoding.EncodeToString([]byte(username))

	// Construct the request string
	requestString := fmt.Sprintf(`AUTH SFTP 1.0
Authorization: %s`, base64Username)

	result := createHeaders(requestString)

	_, err := conn.Write([]byte(result))
	if err != nil {
		fmt.Println("Error with authorization:", err)
		return
	}

	// Read response from the server
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	response = response[:n]

	// Extract status code from the response
	statusCode := extractStatusCode(response)

	// Check status code and act accordingly
	if statusCode == "200" {
		// Resume program
		fmt.Println("Authorization successful")
		runSyncMenu()
	} else {
		// Throw error code
		fmt.Println("Authorization error:", statusCode)
		runMainMenu()
	}
}

func extractStatusCode(response []byte) string {
	lines := strings.Split(string(response), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Statuscode:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "" // Status code not founding
}
