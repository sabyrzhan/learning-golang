package decorator_pattern

import "fmt"
/*
Wrap target class with a class with additional functionality. This way we are extending the functionality using
composition. Though it is also possible with inheritance but not the primary way. The wrapper class must comply
to the signature of the wrapped class, to be able to also use stacked or multi layer composition.
 */

type FileManager interface {
	ReadFully() string
	Write(data string)
}

type TextFileManager struct {
	Path string
}

func (t *TextFileManager) ReadFully() string {
	return "this is the content"
}

func (t *TextFileManager) Write(data string) {
	fmt.Println(fmt.Sprintf("data='%s' written to file", data))
}

type EncryptedTextFileManager struct {
	manager FileManager
	encType string
}

func (e *EncryptedTextFileManager) ReadFully() string {
	encrypted := e.encryptContent(e.manager.ReadFully())
	return encrypted
}

func (e *EncryptedTextFileManager) Write(data string) {
	content := e.encryptContent(data)
	e.manager.Write(content)
}

func (e *EncryptedTextFileManager) encryptContent(content string) string {
	return fmt.Sprintf("%s encr: %s", e.encType, content)
}