package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Create a new scanner to read input from the command line
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type ':' to open the command prompt.")

	// Start a goroutine to save periodically
	go savePeriodically()

	// Infinite loop to continuously read user input
	for {
		// Scan for the next token (which should be a line)
		scanner.Scan()

		// Get the text that was scanned
		command := scanner.Text()

		// Check if the command starts with ':'
		if strings.HasPrefix(command, ":") {
			// Open command prompt
			processCommand(command[1:])
		} else if command != "" {
			fmt.Println("Unknown command. Type ':' to open the command prompt.")
		}
	}
}

// Process commands entered in the command prompt
func processCommand(command string) {
	switch command {
	case "save":
		fmt.Println("Saving...")
		// Implement save logic here
	case "exit":
		fmt.Println("Exiting command prompt.")
		os.Exit(0)
	default:
		fmt.Println("Unknown command:", command)
	}
}

// Save something periodically
func savePeriodically() {
	for {
		time.Sleep(10 * time.Second) // Sleep for 10 seconds
		fmt.Println("Automatically saving...")
		// Implement save logic here
	}
}
