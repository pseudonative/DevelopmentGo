package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game")

	var name string
	fmt.Printf("Enter your name: \n")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game", name)
}
