package array

type Iterator struct {
	Data    []interface{}
	current int
	next    interface{}
	count   int
}

func NewIterator(data []interface{}) *Iterator {
	iterator := &Iterator{Data: data}

	return iterator
}

func (a *Iterator) Current() interface{} {
	return a.Data[a.current]
}

func (a *Iterator) Prev() bool {
	if a.current-1 < 0 {
		return false
	}

	a.current--
	return true
}

func (a *Iterator) Next() bool {
	if a.current+1 > a.Count() {
		return false
	}

	a.current++
	return true
}

func (a *Iterator) Count() int {
	return len(a.Data)
}

func (a *Iterator) Key() int {
	return a.current
}
