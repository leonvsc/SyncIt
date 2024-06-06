package main

import "fmt"

func syncServer() {
	fmt.Println("Enter server address:")
	_, err := fmt.Scanln(&serverAddr)
	if err != nil {
		return
	}
	fmt.Println("Server address set to:", serverAddr)
	runOptionsMenu()
}

func showCurrentServer() {
	fmt.Println("Server address set to:", serverAddr)
	runOptionsMenu()
}
