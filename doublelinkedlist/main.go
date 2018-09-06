package main

import (
	"fmt"
	"go-data-structures-algorithms/doublelinkedlist/doublelinkedlist"
)

func main() {
	list := doublelinkedlist.DoubleLinkedList{}
	list.Prepend(4)
	list.Prepend(5)
	list.Prepend(6)
	list.Prepend(7)

	list.Append(8)
	list.Append(9)
	list.Append(10)

	list.String()
	fmt.Println()
	fmt.Println(list.Contains(8))
	list.Remove(5)
	list.String()
	fmt.Println()
	list.Reverse()
}
