package main

import (
	"bufio"
	"errors"
	"fmt"
	"note-app/note"
	"os"
	"strings"
)

func main() {
	title, content, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote := note.New(title, content)
	newNote.Display()
	newNote.SaveAsJSON("note_data.json")
}

func readInput() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	var title, content string

	fmt.Print("Note title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return "", "", err
	}
	title = strings.TrimSpace(title)

	fmt.Print("Note content: ")
	content, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return "", "", err
	}
	content = strings.TrimSpace(content)

	if title == "" || content == "" {
		return "", "", errors.New("Title and content are required")
	}

	return title, content, nil
}
