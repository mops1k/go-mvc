package array

import "fmt"

type Collection struct {
	data []interface{}
}

func (a *Collection) Get(index int) (interface{}, error) {
	if !a.Exists(index) {
		return nil, fmt.Errorf(`index "%d" does not exists`, index)
	}

	return a.data[index], nil
}

func (a *Collection) Add(value interface{}) {
	a.data = append(a.data, value)
}

func (a *Collection) Update(index int, value interface{}) error {
	if !a.Exists(index) {
		return fmt.Errorf(`index "%d" does not exists`, index)
	}

	a.data[index] = value

	return nil
}

func (a *Collection) Remove(index int) error {
	if !a.Exists(index) {
		return fmt.Errorf(`index "%d" does not exists`, index)
	}

	a.data = append(a.data[:index], a.data[index+1:]...)

	return nil
}

func (a *Collection) List() []interface{} {
	return a.data
}

func (a *Collection) Exists(index int) bool {
	return !(index >= len(a.data) || index < 0)
}
