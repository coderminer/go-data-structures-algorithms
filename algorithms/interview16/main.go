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
	}
	list.size += 1
}

func (list *LinkedList) Print(head *Node) {
	cur := head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

/*
	反转链表，并返回链表的头节点
*/
func (list *LinkedList) Reverse() *Node {
	var reversedHead *Node = nil
	node := list.head
	var prev *Node = nil
	for node != nil {
		next := node.next
		if next == nil {
			reversedHead = node
		}
		node.next = prev
		prev = node
		node = next
	}
	return reversedHead

}

func main() {
	list := LinkedList{}
	list.Append(3)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(8)

	//list.Print(list.head)
	head := list.Reverse()
	list.Print(head)
}
