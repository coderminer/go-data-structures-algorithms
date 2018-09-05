package hashtable

import (
	"fmt"
	"sync"
)

type HashTable struct {
	element map[interface{}]interface{}
	lock    sync.RWMutex
}

func hash(key interface{}) int {
	k := fmt.Sprintf("%s", key)
	h := 0
	for i := 0; i < len(k); i++ {
		h = 31*h + int(k[i])
	}
	fmt.Println(h)
	return h
}

func (h *HashTable) Put(key interface{}, value interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.element == nil {
		h.element = make(map[interface{}]interface{})
	}
	i := hash(key)
	h.element[i] = value
}

func (h *HashTable) Remove(key interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	i := hash(key)
	delete(h.element, i)
}

func (h *HashTable) Get(key interface{}) interface{} {
	h.lock.Lock()
	defer h.lock.Unlock()
	i := hash(key)
	return h.element[i]
}

func (h *HashTable) Size() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.element)
}
