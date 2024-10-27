package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

var id int = 0

type Contact struct {
	Id    int
	Name  string
	Email string
}

func newContact(name string, email string) Contact {
	id++
	return Contact{Name: name, Email: email, Id: id}
}

type Contacts = []Contact
type Data struct {
	Contacts
}

func (d Data) indexOf(id int) int {
	for i, c := range d.Contacts {
		if c.Id == id {
			return i
		}
	}

	return -1
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "john@example.com"),
			newContact("Jane", "jane@example.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func (d Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}

	return false
}

type Page struct {
	Data Data
	Form FormData
}

func newPageData() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplates()
	e.Static("/images", "images")
	e.Static("/css", "css")

	page := newPageData()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"
			return c.Render(http.StatusUnprocessableEntity, "form", formData)
		}
		contact := newContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)
		_ = c.Render(http.StatusOK, "form", newFormData())
		return c.Render(http.StatusOK, "oob-contact", contact)
	})

	e.DELETE("/contacts/:id", func(c echo.Context) error {
		time.Sleep(3 * time.Second)
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}
		index := page.Data.indexOf(id)
		if index == -1 {
			return c.String(http.StatusNotFound, "Contact not found")
		}

		page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
