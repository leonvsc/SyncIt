package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func readHeader(conn net.Conn) map[string]string {
	headerMap := make(map[string]string)
	buffer := make([]byte, 1024*1024) // 1 MB buffer size
	bytesRead, err := conn.Read(buffer)
	text := string(buffer[:bytesRead])

	getHeaderLength(text, headerMap)
	parseRequestType(text, headerMap)
	for _, line := range strings.Split(text, "\n") {
		parseHeader(line, headerMap)
	}
	if err != nil {
		fmt.Println(err)
	}

	return headerMap
}

func getHeaderLength(header string, headerMap map[string]string) {
	headerLength := make([]byte, 4)
	copy(headerLength, header[:4])
	// Convert the header length to an integer to remove leading zeros
	length, _ := strconv.Atoi(string(headerLength))
	headerMap["HeaderLength"] = strconv.Itoa(length)
}

func parseRequestType(line string, headerMap map[string]string) {
	switch {
	case strings.Contains(line, "GET"):
		headerMap["RequestType"] = "GET"
	case strings.Contains(line, "POST"):
		headerMap["RequestType"] = "POST"
	case strings.Contains(line, "PUT"):
		headerMap["RequestType"] = "PUT"
	case strings.Contains(line, "DELETE"):
		headerMap["RequestType"] = "DELETE"
	default:
	}
}

func parseHeader(line string, headerMap map[string]string) {
	switch {
	case strings.Contains(line, "RequestType"):
		headerMap["RequestType"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "ContentLength"):
		headerMap["ContentLength"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "GUID"):
		headerMap["GUID"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "Path"):
		headerMap["Path"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "FileSystem"):
		headerMap["FileSystem"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "FileName"):
		headerMap["FileName"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "FileExtension"):
		headerMap["FileExtension"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "Authorisation"):
		headerMap["Authorisation"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "MimeType"):
		headerMap["MimeType"] = strings.Split(line, ": ")[1]
	default:
	}
}
