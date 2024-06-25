package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, expectedReturnRate, years, inflationRate float64
	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	fmt.Print("Inflation rate: ")
	fmt.Scan(&inflationRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("%d\n", int(futureValue))
	fmt.Printf("%d\n", int(futureRealValue))
}
