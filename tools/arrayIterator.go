package tools

type ArrayIterator struct {
	data    []interface{}
	current int
	next    interface{}
	count   int
}

func NewArrayIterator(data []interface{}) *ArrayIterator {
	iterator := &ArrayIterator{data: data, next: false, count: len(data)}
	if iterator.count > 0 {
		iterator.current = 0
	}

	return iterator
}

func (a *ArrayIterator) Current() interface{} {
	return a.data[a.current]
}

func (a *ArrayIterator) Prev() bool {
	if a.current-1 < 0 {
		return false
	}

	a.current--
	return true
}

func (a *ArrayIterator) Next() bool {
	if a.current+1 > a.count {
		return false
	}

	a.current++
	return true
}

func (a *ArrayIterator) Count() int {
	return len(a.data)
}

func (a *ArrayIterator) Key() int {
	return a.current
}
