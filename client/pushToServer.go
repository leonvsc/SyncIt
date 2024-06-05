package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func pushFolderToServer(conn net.Conn, folderPath string) {
	// Get a list of files and directories in the folder
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		return
	}

	for i, entry := range entries {
		// Construct the full path of the entry
		entryPath := filepath.Join(folderPath, entry.Name())

		if entry.IsDir() {
			// If the entry is a directory, recursively call pushFolderToServer
			err := sendFile(entryPath, conn)
			if err != nil {
				return
			}
		} else {
			// If the entry is a file, push it to the server
			err := sendFile(entryPath, conn)
			if err != nil {
				return
			}

			// Add a delimiter between files (except for the last file)
			if i < len(entries)-1 {
				_, err := conn.Write([]byte("\n\n"))
				if err != nil {
					return
				}
			}
		}
	}

	// Add a delimiter after syncing the folder
	_, err = conn.Write([]byte("\n\n"))
	if err != nil {
		return
	}
	fmt.Println("Folder synced successfully.")
}
