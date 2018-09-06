package main

import (
	"fmt"
	"sync"
)

type Node struct {
	val  int
	next *Node
}

type SingleLinkedList struct {
	head *Node
	size int
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (list *SingleLinkedList) Size() int {
	return list.size
}

func (list *SingleLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *SingleLinkedList) AddHead(val int) {
	list.head = &Node{val, list.head}
	list.size++
}

func (list *SingleLinkedList) AddTail(val int) {
	node := &Node{val, nil}
	if list.head == nil {
		list.head = node
		list.size++
		return
	}
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
	list.size++
}

func (list *SingleLinkedList) IndexOf(val int) int {
	index := -1
	cur := list.head
	for cur != nil {
		index++
		if cur.val == val {
			return index
		}
		cur = cur.next
	}
	return index
}

func (list *SingleLinkedList) InsertAt(index int, val int) {
	if index > list.size || index < 0 {
		return
	}
	if index == 0 {
		list.AddHead(val)
		list.size++
		return
	}
	count := 1
	cur := list.head
	for cur != nil && count < index {
		count++
		cur = cur.next
	}

	node := &Node{val, cur.next}
	cur.next = node
	list.size++
	//fmt.Println(cur.val)
}

func (list *SingleLinkedList) SortedInsert(val int) {
	node := &Node{val, nil}
	cur := list.head
	if cur == nil || cur.val > val {
		node.next = list.head
		list.head = node
		return
	}

	for cur.next != nil && cur.next.val < val {
		cur = cur.next
	}
	node.next = cur.next
	cur.next = node
}

func (list *SingleLinkedList) Contains(val int) bool {
	cur := list.head
	for cur != nil {
		if cur.val == val {
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *SingleLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}
	value := list.head.val
	list.head = list.head.next
	list.size--
	return value, true
}

func (list *SingleLinkedList) DeleteNode(val int) bool {
	if list.IsEmpty() {
		return false
	}
	if val == list.head.val {
		list.head = list.head.next
		list.size--
		return true
	}
	cur := list.head
	for cur.next != nil {
		if cur.next.val == val {
			cur.next = cur.next.next
			list.size--
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *SingleLinkedList) RemoveDuplicate() {
	cur := list.head
	for cur != nil {
		if cur.next != nil && cur.val == cur.next.val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
}

func (list *SingleLinkedList) CompareList(l *SingleLinkedList) bool {
	return list.compareListUtil(list.head, l.head)
}

func (list *SingleLinkedList) compareListUtil(head1 *Node, head2 *Node) bool {
	if head1 == nil && head2 == nil {
		return true
	} else if (head1 == nil) || (head2 == nil) || (head1.val != head2.val) {
		return false
	} else {
		return list.compareListUtil(head1.next, head2.next)
	}
}

func (list *SingleLinkedList) FindLength() int {
	cur := list.head
	count := 0
	for cur != nil {
		count++
		cur = cur.next
	}
	return count
}

func (list *SingleLinkedList) NthNodeFromBeginning(index int) (int, bool) {
	if index > list.Size() || index < 1 {
		return 0, false
	}
	count := 0
	cur := list.head
	for cur != nil && count < index-1 {
		count++
		cur = cur.next
	}
	return cur.val, true
}

func (list *SingleLinkedList) FreeList() {
	list.head = nil
	list.size = 0
}

func (list *SingleLinkedList) Display() {
	cur := list.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func (list *SingleLinkedList) Reverse() {
	cur := list.head
	var prev, next *Node
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	list.head = prev
}

func (list *SingleLinkedList) LoopDetect() bool {
	slow := list.head
	fast := list.head
 
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func (list *LinkedList) Append(val int) {
	list.lock.Lock()
	defer list.lock.Unlock()
	newNode := &Node{val, nil}
	if list.head == nil {
		list.head = newNode
	} else {

	}
}

func main() {
	list := SingleLinkedList{}

	list.SortedInsert(12)
	list.SortedInsert(5)
	list.SortedInsert(8)
	list.SortedInsert(23)
	list.SortedInsert(34)
	list.SortedInsert(3)
	list.SortedInsert(7)
	list.SortedInsert(5)

	r, err := list.NthNodeFromBeginning(2)
	fmt.Println(r, err)
}
