package types

type TaxCalculator struct {
}

func NewTaxCalculator() *TaxCalculator {
	return &TaxCalculator{}
}

func (c *TaxCalculator) calculate(tax float64, price float64) float64 {
	return price*tax/100 + price
}

func (c *TaxCalculator) ApplyTaxToPrice(tax *Tax, price *Price) *Price {
	return &Price{Value: c.calculate(tax.GetValue(), price.GetValue())}
}

func (c *TaxCalculator) ApplyTaxToPrices(tax *Tax, prices []*Price) []*Price {
	result := make([]*Price, len(prices))
	for i, price := range prices {
		result[i] = c.ApplyTaxToPrice(tax, price)
	}

	return result
}
