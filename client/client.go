package main

import (
	"fmt"
	"net"
	"os"
)

var (
	serverAddr string
	conn       net.Conn
	folderPath = "../local"
)

func main() {
	runMainMenu()
}

func makeConnection() {
	// Establish TCP connection
	if err := establishConnection(); err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection:", err)
		}
	}(conn)
	authorization()
	runSyncMenu()
}

func establishConnection() error {
	var err error
	conn, err = net.Dial("tcp", serverAddr)
	return err
}

func closeConnection() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			return
		}
	}
}

func runMainMenu() {
	options := []string{"Connect to server", "Sync", "Options", "Quit"}
	displayMenu(options)

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice. Please enter a valid option number.")
		return
	}

	switch choice {
	case 1:
		makeConnection()
	case 2:
		runSyncMenu()
	case 3:
		runOptiesMenu()
	case 4:
		fmt.Println("Exiting program...")
		closeConnection() // Close connection before exiting
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please enter a valid option number.")
	}
}

func runSyncMenu() {
	options := []string{"Push Folder", "Push File", "Listen to server", "Back"}
	displayMenu(options)

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice. Please enter a valid option number.")
		return
	}

	switch choice {
	case 1:
		pushFolderToServer(conn, folderPath)
	case 2:
		fileToSync := getFileList()
		err := sendFile(fileToSync, conn)
		if err != nil {
			return
		}
	case 3:
		headerMap := readHeader(conn)
		handleRequest(headerMap, conn)
	case 4:
		runMainMenu()
	default:
		fmt.Println("Invalid choice. Please enter a valid option number.")
	}
}

func runOptiesMenu() {
	options := []string{"Set sync server", "Show current sync server", "Back"}
	displayMenu(options)

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice. Please enter a valid option number.")
		return
	}

	switch choice {
	case 1:
		syncServer()
	case 2:
		showCurrentServer()
	case 3:
		runMainMenu()
	default:
		fmt.Println("Invalid choice. Please enter a valid option number.")
	}
}

func displayMenu(options []string) {
	fmt.Println("Select an option:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Print("Enter option number: ")
}
