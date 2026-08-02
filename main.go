package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// const appName string = "Notes"
const appName = "Notes"

func main() {
	// first explain array and why not?
	notes := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println("Welcome to", appName+" app!")
	fmt.Printf("Welcome to %s app!\n", appName)
	fmt.Println(`Type "help" to see available commands.`)

	for {
		fmt.Print("> ")

		// fmt.Scan(&command)
		scanner.Scan()
		userInput := scanner.Text()
		// why not .Split() ?
		inputParts := strings.Fields(userInput)

		switch inputParts[0] {
		case "":
			continue

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("   help")
			fmt.Println("   list")
			fmt.Println("   add")
			fmt.Println("   exit")

		case "add":
			notes = append(notes, strings.Join(inputParts[1:], " "))

		case "list":
			// fmt.Println(notes)
			fmt.Println("List of notes:")
			for index, note := range notes {
				fmt.Printf("%d. %s\n", index, note)
			}

		case "exit":
			fmt.Println("Goodbye!")
			// break - instead since it ends loop, but not fn
			return

		default:
			fmt.Println(`Unknown command. Type "help".`)
		}
	}
}
