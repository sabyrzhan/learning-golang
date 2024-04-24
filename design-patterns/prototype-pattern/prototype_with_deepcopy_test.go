package prototype_pattern

import (
	"fmt"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	doc := NewDocumentWithHistory()

	docCopy := doc.DeepCopy()
	docCopy.Author.Name = "Another new author"
	docCopy.AddPhoneNumber("77777")
	docCopy.History = doc


	fmt.Println("==== Doc ====")
	fmt.Println(doc)

	fmt.Println("==== Doc copy ====")
	fmt.Println(docCopy)
}