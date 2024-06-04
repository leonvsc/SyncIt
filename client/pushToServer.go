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
			sendFile(entryPath, conn)
		} else {
			// If the entry is a file, push it to the server
			sendFile(entryPath, conn)

			// Add a delimiter between files (except for the last file)
			if i < len(entries)-1 {
				conn.Write([]byte("\n\n"))
			}
		}
	}

	// Add a delimiter after syncing the folder
	conn.Write([]byte("\n\n"))
	fmt.Println("Folder synced successfully.")
}
