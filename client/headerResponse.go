package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func headerResponse(contentType string, contentLength int, pathLocation string, fileName string, fileExtension string) string {
	// Construct the request string
	requestString := fmt.Sprintf(`POST SFTP 1.0
ContentType: %s
ContentLength: %d
Path: %s
GUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4
FileName: %s
FileSystem: Unix
FileExtension: %s
Authorization: null`, contentType, contentLength, pathLocation, fileName, fileExtension)

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
	return result
}
