package main

import (
	"fmt"
)

type Queue struct {
	stack1 []int
	stack2 []int
}

/*
	利用两个栈(slice)实现队列
*/
func (q *Queue) Append(val int) {
	q.stack1 = append(q.stack1, val)
}

func (q *Queue) DeleteHead() int {
	if len(q.stack2) <= 0 {
		for len(q.stack1) > 0 {
			index := len(q.stack1) - 1
			data := q.stack1[index]
			q.stack1 = append(q.stack1[:index], q.stack1[index+1:]...)
			q.stack2 = append(q.stack2, data)
		}
	}

	index := len(q.stack2) - 1
	head := q.stack2[index]
	q.stack2 = append(q.stack2[:index], q.stack2[index+1:]...)

	return head

}

func main() {
	q := Queue{}
	q.Append(2)
	q.Append(4)
	q.Append(5)

	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	q.Append(7)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())


}
