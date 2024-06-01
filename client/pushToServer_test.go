package main

//
//import (
//	"mime"
//	"testing"
//)
//
//func TestGetContentLength(t *testing.T) {
//	tests := []struct {
//		name           string
//		bodyResponse   []byte
//		expectedLength int
//	}{
//		{
//			name:           "Empty body",
//			bodyResponse:   []byte{},
//			expectedLength: 0,
//		},
//		{
//			name:           "Non-empty body",
//			bodyResponse:   []byte("Hello, world!"),
//			expectedLength: 13,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			length := getContentLength(test.bodyResponse)
//			if length != test.expectedLength {
//				t.Errorf("Expected length %d, got %d", test.expectedLength, length)
//			}
//		})
//	}
//}
//
//func TestGetFileName(t *testing.T) {
//	tests := []struct {
//		name             string
//		localFilePath    string
//		expectedFileName string
//	}{
//		{
//			name:             "Basic file path",
//			localFilePath:    "/path/to/file.txt",
//			expectedFileName: "file.txt",
//		},
//		{
//			name:             "File path with directory",
//			localFilePath:    "/path/to/some/directory/file.txt",
//			expectedFileName: "file.txt",
//		},
//		{
//			name:             "File path without extension",
//			localFilePath:    "/path/to/file",
//			expectedFileName: "file",
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			fileName := getFileName(test.localFilePath)
//			if fileName != test.expectedFileName {
//				t.Errorf("Expected file name %s, got %s", test.expectedFileName, fileName)
//			}
//		})
//	}
//}
//
//func TestGetFileExtension(t *testing.T) {
//	tests := []struct {
//		name              string
//		localFilePath     string
//		expectedExtension string
//	}{
//		{
//			name:              "Basic file path with extension",
//			localFilePath:     "/path/to/file.txt",
//			expectedExtension: "txt",
//		},
//		{
//			name:              "File path with multiple dots",
//			localFilePath:     "/path/to/some/directory/file.tar.gz",
//			expectedExtension: "gz",
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			extension := getFileExtension(test.localFilePath)
//			if extension != test.expectedExtension {
//				t.Errorf("Expected extension %s, got %s", test.expectedExtension, extension)
//			}
//		})
//	}
//}
//
//func TestGetContentType(t *testing.T) {
//	// Adding ".json" extension to the mime types for testing purpose
//	_ = mime.AddExtensionType(".json", "application/json")
//
//	tests := []struct {
//		name                string
//		filePath            string
//		expectedContentType string
//	}{
//		{
//			name:                "Text file",
//			filePath:            "file.txt",
//			expectedContentType: "text/plain; charset=utf-8",
//		},
//		{
//			name:                "HTML file",
//			filePath:            "file.html",
//			expectedContentType: "text/html; charset=utf-8",
//		},
//		{
//			name:                "JSON file",
//			filePath:            "file.json",
//			expectedContentType: "application/json",
//		},
//		{
//			name:                "Empty file path",
//			filePath:            "",
//			expectedContentType: "application/octet-stream",
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			contentType := getContentType(test.filePath)
//			if contentType != test.expectedContentType {
//				t.Errorf("Expected content type %s, got %s", test.expectedContentType, contentType)
//			}
//		})
//	}
//}
