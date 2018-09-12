package linkedqueue

import (
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

type Queue struct {
	head *Node
	size int
}

func (q *Queue) Enqueue(val interface{}) {
	node := &Node{val, nil}
	if q.head == nil {
		q.head = node
	} else {
		cur := q.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	q.size += 1
}

func (q *Queue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	} else {
		h := q.head
		q.head = q.head.next
		q.size -= 1
		return h.val
	}
}

func (q *Queue) Peek() interface{} {
	if q.head == nil {
		return nil
	} else {
		return q.head.val
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) String() {
	cur := q.head
	for cur != nil {
		fmt.Print(cur.val, "->")
		cur = cur.next
	}
}
