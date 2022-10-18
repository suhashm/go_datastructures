package queue

import (
	"errors"
	"fmt"
	"sync"
)

type queue[T any] struct {
	sync.RWMutex
	defaultValue T
	front        int
	back         int
	values       []T
}

func New[T any](defaultValue T) *queue[T] {
	return NewWithCapacity(defaultValue, 10000)
}

func NewWithCapacity[T any](defaultValue T, capacity int) *queue[T] {
	return &queue[T]{
		defaultValue: defaultValue,
		values:       make([]T, capacity),
	}
}

func (q *queue[T]) Add(value T) {

	q.values[q.back] = value
	q.back++
}

func (q *queue[T]) Remove() (T, error) {
	if q.Size() <= 0 {
		return q.defaultValue, errors.New("queue empty")
	}

	result := q.values[q.front]
	q.front++

	if q.Size() == 0 {
		q.front = 0
		q.back = 0
	}

	return result, nil
}

func (q *queue[T]) Peek() (T, error) {
	if len(q.values) == 0 {
		return q.defaultValue, errors.New("queue empty")
	}

	return q.values[q.front], nil
}

func (q *queue[T]) Size() int {
	return q.back - q.front
}

func (q *queue[T]) Print() {
	for i := q.front; i < q.back; i++ {
		fmt.Printf(" %v ||", q.values[i])
	}
	fmt.Println()
}
