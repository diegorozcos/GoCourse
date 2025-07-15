package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structsPractice/note"
	"example.com/structsPractice/todo"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the todo not successful")
		return
	}

	fmt.Println("Saving the todo successful")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note was not successful")
		return
	}

	fmt.Println("Saving the note successful")
}
