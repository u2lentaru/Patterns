package main

import "fmt"

func main() {

	stack := &Stack{}

	stack.Push(5)
	stack.Push(5)
	stack.Push(5)

	//fmt.Println(list.Get(3))
	fmt.Println(stack.ToArray())
	fmt.Println(stack.Pop())
	fmt.Println(stack.ToArray())

}
