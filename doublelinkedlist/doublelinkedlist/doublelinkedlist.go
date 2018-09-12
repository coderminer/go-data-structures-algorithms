package doublelinkedlist

import (
	"fmt"
)

type Node struct {
	val  interface{}
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (list *DoubleLinkedList) Prepend(val interface{}) {
	node := &Node{val, nil, nil}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}
	list.size += 1
}

func (list *DoubleLinkedList) Append(val interface{}) {
	node := &Node{val, nil, nil}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		node.prev = list.tail
		list.tail = node
	}
	list.size += 1
}

func (list *DoubleLinkedList) Remove(val interface{}) bool {
	if list.head == nil {
		return false
	}
	if list.head.val == val {
		if list.head == list.tail {
			list.head = nil
			list.tail = nil
		} else {
			list.head = list.head.next
			list.head.prev = nil
		}
		return true
	}
	cur := list.head.next
	for cur != nil {
		if cur.val == val {
			if cur == list.tail {
				list.tail = list.tail.prev
				list.tail.next = nil
			} else {
				cur.prev.next = cur.next
				cur.next.prev = cur.prev
			}
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *DoubleLinkedList) Contains(val interface{}) bool {
	cur := list.head
	for cur != nil {
		if cur.val == val {
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *DoubleLinkedList) Reverse() {
	cur := list.tail
	for cur != nil {
		fmt.Print(cur.val, "->")
		cur = cur.prev
	}
}

func (list *DoubleLinkedList) String() {
	cur := list.head
	for cur != nil {
		fmt.Print(cur.val, "->")
		cur = cur.next
	}
}

func (list *DoubleLinkedList) ReverseList() {
	cur := list.head
	for cur != nil {
		cur.next, cur.prev = cur.prev, cur.next
		if cur.prev == nil {
			list.tail = list.head
			list.head = cur
			return
		}
		cur = cur.prev
	}
}

func (list *DoubleLinkedList) CopyListReversed(dll *DoubleLinkedList) {
	cur := list.head
	for cur != nil {
		dll.Prepend(cur.val)
		cur = cur.next
	}
}

func (list *DoubleLinkedList) CopyList(dll *DoubleLinkedList) {
	cur := list.head
	for cur != nil {
		dll.Append(cur.val)
		cur = cur.next
	}
}
