package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestProcessMessage(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		headerMap   map[string]string
		expected    []byte
		expectedErr error
	}{
		{
			name:  "Valid message with GET request",
			input: "GET HTTP/1.1 ContentLength: 5 Hello",
			headerMap: map[string]string{
				"RequestType":   "",
				"ContentLength": "",
			},
			expected: []byte("GET HTTP/1.1 ContentLength: 5 Hello"),
		},
		{
			name:  "Valid message with POST request",
			input: "POST HTTP/1.1 ContentLength: 10 UploadFile",
			headerMap: map[string]string{
				"RequestType":   "",
				"ContentLength": "",
			},
			expected: []byte("POST HTTP/1.1 ContentLength: 10 UploadFile"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tc.input))
			output := processMessage(scanner, tc.headerMap)

			if !bytes.Equal(output, tc.expected) {
				t.Errorf("expected: %v, got: %v", tc.expected, output)
			}
		})
	}
}

func TestParseRequestType(t *testing.T) {
	headerMap := make(map[string]string)

	parseRequestType("GET / HTTP/1.1", headerMap)
	if headerMap["RequestType"] != "GET" {
		t.Errorf("expected: GET, got: %v", headerMap["RequestType"])
	}

	parseRequestType("POST /upload HTTP/1.1", headerMap)
	if headerMap["RequestType"] != "POST" {
		t.Errorf("expected: POST, got: %v", headerMap["RequestType"])
	}
}

func TestParseHeader(t *testing.T) {
	headerMap := make(map[string]string)

	parseHeader("RequestType: GET", headerMap)
	if headerMap["RequestType"] != "GET" {
		t.Errorf("expected: GET, got: %v", headerMap["RequestType"])
	}

	parseHeader("ContentLength: 10", headerMap)
	if headerMap["ContentLength"] != "10" {
		t.Errorf("expected: 10, got: %v", headerMap["ContentLength"])
	}
}
