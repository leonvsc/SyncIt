package main

import (
	"encoding/base64"
	"fmt"
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
	fmt.Scanln(&username)
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

	//TODO: Receive response code from server and process the code.
}
