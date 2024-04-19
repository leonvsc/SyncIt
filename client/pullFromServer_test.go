package main

//
//import (
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"testing"
//)
//
//func TestGenerateGetHeader(t *testing.T) {
//	expected := "0200\nRequestType: GET\nContentType: text/plain\nContentLength: 21\nPath: files/test2.txt\nGUID: 0dada1dc-cb0a-463a-b028-7d04a8a5d3e4\nFileName: test2.txt\nFileSystem: Unix\nFileExtension: .txt\nAuthorization: null"
//	generated := generateGetHeader()
//	if generated != expected {
//		t.Errorf("generateGetHeader() = %s; want %s", generated, expected)
//	}
//}
//
//func TestGenerateFile(t *testing.T) {
//	// Define test data
//	testPath := "./testdata"
//	testFileName := "testfile.txt"
//	testBody := "This is a test file."
//
//	// Clean up test directory
//	defer os.RemoveAll(testPath)
//
//	// Call the function
//	err := generateFile(testPath, testFileName, testBody)
//	if err != nil {
//		t.Fatalf("generateFile returned error: %v", err)
//	}
//
//	// Check if file exists
//	fullPath := filepath.Join(testPath, testFileName)
//	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
//		t.Fatalf("generated file does not exist: %s", fullPath)
//	}
//
//	// Read file content
//	content, err := ioutil.ReadFile(fullPath)
//	if err != nil {
//		t.Fatalf("failed to read generated file: %v", err)
//	}
//
//	// Check file content
//	if string(content) != testBody {
//		t.Fatalf("file content mismatch. Expected: %s, Got: %s", testBody, string(content))
//	}
//}
