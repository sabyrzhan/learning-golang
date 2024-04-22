package factory_pattern

import (
	"fmt"
	"strconv"
	"strings"
)

type Page struct {
	title     string
	body      string
	list      []string
	paragraph string
	header    string
}

type TemplateFactory interface {
	Title() string
	Body() string
	List() string
	Paragraph() string
	Header(level int) string
	Render() string
}

func NewTemplateFactory(tplType string, page Page) TemplateFactory {
	switch tplType {
	case "html":
		return &HtmlTemplateFactory{page, 2}
	case "markdown":
		return &MarkdownTemplateFactory{page}
	case "json":
		return &JsonTemplateFactory{page}
	case "xml":
		return &XmlTemplateFactory{page}
	default:
		return &PlainTextTemplateFactory{page}
	}
}

type HtmlTemplateFactory struct {
	page Page
	currentIndent int
}

func (f *HtmlTemplateFactory) Title() string {
	return fmt.Sprintf("<title>%s</title>", f.page.title)
}

func (f *HtmlTemplateFactory) Body() string {
	return fmt.Sprintf("<div>%s</div>", f.page.body)
}

func (f *HtmlTemplateFactory) List() string {
	result := "<ul>\n"
	for _, l := range f.page.list {
		result += fmt.Sprintf("%s<li>%s</li>\n", strings.Repeat(" ", f.currentIndent * 4), l)
	}
	result += strings.Repeat(" ", f.currentIndent * 2) + "</ul>"

	return result
}

func (f *HtmlTemplateFactory) Paragraph() string {
	return fmt.Sprintf("<p>%s</p>", f.page.paragraph)
}

func (f *HtmlTemplateFactory) Header(level int) string {
	if level < 1 || level > 6 {
		panic("level must be between 1 and 6")
	}

	return fmt.Sprintf("<h%d>%s</h%d>", level, f.page.header, level)
}

func (f *HtmlTemplateFactory) Render() string {
	result := `
<html>
	%s
<body>
	%s
	%s
	%s
	%s
</body>
	
`

	return fmt.Sprintf(result, f.Title(), f.Header(1), f.Paragraph(), f.Body(), f.List())
}

type MarkdownTemplateFactory struct {
	page Page
}

func (f *MarkdownTemplateFactory) Title() string {
	return fmt.Sprintf("# %s", f.page.title)
}

func (f *MarkdownTemplateFactory) Body() string {
	return fmt.Sprintf("%s", f.page.body)
}

func (f *MarkdownTemplateFactory) List() string {
	var result string
	for _, l := range f.page.list {
		result += fmt.Sprintf("- %s\n", l)
	}

	return result
}

func (f *MarkdownTemplateFactory) Paragraph() string {
	return fmt.Sprintf("### %s", f.page.paragraph)
}

func (f *MarkdownTemplateFactory) Header(level int) string {
	if level < 1 || level > 3 {
		panic("level must be between 1 and 3")
	}

	return fmt.Sprintf("%s %s", strings.Repeat("#", level), f.page.header)
}

func (f *MarkdownTemplateFactory) Render() string {
	result := fmt.Sprintf("%s\n", f.Title())
	result += fmt.Sprintf("%s\n", f.Header(1))
	result += fmt.Sprintf("%s\n", f.Paragraph())
	result += fmt.Sprintf("%s\n", f.Body())
	result += fmt.Sprintf("%s\n", f.List())

	return result
}

type JsonTemplateFactory struct {
	page Page
}

func (f *JsonTemplateFactory) Title() string {
	return fmt.Sprintf(`title: %s`, strconv.Quote(f.page.title))
}

func (f *JsonTemplateFactory) Body() string {
	return fmt.Sprintf(`body: %s`, strconv.Quote(f.page.body))
}

func (f *JsonTemplateFactory) List() string {
	var items []string
	for _, l := range f.page.list {
		items = append(items, fmt.Sprintf(`%s`, strconv.Quote(l)))
	}

	return fmt.Sprintf(`list: [%s]`, strings.Join(items, ", "))
}

func (f *JsonTemplateFactory) Paragraph() string {
	return fmt.Sprintf(`paragraph: %s`, strconv.Quote(f.page.paragraph))
}

func (f *JsonTemplateFactory) Header(level int) string {
	return fmt.Sprintf(`header: %s`, strconv.Quote(f.page.header))
}

func (f *JsonTemplateFactory) Render() string {
	result := "{\n"
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Title())
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Header(1))
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Paragraph())
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Body())
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.List())
	result += "}"

	return result
}

type XmlTemplateFactory struct {
	page Page
}
func (f *XmlTemplateFactory) Title() string {
	return fmt.Sprintf("<title>%s</title", f.page.title)
}
func (f *XmlTemplateFactory) Body() string {
	return fmt.Sprintf(`<text>%s</text>`, f.page.body)
}
func (f *XmlTemplateFactory) List() string {
	result := "<items>\n"
	for _, l := range f.page.list {
		result += fmt.Sprintf("%s<item>%s</item>\n", strings.Repeat(" ", 4), l)
	}
	result += strings.Repeat(" ", 2) + "</items>"

	return result
}

func (f *XmlTemplateFactory) Paragraph() string {
	return fmt.Sprintf("<paragraph>%s</paragraph", f.page.paragraph)
}

func (f *XmlTemplateFactory) Header(level int) string {
	if level < 1 || level > 6 {
		panic("level must be between 1 and 6")
	}

	return fmt.Sprintf("<h%d>%s</h%d>", level, f.page.header, level)
}

func (f *XmlTemplateFactory) Render() string {
	result := "<xml>\n"
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Title())
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Header(1))
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Paragraph())
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.Body())
	result += fmt.Sprintf("%s%s\n", strings.Repeat(" ", 2), f.List())
	result += "</xml>"

	return result
}

type PlainTextTemplateFactory struct {
	page Page
}

func (f *PlainTextTemplateFactory) Title() string {
	return f.page.title
}

func (f *PlainTextTemplateFactory) Body() string {
	return f.page.body
}

func (f *PlainTextTemplateFactory) List() string {
	var result string
	for _, l := range f.page.list {
		result += l + "\n"
	}

	return result
}

func (f *PlainTextTemplateFactory) Paragraph() string {
	return f.page.paragraph
}

func (f *PlainTextTemplateFactory) Header(level int) string {
	return f.page.header
}

func (f *PlainTextTemplateFactory) Render() string {
	result := fmt.Sprintf("%s\n", f.Title())
	result += fmt.Sprintf("%s\n", f.Header(1))
	result += fmt.Sprintf("%s\n", f.Paragraph())
	result += fmt.Sprintf("%s\n", f.Body())
	result += fmt.Sprintf("%s\n", f.List())

	return result
}
