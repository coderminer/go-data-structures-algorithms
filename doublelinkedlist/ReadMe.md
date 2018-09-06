### Go语言数据结构和算法-DoubleLinkedList(双向链表)

```
Prepend(val)    // 在双向链表的头部添加新数据
Append(val)     // 在双向链表的尾部添加新数据
Remove(val)     // 在双向链表中删除一个数据
Contains(val)   // 在双向链表中是否包含这个元素
Reverse()		// 倒序遍历双向链表
String()        // 遍历打印双向链表的所有元素
```

#### 双向链表的数据结构

```
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
```

#### Prepend(val) 在双向链表的头部添加新数据

```
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
```

#### Append(val) 在双向链表的尾部添加新数据

```
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
```

#### Remove(val) 在双向链表中删除一个数据

```
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
```

#### Contains(val) 在双向链表中是否包含这个元素

```
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
```

#### Reverse() 倒序遍历双向链表

```
func (list *DoubleLinkedList) Reverse() {
	cur := list.tail
	for cur != nil {
		fmt.Print(cur.val, "->")
		cur = cur.prev
	}
}
```

#### String() 遍历打印双向链表的所有元素

```
func (list *DoubleLinkedList) String() {
	cur := list.head
	for cur != nil {
		fmt.Print(cur.val, "->")
		cur = cur.next
	}
}
```
