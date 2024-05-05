package decorator_pattern

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	textFileManager := &TextFileManager{"/path/to/file.txt"}
	encryptedTextFileManager := EncryptedTextFileManager{textFileManager, "DES"}
	fmt.Println(fmt.Sprintf("Reading: %s", encryptedTextFileManager.ReadFully()))
	fmt.Println("Writing:")
	encryptedTextFileManager.Write("some new data")
}
