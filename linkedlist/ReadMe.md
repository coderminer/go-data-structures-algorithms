### Go语言数据结构和算法-LinkedList(链表)

```
Prepend(item)       //  在链表头新增一个元素
Append(item)        //  在链表尾追加一个元素
InsertAt(i,item)    //  在索引i处插入一个元素
RemoveAt(i)         //  删除再索引i处的一个元素
Remove(item)        //  删除item元素
IndexOf(item)       //  返回item的索引
IsEmpty()           //  链表是否为空
Size()              //  链表的大小
Head()              //  链表的头节点
Tail()              //  链表的尾节点
String()            //  遍历链表
```

#### 节点和链表数据结构

```
type Node struct {
	item interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
  lock sync.RWMutex
}
```

#### Prepend(item) 在链表头新增一个元素

```
func (list *LinkedList) Prepend(item interface{}) {
  list.lock.Lock()
	defer list.lock.Unlock()
	node := &Node{item,nil}
	if list.head == nil{
		list.head = node
	}else{
		node.next = list.head
		list.head = node
	}
	list.size += 1
}
```

#### Append(item) 在链表尾追加一个元素

```
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
```

#### InsertAt(i,item) 在索引i处插入一个元素

```
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
```

#### RemoveAt(i) 删除再索引i处的一个元素

```
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
```

#### Remove(item) 删除item元素

```
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
```

#### IndexOf(item) 返回item的索引,如果不存在就返回-1

```
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
```

#### IsEmpty() 链表是否为空

```
func (list *LinkedList) IsEmpty() bool {
  list.lock.Lock()
	defer list.lock.Unlock()
	return list.size == 0
}
```

#### Size() 链表的大小

```
func (list *LinkedList) Size() int {
  list.lock.Lock()
	defer list.lock.Unlock()
	return list.size
}
```

#### Head() 链表头节点

```
func (list *LinkedList) Head() *Node {
  list.lock.Lock()
	defer list.lock.Unlock()
	return list.head
}
```

#### Tail() 链表尾节点

```
func (list *LinkedList) Tail() *Node {
  list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}
```

#### String() 遍历真个链表

```
func (list *LinkedList) String() {
  list.lock.Lock()
	defer list.lock.Unlock()
	cur := list.head
	for cur != nil {
		fmt.Print(cur.item, "->")
		cur = cur.next
	}
}
```
