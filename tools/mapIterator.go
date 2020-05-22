package tools

import (
	"reflect"
)

type MapIterator struct {
	data    map[interface{}]interface{}
	current int
	keys    []reflect.Value
	next    interface{}
	count   int
}

func NewMapIterator(data map[interface{}]interface{}) *MapIterator {
	iterator := &MapIterator{data: data, next: false, count: len(data)}
	if iterator.count > 0 {
		iterator.current = 0
	}

	return iterator
}

func (m *MapIterator) Current() interface{} {
	if m.keys == nil {
		m.readKeys()
	}

	return m.data[m.Key()]
}

func (m *MapIterator) Prev() bool {
	if m.keys == nil {
		m.readKeys()
	}

	if m.current-1 < 0 {
		return false
	}

	m.current--
	return true
}

func (m *MapIterator) Next() bool {
	if m.keys == nil {
		m.readKeys()
	}

	if m.current+1 > m.count {
		return false
	}

	m.current++
	return true
}

func (m *MapIterator) Count() int {
	return len(m.data)
}

func (m *MapIterator) Key() interface{} {
	key := m.keys[m.current]

	return key.Interface()
}

func (m *MapIterator) readKeys() {
	m.keys = reflect.ValueOf(m.data).MapKeys()
}
