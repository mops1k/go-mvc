package _map

import (
	"reflect"
)

type Iterator struct {
	data    map[interface{}]interface{}
	current int
	keys    []reflect.Value
	next    interface{}
	count   int
}

func NewIterator(data map[interface{}]interface{}) *Iterator {
	iterator := &Iterator{data: data, next: false, count: len(data)}
	if iterator.count > 0 {
		iterator.current = 0
	}

	return iterator
}

func (m *Iterator) Current() interface{} {
	if m.keys == nil {
		m.readKeys()
	}

	return m.data[m.Key()]
}

func (m *Iterator) Prev() bool {
	if m.keys == nil {
		m.readKeys()
	}

	if m.current-1 < 0 {
		return false
	}

	m.current--
	return true
}

func (m *Iterator) Next() bool {
	if m.keys == nil {
		m.readKeys()
	}

	if m.current+1 > m.count {
		return false
	}

	m.current++
	return true
}

func (m *Iterator) Count() int {
	return len(m.data)
}

func (m *Iterator) Key() interface{} {
	key := m.keys[m.current]

	return key.Interface()
}

func (m *Iterator) readKeys() {
	m.keys = reflect.ValueOf(m.data).MapKeys()
}
