package main

import "fmt"

func syncServer() {
	fmt.Println("Enter server address:")
	fmt.Scanln(&serverAddr)
	fmt.Println("Server address set to:", serverAddr)
	runOptiesMenu()
}

func showCurrentServer() {
	fmt.Println("Server address set to:", serverAddr)
	runOptiesMenu()
}
