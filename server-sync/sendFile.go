package main

import (
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024*1024) // 1 MB buffer size
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		_, err = conn.Write(buffer[:bytesRead])
		if err != nil {
			return err
		}
	}

	return nil
}
