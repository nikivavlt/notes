package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// const appName string = "Notes"
const appName = "Notes"

type Storage struct {
	notes  []Note
	nextID int
}

func NewStorage() *Storage {
	return &Storage{nextID: 1}
}

func (storage *Storage) Add(text string) bool {
	note := Note{
		ID:   storage.nextID,
		Text: text,
	}

	storage.notes = append(storage.notes, note)
	storage.nextID += 1

	return true
}

func (storage *Storage) Find(id int) (*Note, int) {
	for index := range storage.notes {
		if storage.notes[index].ID == id {
			return &storage.notes[index], index
		}
	}

	// nil or &Note{}
	return nil, -1
}

// split to func (storage *Storage) findIndex(id int) int?
func (storage *Storage) Remove(id int) bool {
	_, index := storage.Find(id)

	if index == -1 {
		return false
	}

	// One more important detail: slices.Delete returns the new slice, so this reassignment is required:
	storage.notes = slices.Delete(storage.notes, index, index+1)
	return true
}

func (storage *Storage) Count() int {
	return len(storage.notes)
}

func (storage *Storage) List() []Note {
	notes := make([]Note, len(storage.notes))
	// why not double? I mean 0, 0, 0 ..., and then notes
	copy(notes, storage.notes)
	return notes
}

func (storage *Storage) Rename(id int, text string) bool {
	note, _ := storage.Find(id)
	if note == nil {
		return false
	}

	note.Rename(text)
	return true
}

type Note struct {
	ID   int
	Text string
}

// A method is a function associated with a type, value receiver /pointer receiver

func (n *Note) Rename(text string) {
	n.Text = text
}

// keep main small
func main() {
	// first explain array and why not? arrays length is part of its type
	storage := NewStorage()

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
		// function handlers split
		switch inputParts[0] {
		// case "":
		//	continue

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("   help")
			fmt.Println("   list")
			fmt.Println("   count")
			fmt.Println("   add <text>")
			fmt.Println("   remove <id>")
			fmt.Println("   show <id>")
			fmt.Println("   rename <id> <text>")
			fmt.Println("   exit")

		case "add":
			// The caller says:
			// Add this note.
			// not:
			// Append this struct into your internal slice and increment this counter.
			// That distinction becomes more important as complexity grows.

			// Explain panic on inputParts[1]
			if len(inputParts) < 2 {
				fmt.Println("Usage: add <text>")
				continue
			}

			storage.Add(strings.Join(inputParts[1:], " "))
			fmt.Println("Note added successfully.")

		case "remove":
			if len(inputParts) != 2 {
				fmt.Println("Usage: remove <id>")
				continue
			}

			id, err := strconv.Atoi(inputParts[1])
			if err != nil || id <= 0 {
				fmt.Println("Invalid note ID.")
				continue
			}

			removed := storage.Remove(id)
			if !removed {
				fmt.Println("Note not found.")
				continue
			}

			fmt.Println("Note removed successfully.")

		case "list":
			if storage.Count() == 0 {
				fmt.Println("No notes in the list.")
				continue
			}
			// leaks internal state if moved to struct method, do copy

			// fmt.Println(notes)
			fmt.Println("List of notes:")
			for _, note := range storage.List() {
				// this mixed business logic with presentation, cli must decide
				fmt.Printf("%d. %s\n", note.ID, note.Text)
			}

		case "count":
			// what is d?
			fmt.Printf("%d note(s) in the list.\n", storage.Count())

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

			// to one line condition with var inside if ?
			renamed := storage.Rename(id, strings.Join(inputParts[2:], " "))
			if renamed == false {
				fmt.Println("Note not found.")
				continue
			}

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

			// why needed, _??
			note, _ := storage.Find(id)
			if note == nil {
				fmt.Println("Note not found.")
				continue
			}

			fmt.Printf("%d. %s\n", note.ID, note.Text)

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
