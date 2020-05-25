package _map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mcMap = map[interface{}]interface{}{
		"test 1": "value 1",
		"test 2": "value 2",
		"test 3": "value 3",
		"test 4": "value 4",
		"test 5": "value 5",
		"test 6": "value 6",
	}
	mapCollection = Collection{data: mcMap}
)

func TestArrayCollection_Add_List(t *testing.T) {
	err := mapCollection.Add("test 7", "value 7")
	assert.NoError(t, err)
	assert.Equal(t, map[interface{}]interface{}{
		"test 1": "value 1",
		"test 2": "value 2",
		"test 3": "value 3",
		"test 4": "value 4",
		"test 5": "value 5",
		"test 6": "value 6",
		"test 7": "value 7",
	}, mapCollection.List())
}

func TestArrayCollection_Exists(t *testing.T) {
	assert.False(t, mapCollection.Exists("test 432"))
	assert.True(t, mapCollection.Exists("test 1"))
}

func TestArrayCollection_Remove(t *testing.T) {
	err := mapCollection.Remove("test 2")
	assert.NoError(t, err)
	assert.Equal(t, map[interface{}]interface{}{
		"test 1": "value 1",
		"test 3": "value 3",
		"test 4": "value 4",
		"test 5": "value 5",
		"test 6": "value 6",
		"test 7": "value 7",
	}, mapCollection.List())
}

func TestArrayCollection_Update(t *testing.T) {
	err := mapCollection.Update("test 4", "new value")
	assert.NoError(t, err)
	val, err := mapCollection.Get("test 4")
	assert.NoError(t, err)
	assert.Equal(t, "new value", val)
}
