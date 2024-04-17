package main

import (
	"errors"
	"fmt"
)

type Queue[T comparable] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	var result T
	result, q.elements = q.elements[0], q.elements[1:]

	return &result, nil
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) Clear() {
	q.elements = []T{}
}

func (q *Queue[T]) Values() []T {
	return q.elements
}

func (q *Queue[T]) Peek() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return &q.elements[0], nil
}

func (q *Queue[T]) String() string {
	return fmt.Sprint(q.elements)
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{[]T{}}
}
