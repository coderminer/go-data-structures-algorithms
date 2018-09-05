### Go语言数据结构和算法-Queue(队列)

```
New()           // 初始化队列
Enqueue()       // 入队
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