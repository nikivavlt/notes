package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const appName string = "Notes"
const appName = "Notes"

type Note struct {
	ID   int
	Text string
}

// A method is a function associated with a type, value receiver /pointer receiver
func (n Note) Print() {
	fmt.Printf("%d. %s\n", n.ID, n.Text)
}

func main() {
	// first explain array and why not? arrays length is part of its type
	notes := []Note{}
	globalID := 1

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

		// simplify this
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
			fmt.Println("   add <text>")
			fmt.Println("   show <id>")
			fmt.Println("   exit")

		case "add":
			// Explain panic on inputParts[1]
			if len(inputParts) < 2 {
				fmt.Println("Usage: add <text>")
			} else {
				notes = append(
					notes,
					Note{
						ID:   globalID,
						Text: strings.Join(inputParts[1:], " "),
					})

				globalID += 1

				fmt.Println("Note added successfully.")
			}

		case "list":
			if len(notes) == 0 {
				fmt.Println("No notes in the list.")
				continue
			}

			// fmt.Println(notes)
			fmt.Println("List of notes:")
			for _, note := range notes {
				note.Print()
			}

		case "show":
			id, err := strconv.Atoi(inputParts[1])
			if err != nil {
				fmt.Println("Invalid note ID.")
				continue
			}

			note, found := findNote(id, notes)
			if !found {
				fmt.Println("Note not found.")
				return
			}

			note.Print()

		case "exit":
			fmt.Println("Goodbye!")
			// break - instead since it ends loop, but not fn
			return

		default:
			fmt.Println(`Unknown command. Type "help".`)
		}
	}
}

func findNote(id int, notes []Note) (Note, bool){
	for _, note := range notes {
		if note.ID == id {
			return note, true
		}
	}

	return Note{}, false
}
