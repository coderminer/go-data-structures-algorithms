### Go语言数据结构和算法-Queue(队列)

```
New()           // 初始化队列
Enqueue(val)    // 入队
Dequeue()       // 出队
Front()         // 队列头节点
IsEmpty()       // 队列是否为空
Size()          // 队列的大小
```

#### 队列的数据结构

```
type Queue struct {
	elements []interface{}
	lock     sync.RWMutex
}
```

#### Enqueue(val) 入队

```
func (q *Queue) Enqueue(val interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.elements = append(q.elements, val)
}
```

#### Dequeue() 出队

```
func (q *Queue) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	item := q.elements[0]
	q.elements = q.elements[1:]
	return item
}

```

#### Front() 队列头节点

```
func (q *Queue) Front() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.elements[0]
}
```

#### IsEmpty() 队列是否为空 

```
func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.elements) == 0
}
```

#### Size() 队列的大小

```
func (q *Queue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.elements)
}
```