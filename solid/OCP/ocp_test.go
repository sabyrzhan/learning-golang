package OCP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOCP_ColorFilter(t *testing.T) {
	products := []Product{
		{Name: "MacBook Pro 14", Color: Black, Size: Medium},
		{Name: "MacBook Air 13", Color: White, Size: Small},
		{Name: "Lenovo Legion 3", Color: Black, Size: Large},
	}
	colorFilter := FilterByColor{Black}

	result := FilterProducts(products, colorFilter)

	expected := []Product{products[0], products[2]}
	assert.Equal(t, expected, result)
}

func TestOCP_SizeFilter(t *testing.T) {
	products := []Product{
		{Name: "MacBook Pro 14", Color: Black, Size: Medium},
		{Name: "MacBook Air 13", Color: White, Size: Small},
		{Name: "Lenovo Legion 3", Color: Black, Size: Large},
	}
	sizeFilter := FilterBySize{Small}

	result := FilterProducts(products, sizeFilter)

	expected := []Product{products[1]}
	assert.Equal(t, expected, result)
}

func TestOCP_CompositeFilter(t *testing.T) {
	products := []Product{
		{Name: "MacBook Pro 14", Color: Black, Size: Medium},
		{Name: "MacBook Air 13", Color: White, Size: Small},
		{Name: "Lenovo Legion 3", Color: Black, Size: Large},
	}
	compositeFilter := CompositeFilter{
		[]FilterSpec{FilterBySize{Large}, FilterByColor{Black}},
	}

	result := FilterProducts(products, compositeFilter)

	expected := []Product{products[2]}
	assert.Equal(t, expected, result)
}
