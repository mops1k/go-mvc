package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	acTest          = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	arrayCollection = Collection{collection: acTest}
)

func TestArrayCollection_Add_List(t *testing.T) {
	arrayCollection.Add(9)
	assert.Greater(t, len(arrayCollection.List()), len(acTest))
}

func TestArrayCollection_Exists(t *testing.T) {
	assert.False(t, arrayCollection.Exists(-1))
	assert.True(t, arrayCollection.Exists(5))
	assert.False(t, arrayCollection.Exists(11))
}

func TestArrayCollection_Remove(t *testing.T) {
	err := arrayCollection.Remove(2)
	assert.NoError(t, err)
	assert.NotEqual(t, acTest, arrayCollection.List())
	assert.Equal(t, []interface{}{1, 2, 4, 5, 6, 7, 8, 9}, arrayCollection.List())
}

func TestArrayCollection_Update(t *testing.T) {
	err := arrayCollection.Update(0, 9)
	assert.NoError(t, err)
	val, err := arrayCollection.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 9, val)
}
