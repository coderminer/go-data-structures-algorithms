package dictionary

import (
	"sync"
)

type Dictionary struct {
	element map[interface{}]interface{}
	lock    sync.RWMutex
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.element == nil {
		d.element = make(map[interface{}]interface{})
	}
	d.element[key] = value
}

func (d *Dictionary) Delete(key interface{}) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.element[key]
	if ok {
		delete(d.element, key)
	}
	return ok
}

func (d *Dictionary) Has(key interface{}) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.element[key]
	return ok
}

func (d *Dictionary) Get(key interface{}) interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.element[key]
}

func (d *Dictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.element = make(map[interface{}]interface{})
}

func (d *Dictionary) Size() int {
	d.lock.Lock()
	defer d.lock.Unlock()
	return len(d.element)
}

func (d *Dictionary) Keys() []interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	keys := make([]interface{}, 0)
	for key := range d.element {
		keys = append(keys, key)
	}
	return keys
}

func (d *Dictionary) Values() []interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	values := make([]interface{}, 0)
	for key := range d.element {
		values = append(values, d.element[key])
	}
	return values
}
