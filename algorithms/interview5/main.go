package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

/*
	从链表的尾部追加元素
*/

func (list *LinkedList) Append(val int) {
	node := &Node{val, nil}
	if list.head == nil {
		list.head = node
		list.size += 1
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
		list.size += 1
	}
}

/*
	从尾到头反过来打印出每个节点的值，借助slice打印
*/

func (list *LinkedList) ReverseWithSlice() {
	arr := make([]*Node, 0)
	cur := list.head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.next
	}
	for len(arr) != 0 {
		index := len(arr) - 1
		fmt.Println(arr[index].val)
		arr = append(arr[:index], arr[index+1:]...)
	}
}

/*
	从尾到头反过来打印出每个节点的值,使用递归的方式实现
*/

func (list *LinkedList) ReverseWithRecurse(head *Node) {
	if head != nil {
		if head.next != nil {
			list.ReverseWithRecurse(head.next)
		}
		fmt.Println(head.val)
	}
}

func main() {

	list := LinkedList{}
	list.Append(2)
	list.Append(4)
	list.Append(6)
	list.Append(7)

	list.ReverseWithRecurse(list.head)
}
