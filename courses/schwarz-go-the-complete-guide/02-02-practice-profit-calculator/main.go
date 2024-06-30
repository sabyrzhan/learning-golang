package main

import (
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, tax float64
	revenue = readInput("Revenue: ")
	if revenue <= 0 {
		fmt.Println("Revenue must be greater than zero.")
		return
	}
	expenses = readInput("Expenses: ")
	if expenses <= 0 {
		fmt.Println("Expenses must be greater than zero.")
		return
	}
	tax = readInput("Tax rate: ")
	if tax <= 0 {
		fmt.Println("Tax rate must be greater than zero.")
		return
	}

	ebt, eat, ratio := calculate(revenue, expenses, tax)
	result := fmt.Sprintf("Earnings before tax (EBT): %.2f\n", ebt)
	result += fmt.Sprintf("Earnings after tax (EAT): %.2f\n", eat)
	result += fmt.Sprintf("Ratio: %.2f", ratio)

	err := os.WriteFile("result.txt", []byte(result), 0644)
	if err != nil {
		fmt.Println("ERROR: error writing result.txt", err.Error())
	}
	fmt.Println(result)
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
