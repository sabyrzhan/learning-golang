package generics

import "fmt"

type List[T comparable] interface {
	Add(item T) T
	Get(i int) (*T, error)
	IsEmpty() bool
	Length() int
	Remove(i int) (*T, error)
	Items() []T
}

type ArrayList[T comparable] struct {
	items []T
}

func (a *ArrayList[T]) Add(item T) T {
	a.items = append(a.items, item)
	return item
}

func (a *ArrayList[T]) Get(i int) (*T, error) {
	if a.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if len(a.items) <= i {
		return nil, fmt.Errorf("out of bounds index")
	}

	return &a.items[0], nil
}

func (a *ArrayList[T]) Length() int {
	return len(a.items)
}

func (a *ArrayList[T]) IsEmpty() bool {
	return a.Length() == 0
}

func (a *ArrayList[T]) Remove(i int) (*T, error) {
	if a.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if len(a.items) <= i {
		return nil, fmt.Errorf("out of bounds index")
	}

	v := a.items[i]
	a.items = a.items[1:]

	return &v, nil
}

func (a *ArrayList[T]) Items() []T {
	return a.items
}

func NewArrayList[T comparable]() ArrayList[T] {
	return ArrayList[T]{}
}

func BuildMap[KV comparable](kv ...KV) (map[KV]KV, error) {
	result := map[KV]KV{}
	if len(kv)%2 != 0 {
		return nil, fmt.Errorf("odd number of key/value pairs")
	}

	for i := 0; i < len(kv); i += 2 {
		result[kv[i]] = kv[i+1]
	}

	return result, nil
}
