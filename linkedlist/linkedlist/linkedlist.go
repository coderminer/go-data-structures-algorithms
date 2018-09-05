package linkedlist

import (
	"fmt"
	"sync"
)

type Node struct {
	item interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (list *LinkedList) Prepend(item interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()
	node := &Node{item, nil}
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}
	list.size += 1
}

func (list *LinkedList) Append(item interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()
	node := &Node{item, nil}
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

func (list *LinkedList) InsertAt(i int, item interface{}) error {
	list.lock.Lock()
	defer list.lock.Unlock()
	if i > list.size || i < 0 {
		return fmt.Errorf("index error")
	}

	node := &Node{item, nil}
	if i == 0 {
		node.next = list.head
		list.head = node
		list.size += 1
		return nil
	}
	cur := list.head
	for cur.next != nil && i != 1 {
		i -= 1
		cur = cur.next
	}
	node.next = cur.next
	cur.next = node
	list.size += 1

	return nil
}

func (list *LinkedList) RemoveAt(i int) bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	if i < 0 || i >= list.size {
		return false
	}
	if list.head == nil {
		return false
	}

	if i == 0 {
		list.head = list.head.next
		list.size -= 1
		return true
	}

	cur := list.head
	for cur.next != nil && i != 1 {
		i -= 1
		cur = cur.next
	}
	cur.next = cur.next.next
	list.size -= 1
	return true
}

func (list *LinkedList) Remove(item interface{}) bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	if list.head == nil {
		return false
	}
	if list.head.item == item {
		list.head = list.head.next
		list.size -= 1
		return true
	}

	cur := list.head
	for cur.next != nil {
		if cur.next.item == item {
			cur.next = cur.next.next
			list.size -= 1
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *LinkedList) IndexOf(item interface{}) int {
	list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	index := -1
	for cur != nil {
		index += 1
		if item == cur.item {
			return index
		}
		cur = cur.next
	}
	return -1
}

func (list *LinkedList) IsEmpty() bool {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.size == 0
}

func (list *LinkedList) Size() int {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.size
}

func (list *LinkedList) Head() *Node {
	list.lock.Lock()
	defer list.lock.Unlock()
	return list.head
}

func (list *LinkedList) Tail() *Node {
	list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

func (list *LinkedList) String() {
	list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	for cur != nil {
		fmt.Print(cur.item, "->")
		cur = cur.next
	}
}
