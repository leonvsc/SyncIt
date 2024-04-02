package main

import "testing"

func TestHeaderResponse(t *testing.T) {
	contentType := "application/json"
	contentLength := 100
	pathLocation := "/some/path"
	fileName := "example.txt"
	fileExtension := "txt"

	expectedResult := "0200\nPOST SFTP 1.0\nContentType: application/json\nContentLength: 100\nPath: /some/path\nGUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4\nFileName: example.txt\nFileSystem: Unix\nFileExtension: txt\nAuthorization: null"

	result := headerResponse(contentType, contentLength, pathLocation, fileName, fileExtension)

	if result != expectedResult {
		t.Errorf("Unexpected result, got: %s, want: %s", result, expectedResult)
	}
}
