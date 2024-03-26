package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func parseMessage(msg []byte) string {
	scanner := bufio.NewScanner(strings.NewReader(string(msg)))
	scanner.Split(bufio.ScanLines)

	headerLength := msg[:4]
	println(string(headerLength))

	headerNumber := headerContainsLength(headerLength)
	println(headerNumber)
	if !headerNumber {
		return "Error: Header does not contain a number: 400\n"
	}

	headerInt, err := strconv.Atoi(string(headerLength))
	println(headerInt)
	if err != nil {
		return "Error: Header does not contain a number: 400\n"
	}

	header := msg[4:headerInt]
	println(string(header))

	message := msg[headerInt:]
	println(string(message))

	for scanner.Scan() {
		line := scanner.Text()
		// Process each line individually
		fmt.Println(line)
	}

	return ""
}

func headerContainsLength(header []byte) bool {
	r := string(header)
	sep := 0

	for _, b := range r {
		if unicode.IsNumber(b) {
			continue
		}
		if b == '.' {
			if sep > 0 {
				return false
			}
			sep++
			continue
		}
		return false
	}

	return true
}
