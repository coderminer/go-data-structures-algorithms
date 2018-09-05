package main

import(
	"fmt"
	"go-data-structures-algorithms/queue/queue"
)

func main(){
	q := queue.Queue{}
	q.New()

	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Front())
}