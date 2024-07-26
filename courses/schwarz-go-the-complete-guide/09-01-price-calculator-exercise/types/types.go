package types

type Price struct {
	Value float64
}

type Tax struct {
	Value float64
}

func NewPrice(value float64) *Price {
	return &Price{value}
}
func (p *Price) GetValue() float64 {
	return p.Value
}
func NewPriceFromValues(values []float64) []*Price {
	prices := make([]*Price, len(values))
	for i, v := range values {
		prices[i] = NewPrice(v)
	}

	return prices
}

func NewTax(value float64) *Tax {
	return &Tax{value}
}
func (p *Tax) GetValue() float64 {
	return p.Value
}
func NewTaxFromValues(values []float64) []*Tax {
	taxes := make([]*Tax, len(values))
	for i, v := range values {
		taxes[i] = NewTax(v)
	}

	return taxes
}

type PricesWithTax struct {
	Tax    *Tax
	Prices []*Price
}
