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

func (list *LinkedList) Print() {
	cur := list.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func (list *LinkedList) GetAt(index int) *Node {
	cur := list.head
	count := 0
	for cur != nil && count != index {
		count += 1
		cur = cur.next
	}
	return cur
}

/*
	在O(1)时间内删除一个链表的节点,这个节点必须是在这个链表中，否则无法在
	O(1)时间内删除一个节点
	需要通过循环查找需要删除的节点
*/
func (list *LinkedList) Delete(toBeDelete *Node) bool {

	if list.head == nil || toBeDelete == nil {
		return false
	}
	//要删除的节点不是尾节点
	if toBeDelete.next != nil {
		mNext := toBeDelete.next
		toBeDelete.val = mNext.val
		toBeDelete.next = mNext.next

		mNext = nil
		list.size -= 1
	} else if list.head == toBeDelete {
		list.head = nil
		list.size -= 1
	} else {
		node := list.head
		for node.next != toBeDelete {
			node = node.next
		}
		node.next = nil
		toBeDelete = nil
		list.size -= 1
	}

	return true
}

func main() {
	list := LinkedList{}
	list.Append(4)
	list.Append(6)
	list.Append(7)
	list.Append(8)
	list.Append(10)
	list.Append(3)

	node := list.GetAt(5)
	list.Delete(node)

	list.Print()
}
