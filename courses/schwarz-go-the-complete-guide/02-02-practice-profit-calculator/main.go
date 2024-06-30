package main

import "fmt"

func main() {
	var revenue, expenses, tax float64
	revenue = readInput("Revenue: ")
	expenses = readInput("Expenses: ")
	tax = readInput("Tax rate: ")

	ebt, eat, ratio := calculate(revenue, expenses, tax)
	fmt.Println("Earnings before tax (EBT): ", ebt)
	fmt.Println("Earnings after tax (EAT): ", eat)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func readInput(question string) float64 {
	fmt.Print(question)
	var value float64
	fmt.Scan(&value)

	return value
}

func calculate(revenue, expenses, tax float64) (ebt float64, eat float64, ratio float64) {
	ebt = revenue - expenses
	eat = ebt - ebt*tax/100
	ratio = ebt / eat
	return
}
