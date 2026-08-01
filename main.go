package main

import "fmt"

// const appName string = "Notes"
const appName = "Notes"

func main() {
	// var command string
	command := ""

	// fmt.Println("Welcome to", appName+" app!")
	fmt.Printf("Welcome to %s app!\n", appName)
	fmt.Println(`Type "help" to see available commands.`)
	fmt.Print("> ")

	fmt.Scan(&command)

	if command == "help" {
		fmt.Println("\nAvailable commands:")
		fmt.Println("	help")
	} else {
		fmt.Println(`Unknown command. Type "help".`)
	}
}
