package main

import (
	"bufio"
	"strconv"
	"strings"
)

func processMessage(scanner *bufio.Scanner, headerMap map[string]string) []byte {
	var message []byte
	count := 0
	length := 0

	for scanner.Scan() {
		if scanner.Bytes() == nil || len(scanner.Bytes()) <= 4 {
			continue
		}

		line := scanner.Text()

		if count == 0 {
			headerLength := make([]byte, 4)
			copy(headerLength, line[:4])
			HeaderLength, _ = strconv.Atoi(string(headerLength))
			parseRequestType(line, headerMap)
		}

		count += len(line)
		parseHeader(line, headerMap)

		if headerMap["ContentLength"] != "" {
			contentLength, _ := strconv.Atoi(headerMap["ContentLength"])
			length = contentLength + HeaderLength
		}

		message = append(message, scanner.Bytes()...)

		if count == length {
			break
		}
	}
	return message
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
	case strings.Contains(line, "PUTHEALTH"):
		headerMap["RequestType"] = "PUTHEALTH"
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
		headerMap["FileName"] = strings.Split(line, ":")[1]
	case strings.Contains(line, "FileExtension"):
		headerMap["FileExtension"] = strings.Split(line, ":")[1]
	case strings.Contains(line, "Authorization"):
		headerMap["Authorization"] = strings.Split(line, ":")[1]
	case strings.Contains(line, "MimeType"):
		headerMap["MimeType"] = strings.Split(line, ":")[1]
	default:
	}
}
