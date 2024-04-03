package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addFileToServer(msg []byte, headerMap map[string]string) string {

	contentLength, err := strconv.Atoi(headerMap["ContentLength"])
	message := msg[HeaderLength : HeaderLength+contentLength]

	fmt.Println("Writing file: ", headerMap["FileName"])
	filename := strings.TrimSpace(headerMap["FileName"])

	err = os.WriteFile(
		filename,
		message,
		0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	return ""
}
