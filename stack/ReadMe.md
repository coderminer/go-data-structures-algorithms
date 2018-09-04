### Go语言数据结构和算法-使用Slice实现栈

栈是Last-In-First-Out (LIFO)(后进先出)的数据结构，对应的接口如下：  

```
New()       //初始化栈
Push()      // 压栈
Pop()       // 出栈，返回栈顶的数据并删除栈顶的数据
Top()       // 返回栈顶的数据 不删除栈顶的数据
Size()      // 返回栈的大小
```

利用Go语言的Slice的特性实现栈的结构

```
type Stack struct {
	elements []interface{}
	lock     sync.RWMutex
}
```

#### New()

```
func (s *Stack) New() *Stack {
	s.elements = []interface{}{}
	return s
}
```

#### Push()

```
func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = append(s.elements, item)
}
```

#### Pop()

```
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	index := len(s.elements) - 1
	item := s.elements[index]
	s.elements = s.elements[0:index]
	return item
}
```

#### Top()

```
func (s *Stack) Top() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	index := len(s.elements) - 1
	return s.elements[index]
}
```

#### Size()

```
func (s *Stack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.elements)
}
```