package main

import "fmt"

func main() {
	queue := NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue.Values())
	res, _ := queue.Dequeue()
	println(*res)
	values := queue.Values()
	fmt.Println(values)
	res, _ = queue.Peek()
	fmt.Println(*res)
	fmt.Println(queue.Values())
	fmt.Println(queue.IsEmpty())
	queue.Clear()
	fmt.Println(queue.Values())

}
