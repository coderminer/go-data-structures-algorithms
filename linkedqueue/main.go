package main

import (
	"fmt"
	"go-data-structures-algorithms/linkedqueue/linkedqueue"
)

func main() {
	q := linkedqueue.Queue{}
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(8)

	q.Dequeue()
	q.Dequeue()
	q.String()
	fmt.Println()
	fmt.Println(q.Size())
}
