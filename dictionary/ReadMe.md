### Go语言数据结构和算法-Dictionary(字典)

```
Set(key,value)      // 添加修改字典中的数据
Delete(key)         // 删除字典中key
Has(key)            // 是否包含key
Get(key)            // 根据key获取对应的value
Clear()             // 清空字典中的数据
Size()              // 获取字典的大小
Keys()              // 获取字典中的所有的key
Values()            // 获取字典中的所有的value
```

#### 字典的数据结构

```
type Dictionary struct {
	element map[interface{}]interface{}
	lock    sync.RWMutex
}
```

#### Set(key,value) 添加修改字典中的数据

```
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.element == nil {
		d.element = make(map[interface{}]interface{})
	}
	d.element[key] = value
}
```

#### Delete(key) 删除字典中key

```
func (d *Dictionary) Delete(key interface{}) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.element[key]
	if ok {
		delete(d.element, key)
	}
	return ok
}
```

#### Has(key) 是否包含key

```
func (d *Dictionary) Has(key interface{}) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.element[key]
	return ok
}
```

#### Get(key) 根据key获取对应的value

```
func (d *Dictionary) Get(key interface{}) interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.element[key]
}
```

#### Clear() 清空字典中的数据

```
func (d *Dictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.element = make(map[interface{}]interface{})
}
```

#### Size() 获取字典的大小

```
func (d *Dictionary) Size() int {
	d.lock.Lock()
	defer d.lock.Unlock()
	return len(d.element)
}
```

#### Keys() 获取字典中的所有的key

```
func (d *Dictionary) Keys() []interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	keys := make([]interface{}, 0)
	for key := range d.element {
		keys = append(keys, key)
	}
	return keys
}
```

#### Values() 获取字典中的所有的value

```
func (d *Dictionary) Values() []interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	values := make([]interface{}, 0)
	for key := range d.element {
		values = append(values, d.element[key])
	}
	return values
}
```