package main

import(
	"fmt"
	"go-data-structures-algorithms/heap/heap"
)


func main(){
	heap := heap.Heap{}
	heap.Add(13)
	heap.Add(10)
	heap.Add(14)
	heap.Add(9)
	heap.Add(15)
	heap.Add(8)

	heap.String()
	fmt.Println(heap.Contains(16))
	heap.Delete(14)
	fmt.Println()
	heap.String()
}