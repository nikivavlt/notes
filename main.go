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
	// first explain array and why not? arrays length is part of its type
	notes := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println("Welcome to", appName+" app!")
	fmt.Printf("Welcome to %s app!\n", appName)
	fmt.Println(`Type "help" to see available commands.`)

	for {
		fmt.Print("> ")

		// fmt.Scan(&command)
		//scanner.Scan()
		if !scanner.Scan() { // Scan() returns false when input ends or an error occurs. A safer pattern is:
			break
		}

		userInput := scanner.Text()
		// why not .Split() ?
		inputParts := strings.Fields(userInput)

		if len(inputParts) == 0 {
			continue
		}
		
		// explain why panic possible here
		switch inputParts[0] {
		// case "":
		//	continue

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("   help")
			fmt.Println("   list")
			fmt.Println("   add")
			fmt.Println("   exit")

		case "add":
			// Explain panic on inputParts[1]
			if len(inputParts) < 2 {
				fmt.Println("Usage: add <text>")
			} else {
				notes = append(notes, strings.Join(inputParts[1:], " "))
				fmt.Println("Note added successfully.")
			}

		case "list":
			if len(notes) == 0 {
				fmt.Println("No notes in the list.")
				continue
			}

			// fmt.Println(notes)
			fmt.Println("List of notes:")
			for index, note := range notes {
				fmt.Printf("%d. %s\n", index+1, note)
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
