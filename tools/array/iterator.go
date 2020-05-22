package array

type Iterator struct {
	data    []interface{}
	current int
	next    interface{}
	count   int
}

func NewIterator(data []interface{}) *Iterator {
	iterator := &Iterator{data: data, next: false, count: len(data)}
	if iterator.count > 0 {
		iterator.current = 0
	}

	return iterator
}

func (a *Iterator) Current() interface{} {
	return a.data[a.current]
}

func (a *Iterator) Prev() bool {
	if a.current-1 < 0 {
		return false
	}

	a.current--
	return true
}

func (a *Iterator) Next() bool {
	if a.current+1 > a.count {
		return false
	}

	a.current++
	return true
}

func (a *Iterator) Count() int {
	return len(a.data)
}

func (a *Iterator) Key() int {
	return a.current
}
