package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
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

	// Convert the request string to a byte array
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
	result := paddedLastValue + "\n" + requestString

	_, err := conn.Write([]byte(result))
	if err != nil {
		fmt.Println("Error with authorization:", err)
		return
	}

	//TODO: Receive response code from server and process the code.
}
