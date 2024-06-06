package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
	if serverAddr == "" {
		serverAddr = ":50000"
	}
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
		runOptionsMenu()
	case 4:
		fmt.Println("Exiting program...")
		closeConnection() // Close connection before exiting
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please enter a valid option number.")
	}
}

func runSyncMenu() {
	options := []string{"Push Folder", "Push File", "Get File", "Listen to server", "Back"}
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
		break
	case 2:
		fileToSync := getFileList()
		err := sendFile(fileToSync, conn)
		if err != nil {
			return
		}
		break
	case 3:
		var fileName string
		fmt.Println("Enter filename:")
		_, err := fmt.Scanln(&fileName)
		if err != nil {
			fmt.Println("Error reading filename:", err)
			return
		}
		message := requestFile(fileName)
		conn.Write([]byte(message))

		buffer := make([]byte, 1024*1024) // 1 MB buffer size
		bytesRead, err := conn.Read(buffer)
		text := string(buffer[:bytesRead])
		if strings.Contains(text, "Statuscode: 404") {
			fmt.Println("File not found on server")
			return
		}

		headerMap := readHeader(conn)
		handleRequest(headerMap, conn)
		break
	case 4:
		headerMap := readHeader(conn)
		handleRequest(headerMap, conn)
		break
	case 5:
		runMainMenu()
		break
	default:
		fmt.Println("Invalid choice. Please enter a valid option number.")
		break
	}
}

func runOptionsMenu() {
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
