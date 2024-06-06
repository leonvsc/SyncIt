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
	headerLength, _ := strconv.Atoi(headerMap["HeaderLength"])
	totalContent := 0
	for _, line := range strings.Split(text, "\n") {
		parseHeader(line, headerMap)
		if totalContent >= headerLength {
			break
		}
		totalContent = totalContent + len(line)
	}
	if err != nil {
		fmt.Println(err)
	}

	return headerMap
}

func getHeaderLength(header string, headerMap map[string]string) {
	headerLength := make([]byte, 4)
	if len(header) < 4 {
		fmt.Println("Header length is less than 4 bytes")
		fmt.Println("Header: " + header)
		fmt.Println("HeaderMap: " + fmt.Sprint(headerMap))
		return
	}
	copy(headerLength, header[:4])
	// Convert the header length to an integer to remove leading zeros
	length, err := strconv.Atoi(string(headerLength))
	if err != nil {
		fmt.Println(err)
		return
	}
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
	case strings.Contains(line, "AUTH"):
		headerMap["RequestType"] = "AUTH"
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
	case strings.Contains(line, "Authorization"):
		headerMap["Authorization"] = strings.Split(line, ": ")[1]
	case strings.Contains(line, "MimeType"):
		headerMap["MimeType"] = strings.Split(line, ": ")[1]
	default:
	}
}
