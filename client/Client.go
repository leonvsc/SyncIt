package main

import (
	"fmt"
	"os"
)

func main() {
	runMainMenu()
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
		sync()
	case 2:
		runOptiesMenu()
	case 3:
		fmt.Println("Exiting program...")
		os.Exit(0)
	}
}

func runOptiesMenu() {
	options := []string{"Sync server", "Program", "Back"}
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
		// Add your code for sync server option here
	case 2:
		fmt.Println("Executing program option...")
		// Add your code for program option here
	case 3:
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
