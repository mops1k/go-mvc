package array

import "fmt"

type Collection struct {
	collection []interface{}
}

func (a *Collection) Get(index int) (interface{}, error) {
	if !a.Exists(index) {
		return nil, fmt.Errorf(`index "%d" does not exists`, index)
	}

	return a.collection[index], nil
}

func (a *Collection) Add(value interface{}) {
	a.collection = append(a.collection, value)
}

func (a *Collection) Update(index int, value interface{}) error {
	if !a.Exists(index) {
		return fmt.Errorf(`index "%d" does not exists`, index)
	}

	a.collection[index] = value

	return nil
}

func (a *Collection) Remove(index int) error {
	if !a.Exists(index) {
		return fmt.Errorf(`index "%d" does not exists`, index)
	}

	a.collection = append(a.collection[:index], a.collection[index+1:]...)

	return nil
}

func (a *Collection) List() []interface{} {
	return a.collection
}

func (a *Collection) Exists(index int) bool {
	return !(index >= len(a.collection) || index < 0)
}
