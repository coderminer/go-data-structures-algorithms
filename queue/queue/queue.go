package queue

import (
	"sync"
)

type Queue struct {
	elements []interface{}
	lock     sync.RWMutex
}

func (q *Queue) New() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.elements = make([]interface{}, 0)
}

func (q *Queue) Enqueue(val interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.elements = append(q.elements, val)
}

func (q *Queue) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	item := q.elements[0]
	q.elements = q.elements[1:]
	return item
}

func (q *Queue) Front() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.elements)
}
