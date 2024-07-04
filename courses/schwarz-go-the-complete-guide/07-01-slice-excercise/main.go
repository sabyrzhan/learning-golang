package main

import "fmt"

func main() {
	hobbies := [3]string{"hobby1", "hobby2", "hobby3"}
	fmt.Println(hobbies) // 1

	fmt.Println(hobbies[:1]) // 2
	fmt.Println(hobbies[1:]) //

	subhobbies1 := hobbies[:2] //3
	subhobbies2 := hobbies[0:2]
	fmt.Println(subhobbies1)
	fmt.Println(subhobbies2)

	subhobbies3 := subhobbies1[1:3] //4
	fmt.Println(subhobbies3)

	goals := []string{"goal1", "goal2"} //5
	goals[1] = "goal2modified"
	goals = append(goals, "goal3")
	fmt.Println(goals)

	products := []Product{ //6
		{
			1, "title1", 10.20,
		},
		{
			2, "title2", 30.30,
		},
	}
	products = append(products, Product{3, "title3", 40.40})
	fmt.Println(products)
}

type Product struct {
	id    int
	title string
	price float64
}
