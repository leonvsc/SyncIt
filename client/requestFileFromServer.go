package main

import "fmt"

func requestFile(filename string) string {

	requestString := fmt.Sprintf(`GET SFTP 1.0
	RequestType: GET
	ContentType: test
	ContentLength: test
	Path: test
	GUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4
	FileName: %s
	FileSystem: Unix
	FileExtension: test
	Authorization: test`, filename)

	request := createHeaderLength(requestString)
	return request
}
