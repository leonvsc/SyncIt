package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseMessage(msg []byte) string {

	headerLength := msg[:4]
	println(string(headerLength))

	headerInt, err := strconv.Atoi(string(headerLength))
	println(headerInt)
	if err != nil {
		println("Error: Header does not contain a number: 400")
		return "Error: Header does not contain a number: 400\n"
	}

	if headerInt > len(msg) || headerInt < 4 {
		println("Error: Header length is incorrect: 400")
		return "Error: Header length is incorrect message: 400\n"
	}

	header := msg[4:headerInt]
	println(string(header))

	headerMap := make(map[string]string)

	message := msg[headerInt:] // till the end of the message (get via the content length)
	println(string(message))

	scanner := bufio.NewScanner(strings.NewReader(string(message)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line individually
		fmt.Println(line)
		switch {
		case strings.Contains(line, "RequestType"):
			headerMap["RequestType"] = strings.Split(line, ":")[1]
			fmt.Println("RequestType is...")
		case strings.Contains(line, "ContentLength"):
			headerMap["ContentLength"] = strings.Split(line, ":")[1]
			fmt.Println("ContentLength is...")
		case strings.Contains(line, "GUID"):
			headerMap["GUID"] = strings.Split(line, ":")[1]
			fmt.Println("GUID is...")
		case strings.Contains(line, "Path"):
			headerMap["Path"] = strings.Split(line, ":")[1]
			fmt.Println("Path is...")
		case strings.Contains(line, "FileSystem"):
			headerMap["FileSystem"] = strings.Split(line, ":")[1]
			fmt.Println("FileSystem is...")
		case strings.Contains(line, "FileName"):
			headerMap["FileName"] = strings.Split(line, ":")[1]
			fmt.Println("FileName is...")
		case strings.Contains(line, "FileExtension"):
			headerMap["FileExtension"] = strings.Split(line, ":")[1]
			fmt.Println("FileExtension is...")
		case strings.Contains(line, "Authorization"):
			headerMap["Authorization"] = strings.Split(line, ":")[1]
			fmt.Println("Authorization is...")
		case strings.Contains(line, "MimeType"):
			headerMap["MimeType"] = strings.Split(line, ":")[1]
			fmt.Println("MimeType is...")
		default:
			return "Error: Unknown header field: 400\n"
		}
	}
	println(headerMap)

	return ""
}
