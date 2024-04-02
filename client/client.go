package main

import (
	"fmt"
	"net"
	"os"
)

var (
	serverAddr string = "localhost:50000" // Change this to your server address
	conn       net.Conn
	folderPath = "../local"
	filePath   = "../local/input.txt"
)

func main() {
	// Establish TCP connection
	if err := establishConnection(); err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	runMainMenu()
}

func establishConnection() error {
	var err error
	conn, err = net.Dial("tcp", serverAddr)
	return err
}

func closeConnection() {
	if conn != nil {
		conn.Close()
	}
}

func runMainMenu() {
	options := []string{"Synchroniseren", "Opties", "Afsluiten"}
	displayMenu(options)

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice. Please enter a valid option number.")
		return
	}

	switch choice {
	case 1:
		runSyncMenu()
	case 2:
		runOptiesMenu()
	case 3:
		fmt.Println("Exiting program...")
		closeConnection() // Close connection before exiting
		os.Exit(0)
	}
}

func runSyncMenu() {
	options := []string{"Push Folder", "Push File", "Pull from server", "Back"}
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
		pushFileToServer(conn, filePath)
	case 3:
		pullFromServer(conn)
	case 4:
		runMainMenu()
	}
}

func runOptiesMenu() {
	options := []string{"Sync server", "Back"}
	displayMenu(options)

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice. Please enter a valid option number.")
		return
	}

	switch choice {
	case 1:
		fmt.Println("Executing sync server option...")
		syncServer()
	case 2:
		runMainMenu()
	}
}

func displayMenu(options []string) {
	fmt.Println("Select an option:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Print("Enter option number: ")
}
