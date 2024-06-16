package iterator_pattern

import (
	"fmt"
	"testing"
)

func TestBasicIterator_withSlices(t *testing.T) {
	d := Data{Line1: "line1", Line2: "line2", Line3: "line3"}
	for _, line := range d.IteratorWithSlices() {
		fmt.Println("Slice: " + line)
	}
}

func TestBasicIterator_withGenerator(t *testing.T) {
	d := Data{Line1: "line1", Line2: "line2", Line3: "line3"}
	for line := range d.IteratorWithGenerator() {
		fmt.Println("Generator: " + line)
	}
}

func TestBasicIterator_withCustomIterator(t *testing.T) {
	d := Data{Line1: "line1", Line2: "line2", Line3: "line3"}
	it := NewCustomDataIterator(d)
	for it.Next() {
		fmt.Println(it.Value)
	}
}