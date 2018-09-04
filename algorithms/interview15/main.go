package main

import (
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Append(val interface{}) {
	node := &Node{val, nil}
	if list.head == nil {
		list.head = node
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
		list.size += 1
	}
}

func (list *LinkedList) Print() {
	cur := list.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

/*
	输出链表中倒数第k个节点,
	两个指针，一个指针先定位到k - 1处，
*/
func (list *LinkedList) FindKthToTail(k int) *Node {
	if list.head == nil || k == 0 {
		return nil
	}
	ahead := list.head
	for i := 0; i < k-1; i++ {
		if ahead.next != nil {
			ahead = ahead.next
		} else {
			return nil
		}
	}
	after := list.head
	for ahead.next != nil {
		ahead = ahead.next
		after = after.next
	}
	return after
}

func main() {
	list := LinkedList{}
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(7)

	fmt.Println(list.FindKthToTail(3))
}
