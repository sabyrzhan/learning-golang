package factory_pattern

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	page := Page{
		title: "Page title",
		body: "\"Body text\"",
		paragraph: "Text paragraph",
		header: "Text header",
		list: []string {"list 1", "list 2", "list 3"},
	}

	fmt.Println("=== Html page ===")
	htmlFactory := NewTemplateFactory("html", page)
	fmt.Println(htmlFactory.Render())

	fmt.Println("=== Markdown page ===")
	mardownFactory := NewTemplateFactory("markdown", page)
	fmt.Println(mardownFactory.Render())

	fmt.Println("\n\n=== JSON page ===")
	jsonFactory := NewTemplateFactory("json", page)
	fmt.Println(jsonFactory.Render())

	fmt.Println("\n\n=== XML page ===")
	xmlFactory := NewTemplateFactory("xml", page)
	fmt.Println(xmlFactory.Render())

	fmt.Println("\n\n=== PlainText page ===")
	plaintextFactory := NewTemplateFactory("default", page)
	fmt.Println(plaintextFactory.Render())
}
