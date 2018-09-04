package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Append(val int) {
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
	list.size -= 1
}

func Print(head *Node) {
	cur := head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

/*
	合并两个有序的链表,两个递增的有序链表合并，
	合并之后的链表仍然是按照递增的顺序排序
*/
func Merge(head1 *Node, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var mergedHead *Node = nil
	if head1.val < head2.val {
		mergedHead = head1
		mergedHead.next = Merge(head1.next, head2)
	} else {
		mergedHead = head2
		mergedHead.next = Merge(head1, head2.next)
	}

	return mergedHead
}

func main() {
	list1 := LinkedList{}
	list2 := LinkedList{}

	list1.Append(2)
	list1.Append(4)
	list1.Append(6)
	list1.Append(8)
	list1.Append(10)

	list2.Append(1)
	list2.Append(3)
	list2.Append(5)
	list2.Append(7)

	Print(list1.head)
	fmt.Println()
	Print(list2.head)
	head := Merge(list1.head, list2.head)
	fmt.Println()
	Print(head)
}
