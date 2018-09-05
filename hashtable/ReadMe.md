### Go语言数据结构和算法-HashTable(哈希表)

```
Put(key,value)      // 在HashTable中添加数据
Remove(key)         // 从HashTable中删除数据
Get(key)            // 根据key获取HashTable中的数据
Size()              // 获取HashTable的大小
```

#### 哈希表的数据结构

```
type HashTable struct {
	element map[interface{}]interface{}
	lock    sync.RWMutex
}
```

#### 散列函数 

```
func hash(key interface{}) int {
	k := fmt.Sprintf("%s", key)
	h := 0
	for i := 0; i < len(k); i++ {
		h = 31*h + int(k[i])
	}
	fmt.Println(h)
	return h
}
```

#### Put(key,value) 在HashTable中添加数据

```
func (h *HashTable) Put(key interface{}, value interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.element == nil {
		h.element = make(map[interface{}]interface{})
	}
	i := hash(key)
	h.element[i] = value
}
```

#### Remove(key) 从HashTable中删除数据

```
func (h *HashTable) Remove(key interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	i := hash(key)
	delete(h.element, i)
}
```

#### Get(key) 根据key获取HashTable中的数据

```
func (h *HashTable) Get(key interface{}) interface{} {
	h.lock.Lock()
	defer h.lock.Unlock()
	i := hash(key)
	return h.element[i]
}
```

#### Size() 获取HashTable的大小

```
func (h *HashTable) Size() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.element)
}
```