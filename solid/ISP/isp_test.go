package ISP

import "testing"

func TestISP_MultiFunctionDevice(t *testing.T) {
	doc := Document{}
	mfd := MultiFunctionDevice{}
	mfd.Print(doc)
	mfd.Scan(doc)
	mfd.Fax(doc)
}

func TestISP_OldPrinter(t *testing.T) {
	doc := Document{}
	old := OldPrinter{}
	old.Print(doc)
}
