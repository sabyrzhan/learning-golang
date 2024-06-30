package main

import "fmt"

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthdate := getUserData("Enter your birth date (MM/DD/YYYY): ")

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
