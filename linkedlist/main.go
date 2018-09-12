package main

import (
	"fmt"
	"go-data-structures-algorithms/linkedlist/linkedlist"
)

func main() {
	list := linkedlist.LinkedList{}
	list.Prepend(2)
	list.Prepend(3)
	list.Prepend(4)
	list.Append(5)
	list.Append(6)

	list.String()
	fmt.Println()

	fmt.Println(list.IndexOf(7))
	fmt.Println(list.IsEmpty())
	fmt.Println(list.Size())
	fmt.Println(list.Head())
	fmt.Println(list.Tail())
	//list.Remove(6)
	list.String()
	fmt.Println()
	list.RemoveAt(5)
	list.String()
	fmt.Println()
	list.InsertAt(6, 7)
	list.String()
	list.Reverse()
	fmt.Println()
	list.String()
	fmt.Println()
	r, ok := list.NthNodeFromBegining(4)
	fmt.Println(r, ok)

	r, ok = list.NthNodeFromEnd(5)
	fmt.Println(r, ok)

	r, ok = list.NthNodeFromEnd2(5)
	fmt.Println(r, ok)
}
