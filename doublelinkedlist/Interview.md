### 关于双向链表相关的问题-Go语言实现

#### 反转一个双向链表

```
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
```

#### 把一个双向链表反向Copy到另一个双向链表

遍历链表，并把值按照在头部插入的方式插入到另一个链表  

```
func (list *DoubleLinkedList) CopyListReversed(dll *DoubleLinkedList) {
	cur := list.head
	for cur != nil {
		dll.Prepend(cur.val)
		cur = cur.next
	}
}
```

#### 把一个双向链表Copy到另一个双向链表

```
func (list *DoubleLinkedList) CopyList(dll *DoubleLinkedList) {
	cur := list.head
	for cur != nil {
		dll.Append(cur.val)
		cur = cur.next
	}
}
```