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

func (n *Note) Rename(text string) {
	n.Text = text
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
			fmt.Println("   count")
			fmt.Println("   add <text>")
			fmt.Println("   show <id>")
			fmt.Println("   rename <id> <text>")
			fmt.Println("   exit")

		case "add":
			// Explain panic on inputParts[1]
			if len(inputParts) < 2 {
				fmt.Println("Usage: add <text>")
			} else {

				note := Note{
					ID:   globalID,
					Text: strings.Join(inputParts[1:], " "),
				}

				notes = append(notes, note)

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

		case "count":
			// what is d?
			fmt.Printf("%d note(s) in the list.\n", len(notes))

		case "rename":
			if len(inputParts) < 3 {
				fmt.Println("Usage: rename <id> <text>")
				continue
			}

			id, err := strconv.Atoi(inputParts[1])
			if err != nil || id <= 0 {
				fmt.Println("Invalid note ID.")
				continue
			}

			note, found := findNote(notes, id)
			if !found {
				fmt.Println("Note not found.")
				continue
			}

			note.Rename(strings.Join(inputParts[2:], " "))
			fmt.Println("Note renamed successfully.")

		case "show":
			if len(inputParts) != 2 {
				fmt.Println("Usage: show <id>")
				continue
			}

			id, err := strconv.Atoi(inputParts[1])
			if err != nil || id <= 0 {
				fmt.Println("Invalid note ID.")
				continue
			}

			note, found := findNote(notes, id)
			if !found {
				fmt.Println("Note not found.")
				continue
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

	if err := scanner.Err(); err != nil {
		fmt.Println("input error:", err)
	}
}

// simplify to one return ? without bool?
func findNote(notes []Note, id int) (*Note, bool) {
	for index := range notes {
		if notes[index].ID == id {
			return &notes[index], true
		}
	}

	// nil or &Note{}
	return nil, false
}
