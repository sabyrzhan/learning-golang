package file

import (
	"bufio"
	"encoding/json"
	"os"
	"price-calculator/types"
	"strconv"
)

type File struct {
	path    string
	outPath string
}

func NewFile(path string, outPath string) *File {
	return &File{path, outPath}
}

func (f *File) ReadPrices() []*types.Price {
	file, err := os.Open(f.path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]float64, 0)
	for scanner.Scan() {
		float, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		result = append(result, float)
	}

	return types.NewPriceFromValues(result)
}

func (f *File) WritePricesWithTaxes(pricesWithTax []*types.PricesWithTax) {
	file, err := os.Create(f.outPath)
	if err != nil {
		panic(err)
	}

	type TaxJson struct {
		Tax    float64   `json:"tax"`
		Prices []float64 `json:"prices"`
	}

	jsonData := make([]TaxJson, 0)
	for _, priceTax := range pricesWithTax {
		prices := make([]float64, 0)
		for _, price := range priceTax.Prices {
			prices = append(prices, price.GetValue())
		}
		jsonData = append(jsonData, TaxJson{Tax: priceTax.Tax.GetValue(), Prices: prices})
	}
	marshal, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(marshal)
	if err != nil {
		panic(err)
	}
}
