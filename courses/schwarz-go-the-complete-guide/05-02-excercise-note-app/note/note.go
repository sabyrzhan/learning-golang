package note

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type note struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}

func New(title, content string) *note {
	return &note{title, content, time.Now()}
}

func (n *note) SaveAsJSON(fileName string) error {
	js, err := json.MarshalIndent(n, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, js, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (n *note) Display() {
	fmt.Printf("You entered note with title %s and content:\n%s\n", n.Title, n.Content)
}
