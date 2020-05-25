package _map

import (
	"fmt"
)

// Iterable collection struct
type Collection struct {
	data     map[interface{}]interface{}
	iterator *Iterator
}

// check if value is exists
func (m *Collection) Exists(key interface{}) bool {
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
		return false
	}

	if _, exists := m.data[key]; !exists {
		return false
	}

	return true
}

// Get value
func (m *Collection) Get(key interface{}) (interface{}, error) {
	if !m.Exists(key) || m.data == nil {
		return nil, fmt.Errorf(`key "%s" does not exists`, key)
	}

	return m.data[key], nil
}

// Add value
func (m *Collection) Add(key interface{}, value interface{}) error {
	if m.Exists(key) {
		return fmt.Errorf(`key "%s" already exists`, key)
	}

	m.data[key] = value

	return nil
}

// Update value
func (m *Collection) Update(key interface{}, value interface{}) error {
	if !m.Exists(key) {
		return fmt.Errorf(`key "%s" does not exists`, key)
	}

	m.data[key] = value

	return nil
}

// remove value
func (m *Collection) Remove(key interface{}) error {
	if !m.Exists(key) {
		return fmt.Errorf(`key "%s" does not exists`, key)
	}

	delete(m.data, key)

	return nil
}

// list all values
func (m *Collection) List() map[interface{}]interface{} {
	return m.data
}

func (m *Collection) CreateIterator() *Iterator {
	return NewIterator(m.List())
}
