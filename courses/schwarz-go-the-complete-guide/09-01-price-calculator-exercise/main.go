package main

import (
	"price-calculator/file"
	"price-calculator/types"
)

func main() {
	f := file.NewFile("prices.txt", "result.json")
	priceValues := f.ReadPrices()
	taxes := []float64{30, 40, 50}
	taxValues := types.NewTaxFromValues(taxes)

	taxCalculator := types.NewTaxCalculator()
	result := make([]*types.PricesWithTax, 0)
	for _, taxValue := range taxValues {
		newPricesValuesWithTax := taxCalculator.ApplyTaxToPrices(taxValue, priceValues)
		result = append(result, &types.PricesWithTax{Tax: taxValue, Prices: newPricesValuesWithTax})
	}

	f.WritePricesWithTaxes(result)
}
