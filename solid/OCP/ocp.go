package OCP

// OCP states the class or module should be open for extension but closed for modification
// If there is a new requirement to add or update the classes it is better to extend for example using extra design patterns.
// Below we have products. And we are adding filters by extending FilterSpec interface which is implemented
// using Specification design pattern. Whenever we want to add new filter we just create new one by implementing FilterSpec.
// More over we are creating CompositeFilter using Composite design pattern to combine multiple filter and apply as a single unit

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Color int

const (
	Black Color = iota
	White
	Green
	Blue
)

type Product struct {
	Name  string
	Size  Size
	Color Color
}

type FilterSpec interface {
	IsSatisfied(product Product) bool
}

type FilterByColor struct {
	color Color
}

func (f FilterByColor) IsSatisfied(p Product) bool {
	return f.color == p.Color
}

type FilterBySize struct {
	size Size
}

func (f FilterBySize) IsSatisfied(p Product) bool {
	return f.size == p.Size
}

type CompositeFilter struct {
	filters []FilterSpec
}

func (f CompositeFilter) IsSatisfied(product Product) bool {
	for _, filter := range f.filters {
		if !filter.IsSatisfied(product) {
			return false
		}
	}

	return true
}

func FilterProducts(products []Product, filter FilterSpec) []Product {
	filtered := make([]Product, 0, len(products))
	for _, product := range products {
		if filter.IsSatisfied(product) {
			filtered = append(filtered, product)
		}
	}

	return filtered
}
