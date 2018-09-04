package main

import (
	"fmt"
	"go-data-structures-algorithms/stack/stack"
)

func main() {
	s := stack.Stack{}
	s.New()
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	len := s.Size()
	for i := 0;i < len;i++{
		fmt.Println(s.Pop())
	}
}
