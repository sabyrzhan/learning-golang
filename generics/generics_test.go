package generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	arrayList := NewArrayList[int]()
	arrayList.Add(1)
	arrayList.Add(2)
	arrayList.Add(3)

	// Test get success
	assert.Equal(t, 3, arrayList.Length())
	firstItem, _ := arrayList.Get(0)
	assert.Equal(t, 1, *firstItem)

	// Test get out of bounds index
	_, errorOutOfBounds := arrayList.Get(10)
	assert.NotNil(t, errorOutOfBounds)
	assert.Equal(t, "out of bounds index", errorOutOfBounds.Error())

	// Test remove success
	firstItem, _ = arrayList.Remove(0)
	assert.Equal(t, 1, *firstItem)

	// Test remove out of bounds index
	_, errorOutOfBounds = arrayList.Remove(10)
	assert.NotNil(t, errorOutOfBounds)
	assert.Equal(t, "out of bounds index", errorOutOfBounds.Error())

	// Test length
	assert.Equal(t, 2, arrayList.Length())

	// Test items
	values := arrayList.Items()
	assert.EqualValues(t, []int{2, 3}, values)
}

func TestGenericFunction(t *testing.T) {
	// Test with string
	stringResult, _ := BuildMap[string]("first", "1", "second", "2")
	assert.Equal(t, map[string]string{"first": "1", "second": "2"}, stringResult)

	// Test with integer
	intResult, _ := BuildMap[int](1, 2, 3, 4)
	assert.Equal(t, map[int]int{1: 2, 3: 4}, intResult)

	// Test with odd number of variables
	_, err := BuildMap[int](1, 2, 3)
	assert.NotNil(t, err)
	assert.Equal(t, "odd number of key/value pairs", err.Error())
}
