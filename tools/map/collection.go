package _map

import (
	"fmt"
)

type Collection struct {
	collection map[interface{}]interface{}
}

func (m *Collection) Exists(key interface{}) bool {
	if _, exists := m.collection[key]; !exists {
		return false
	}

	return true
}

func (m *Collection) Get(key interface{}) (interface{}, error) {
	if !m.Exists(key) {
		return nil, fmt.Errorf(`key "%s" does not exists`, key)
	}

	return m.collection[key], nil
}

func (m *Collection) Add(key interface{}, value interface{}) error {
	if m.Exists(key) {
		return fmt.Errorf(`key "%s" already exists`, key)
	}

	m.collection[key] = value

	return nil
}

func (m *Collection) Update(key interface{}, value interface{}) error {
	if !m.Exists(key) {
		return fmt.Errorf(`key "%s" does not exists`, key)
	}

	m.collection[key] = value

	return nil
}

func (m *Collection) Remove(key interface{}) error {
	if !m.Exists(key) {
		return fmt.Errorf(`key "%s" does not exists`, key)
	}

	delete(m.collection, key)

	return nil
}

func (m *Collection) List() map[interface{}]interface{} {
	return m.collection
}
