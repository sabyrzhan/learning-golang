package iterator_pattern

// We can use 3 types of iterators:
// 1. Using range over built-int collections like slices and maps
// 2. Using channel as values generator
// 3. Creating custom Iterator data structure that manages the current state internally

type Data struct {
	Line1 string
	Line2 string
	Line3 string
}

func (d *Data) IteratorWithSlices() []string {
	return []string{d.Line1, d.Line2, d.Line3}
}

func (d *Data) IteratorWithGenerator() chan string {
	out := make(chan string)
	go func(){
		if d.Line1 != "" {
			out <- d.Line1
		}

		if d.Line2 != "" {
			out <- d.Line2
		}

		if d.Line3 != "" {
			out <- d.Line3
		}

		defer close(out)
	}()

	return out
}

type CustomDataIterator struct {
	data Data
	Value string
	index int
}

func NewCustomDataIterator(d Data) *CustomDataIterator {
	return &CustomDataIterator{data: d, Value: "", index: -1}
}

func (c *CustomDataIterator) Next() bool {
	if c.index >= 3 {
		return false
	}

	c.index = c.index + 1

	switch c.index {
	case 0:
		c.Value = c.data.Line1
	case 1:
		c.Value = c.data.Line2
	case 2:
		c.Value = c.data.Line3
	default:
		return false
	}

	return true
}