package main

import (
	"fmt"
	"os"
	"strconv"
)

func readBalanceFromFile() float64 {
	balance, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println("WARN: Failed to read file. Returning 1000")
		return 1000
	}

	parsed, err := strconv.ParseFloat(string(balance), 64)
	if err != nil {
		fmt.Println("WARN: Failed to parse float. Returning 1000")
		return 1000
	}

	return parsed
}

func writeBalanceToFile(balance float64) {
	err := os.WriteFile("balance.txt", []byte(fmt.Sprintf("%.2f", balance)), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var accountBalance = readBalanceFromFile()
	fmt.Println("Welcome to the Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("You choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
			} else {
				accountBalance += depositAmount
				fmt.Printf("Your updated balance is: %.2f\n", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
			} else if accountBalance-withdrawAmount < 0 {
				fmt.Printf("Withdraw amount cannot be larger than %.2f amount\n", accountBalance)
			} else {
				accountBalance -= withdrawAmount
				fmt.Printf("Your updated balance is: %.2f\n", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		default:
			fmt.Println("Bye. Thank you for choosing our bank!")
			return
		}
	}
}
