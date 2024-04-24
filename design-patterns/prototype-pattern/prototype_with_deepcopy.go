package prototype_pattern

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

type author struct {
	Name         string
	MobilePhones []string
}

type Document struct {
	Author *author
	Data    string
	History *Document
	Version int
}

// NewDocumentWithHistory creates Document with history using DeepCopy (Prototype pattern).
// Each history is the clone of the previous version
func NewDocumentWithHistory() *Document {
	docVersion1 := Document{}
	docVersion1.Author = &author{"New author", []string{"1111"}}
	docVersion1.Version = 1
	docVersion1.Data = "Document data. Newly written with version 1"

	docVersion2 := docVersion1.DeepCopy()
	docVersion2.Data = "Document data. Newly written with version 2"
	docVersion2.History.Author.Name = fmt.Sprintf(docVersion2.History.Author.Name + " (Archived in version %d)", docVersion2.History.Version)
	docVersion2.AddPhoneNumber("222")

	docVersion3 := docVersion2.DeepCopy()
	docVersion3.Data = "Document data. Newly written with version 3"
	docVersion3.History.Author.Name = fmt.Sprintf(docVersion3.History.Author.Name + " (Archived in version %d)", docVersion3.History.Version)
	docVersion3.AddPhoneNumber("333")

	return docVersion3
}

// DeepCopy main function that makes deep clone of the document using gob.Encoder and gob.Decoder
func (d *Document) DeepCopy() *Document {
	fn := func() *Document {
		buffer := new(bytes.Buffer)
		encoder := gob.NewEncoder(buffer)
		err := encoder.Encode(d)
		if err != nil {
			panic(err)
		}

		decoder := gob.NewDecoder(buffer)
		result := new(Document)
		err = decoder.Decode(result)
		if err != nil {
			panic(err)
		}

		return result
	}

	result := fn()
	result.History = fn()
	result.Version = result.Version + 1

	return result
}

func (d *Document) AddPhoneNumber(phone string) {
	d.Author.MobilePhones = append(d.Author.MobilePhones, phone)
}

func (d *Document) string(nestLevel int) string {
	if d == nil {
		return ""
	}
	var result string
	result += fmt.Sprintf("%sAuthor: %s\n", strings.Repeat(" ", nestLevel * 2), d.Author.Name)
	result += fmt.Sprintf("%sMobile phones: %s, address: %p\n", strings.Repeat(" ", nestLevel * 2), strings.Join(d.Author.MobilePhones, ", "), d.Author.MobilePhones)
	result += fmt.Sprintf("%sData: %s\n", strings.Repeat(" ", nestLevel * 2), d.Data)
	result += fmt.Sprintf("%sVersion: %d\n", strings.Repeat(" ", nestLevel * 2), d.Version)
	result += fmt.Sprintf("%sHistory address: %p\n", strings.Repeat(" ", nestLevel * 2), d.History)
	result += d.History.string(nestLevel + 1)

	result += "\n"

	return result
}

func (d *Document) String() string {
	return d.string(0)
}
