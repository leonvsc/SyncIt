package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func sendFile(headerMap map[string]string) []byte {
	file, err := os.Open(headerMap["Path"])
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	//parse file name
	fileName := file.Name()
	fileExtension := fileName[len(fileName)-4:]

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return nil
	}
	fmt.Println(bs)
	fmt.Println(stat.Size())
	header := createHeader(fileExtension, int(stat.Size()), fileName)
	headerLength := len(header)
	return []byte(strconv.Itoa(headerLength) + header + string(bs))
}

func sendOkToClient() []byte {
	response := "\n200: message received"
	responseLength := len(response)
	header := createHeader("", responseLength, "")
	headerLength := len(header)
	return []byte(strconv.Itoa(headerLength) + header + response)
}

func createHeader(fileExtension string, contentLength int, fileName string) string {
	header := fmt.Sprintf(
		`POST SFTP 1.0
		ContentType: %s
		ContentLength: %d
		Path: %s
		GUID: %s
		FileName: %s
		FileSystem: %s
		FileExtension: %s
		Authorization: %s`,
		fileExtension,
		contentLength,
		fileName,
		"0dada1dc-cb0a-463a-b028-7d04a8a5d3e4",
		fileName,
		"Unix",
		fileExtension,
		"Bearer")

	return header
}
