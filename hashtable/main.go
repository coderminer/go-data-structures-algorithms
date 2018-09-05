package main

import (
	"fmt"
	"go-data-structures-algorithms/hashtable/hashtable"
)

func main() {
	h := hashtable.HashTable{}
	h.Put("first","the first value")
	h.Put("second","the second value")
	h.Put("third","the third value")
	h.Put("forth","the forth value")

	fmt.Println(h.Get("second"))
	fmt.Println(h.Size())
	h.Remove("third")
	
}
