package main

import "fmt"

func main() {
	var revenue, expenses, tax float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax rate: ")
	fmt.Scan(&tax)

	ebt := revenue - expenses
	eat := ebt - ebt*tax/100
	ratio := ebt / eat
	fmt.Println("Earnings before tax (EBT): ", ebt)
	fmt.Println("Earnings after tax (EAT): ", eat)
	fmt.Println("Ratio: ", ratio)
}
