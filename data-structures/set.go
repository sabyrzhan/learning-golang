package main

import "fmt"

type Set[T comparable] struct {
	elements map[T]struct{}
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

func (s *Set[T]) Contains(value T) bool {
	if _, ok := s.elements[value]; ok {
		return true
	}

	return false
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *Set[T]) Values() []T {
	result := make([]T, 0, len(s.elements))
	for v := range s.elements {
		result = append(result, v)
	}

	return result
}

func (s *Set[T]) Len() int {
	return len(s.elements)
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.Values())
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]struct{})}
}
