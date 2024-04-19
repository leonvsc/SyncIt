package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func createHeaders(requestString string) string {
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
