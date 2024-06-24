package ISP

import "fmt"

// ISP states interface should not provide too many functions, but segregated into multiple separate interfaces.
// This will let classes to implement the functions that they really need instead of implementing all of them or
// leaving with empty default body that are not required.
// Below I separated multi function device interface into separate Printer, Scanner and Fax interfaces, so I can
// aggregate all of them in multi function device and use Printer only for printing device.

type Document struct {
}

type Printer interface {
	Print(doc Document)
}

type Scanner interface {
	Scan(doc Document)
}

type Fax interface {
	Fax(doc Document)
}

type OldPrinter struct {
}

func (o *OldPrinter) Print(doc Document) {
	fmt.Println("OldPrinter is printing. Capable to only print")
}

type MultiFunctionCopier interface {
	Printer
	Scanner
	Fax
}

type MultiFunctionDevice struct{}

func (d *MultiFunctionDevice) Print(doc Document) {
	fmt.Println("MultiFunctionDevice is printing")
}

func (d *MultiFunctionDevice) Scan(doc Document) {
	fmt.Println("MultiFunctionDevice is scanning")
}

func (d *MultiFunctionDevice) Fax(doc Document) {
	fmt.Println("MultiFunctionDevice is faxing")
}
