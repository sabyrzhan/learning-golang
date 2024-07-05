package main

import "fmt"

type Data struct {
	ID    int
	Title string
}

func (d *Data) String() string {
	return fmt.Sprintf("ID: %d, Title: %s", d.ID, d.Title)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	printer := createPrinter(nums)
	printer()
	nums[0] = 6
	printer()

	data := &Data{1, "Some title"}
	structPrinter := createStructPrinter(data)
	structPrinter()
	data.ID = 2
	structPrinter()
}

func createStructPrinter(data *Data) func() {
	return func() {
		fmt.Println(data)
	}
}

func createPrinter(numbers []int) func() {
	return func() {
		fmt.Println(numbers)
	}
}
