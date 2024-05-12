package main

import (
	"io"
	"net"
	"os"
)

func receiveFile(conn net.Conn, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024*1024) // 1 MB buffer size
	for {
		bytesRead, err := conn.Read(buffer)
		println("bytesRead: ", bytesRead)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		_, err = file.Write(buffer[:bytesRead])
		if err != nil {
			return err
		}
	}

	return nil
}
