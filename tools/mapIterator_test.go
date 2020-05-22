package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMap = map[interface{}]interface{}{
		"test 1": "value 1",
		"test 2": "value 2",
		"test 3": "value 3",
		"test 4": "value 4",
		"test 5": "value 5",
		"test 6": "value 6",
	}
	mapIterator = NewMapIterator(testMap)
)

func TestMapIterator_Count(t *testing.T) {
	assert.Equal(t, mapIterator.count, len(testMap))
}

func TestMapIterator(t *testing.T) {
	for a := 0; a < 4; a++ {
		mapIterator.Next()
	}

	assert.Equal(t, mapIterator.Current(), testMap[mapIterator.Key()])

	for mapIterator.Next() == true {
	}
	assert.False(t, mapIterator.Next())
	assert.True(t, mapIterator.Prev())

	for mapIterator.Prev() == true {
	}
	assert.False(t, mapIterator.Prev())
	assert.True(t, mapIterator.Next())
}
