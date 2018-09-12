package main

import (
	"fmt"
	"go-data-structures-algorithms/linkedstack/linkedstack"
)

func main() {
	stack := linkedstack.Stack{}
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)

	stack.String()
	fmt.Println()
	fmt.Println(stack.Top())
	stack.Pop()
	stack.String()
	fmt.Println()
	fmt.Println(stack.Top())
}
