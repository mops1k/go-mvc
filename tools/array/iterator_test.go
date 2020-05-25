package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testArray = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i         = NewIterator(testArray)
)

func TestArrayIterator_Count(t *testing.T) {
	assert.Equal(t, i.Count(), len(testArray))
}

func TestArrayIterator(t *testing.T) {
	for a := 0; a < 7; a++ {
		i.Next()
	}

	assert.Equal(t, i.Key(), 7)
	assert.Equal(t, i.Current(), testArray[i.Key()])
	i.Prev()
	assert.Equal(t, i.Key(), 6)
	assert.NotEqual(t, i.Key(), 7)

	for i.Next() == true {
	}
	assert.False(t, i.Next())
	assert.True(t, i.Prev())

	for i.Prev() == true {
	}
	assert.False(t, i.Prev())
	assert.True(t, i.Next())
}
