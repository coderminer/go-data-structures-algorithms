package main

import (
	"fmt"
	"go-data-structures-algorithms/linkedstack/linkedstack"
)

func main() {
	stack := linkedstack.Stack{}

	exp := "{[()]}"
	fmt.Println(stack.IsBalancedParenthesis(exp))

}
