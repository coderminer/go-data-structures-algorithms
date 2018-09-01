package main

import (
	"fmt"
	"sync"
)

type Node struct {
	item interface{}
	next *Node
}

type SingleList struct {
	head *Node
	tail *Node
	size int
	lock sync.RWMutex
}

func (list *SingleList) Append(item interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()

	node := &Node{item, nil}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.size++
}

func (list *SingleList) Size() int {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.size
}

func (list *SingleList) IsEmpty() bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.size == 0
}

func (list *SingleList) IndexOf(item interface{}) int {
	list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	index := 0
	for cur != nil {
		if item == cur.item {
			return index
		}
		cur = cur.next
		index++
	}
	return -1
}

func (list *SingleList) Display() {
	list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	for cur != nil {
		fmt.Println(cur.item)
		cur = cur.next
	}
}

func (list *SingleList) GetAt(index int) *Node {
	list.lock.Lock()
	defer list.lock.Unlock()
	if index > list.size || index < 0 {
		return nil
	}
	cur := list.head
	i := 0
	for i != index {
		i++
		cur = cur.next
	}
	return cur
}

func (list *SingleList) RemoveAt(index int) bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	if index >= list.size || index < 0 {
		return false
	}
	if index == 0 {
		list.head = list.head.next
		list.size--
		return true
	}

	cur := list.head
	for index != 1 {
		cur = cur.next
		index--
	}

	cur.next = cur.next.next
	list.size--
	return true
}

func (list *SingleList) InsertAt(index int, item interface{}) bool {
	if index < 0 || index > list.size {
		return false
	}
	node := &Node{item, nil}
	if index == 0 {
		node.next = list.head
		list.head = node
		list.size++
		return true
	}
	s := list.size - 1
	if index == s {
		list.tail.next = node
		list.tail = node
		list.size++
		return true
	}
	cur := list.head
	for index != 1 {
		cur = cur.next
		index--
	}

	node.next = cur.next
	cur.next = node
	list.size++

	return true
}

func (list *SingleList) Contains(item interface{}) bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	for cur != nil {
		if cur.item == item {
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *SingleList) Remove(item interface{}) bool {
	list.lock.Lock()
	defer list.lock.Unlock()

	if list.head == nil {
		return false
	}

	if list.head.item == item {
		list.head = list.head.next
		list.size--
		return true
	}

	cur := list.head
	for cur.next != nil {
		if cur.next.item == item {
			cur.next = cur.next.next
			return true
		}
		cur = cur.next
	}

	return false
}

func (list *SingleList) String() {
	list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	for cur != nil {
		fmt.Print(cur.item, " ")
		cur = cur.next
	}
}

func main() {
	list := SingleList{}

	list.Append(2) //0
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)

	list.Remove(6)
	list.String()
}
