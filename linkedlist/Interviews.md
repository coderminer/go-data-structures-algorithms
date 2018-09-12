### 关于单向链表相关的问题-Go语言实现

#### 单向链表反转 

```
func (list *LinkedList) Reverse() {
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
```

#### 两个单向链表比较

```
func (list *LinkedList) CompareList(ll *LinkedList) bool {
	return compareListUtils(list.head, ll.head)
}

func compareListUtils(head1 *Node, head2 *Node) bool {
	if head1 == nil && head2 == nil {
		return true
	} else if head1 == nil || head2 == nil || head1.item != head2.item {
		return false
	} else {
		return compareListUtils(head1.next, head2.next)
	}

}
```

#### 正数第n个元素

```
func (list *LinkedList) NthNodeFromBegining(index int) (interface{}, bool) {
	if index > list.Size() || index < 1 {
		fmt.Println("TooFewNodes")
		return nil, false
	}
	cur := list.head
	for cur != nil && index != 1 {
		index -= 1
		cur = cur.next
	}
	return cur.item, true
}
```

#### 倒数第n个元素

根据链表的长度，计算出正数的位置 `startIndex = size() - index + 1`,然后调用`NthNodeFromBegining`  

```
func (list *LinkedList) NthNodeFromEnd(index int) (interface{}, bool) {
	if index > list.Size() || index < 1 {
		fmt.Println("TooFewNodes")
		return nil, false
	}
	startIndex := list.Size() - index + 1
	return list.NthNodeFromBegining(startIndex)
}
```

另一种解法,利用两个指针，快的指针先到index - 1处，然后慢指针开始从头遍历直到遍历到尾节点，此时慢节点就是要找的节点
利用这种方法还可以检测一个单向链表是否是循环链表或是否包含环路

```
func (list *LinkedList) NthNodeFromEnd2(index int) (interface{}, bool) {
	if list.head == nil || index < 1 {
		return nil, false
	}
	fast := list.head
	slow := list.head

	for fast.next != nil && index != 1 {
		index -= 1
		fast = fast.next
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow.item, true
}
```

#### 判断一个单向链表中是否有环路

```
func (list *LinkedList) LoopDetect() (*Node, bool) {
	if list.head == nil {
		return nil, false
	}
	fast := list.head
	slow := list.head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return slow, true
		}
	}
	return nil, false
}
```