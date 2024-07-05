package main

import (
	"fmt"
	"strconv"
)

func main() {
	printVariadicFunction(1, 2, 3, 4, 5)
	printVariadicFunction(6)

	otherNumbers := []int{11, 22, 33, 44}
	printVariadicFunction(1, otherNumbers...)
}

func printVariadicFunction(a int, b ...int) {
	fmt.Println("First param: " + strconv.Itoa(a))
	fmt.Println("Other params")
	for _, v := range b {
		fmt.Print(strconv.Itoa(v) + " ")
	}
	fmt.Println()
}
