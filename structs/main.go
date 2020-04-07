package main

import "fmt"

func main() {

	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(5)

	fmt.Println("stack")
	fmt.Println(stack.ToArray())
	fmt.Println(stack.Pop())
	fmt.Println(stack.ToArray())

	queue := &Queue{}

	queue.Push(7)
	queue.Push(8)
	queue.Push(9)

	fmt.Println("queue")
	fmt.Println(queue.ToArray())
	fmt.Println(queue.Pop())
	fmt.Println(queue.ToArray())
}
