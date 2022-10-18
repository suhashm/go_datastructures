package stack

import (
	"errors"
	"fmt"
	"sync"
)

type stack[T any] struct {
	sync.RWMutex
	top          int
	defaultValue T
	capacity     int
	values       []T
}

func New[T any](defaultValue T) *stack[T] {
	return NewWithCapacity(10000, defaultValue)
}

func NewWithCapacity[T any](capacity int, defaultValue T) *stack[T] {
	return &stack[T]{
		top:          -1,
		defaultValue: defaultValue,
		capacity:     capacity,
		values:       make([]T, capacity),
	}
}

func (s *stack[T]) IsFull() bool {
	s.RLock()
	defer s.RUnlock()
	return s.top == s.capacity-1
}

func (s *stack[T]) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.top == -1
}

func (s *stack[T]) Push(val T) error {
	if s.IsFull() {
		return errors.New("stack full")
	}

	s.Lock()
	defer s.Unlock()
	s.top++

	s.values[s.top] = val
	return nil
}

func (s *stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return s.defaultValue, errors.New("stack empty")
	}

	s.Lock()
	defer s.Unlock()

	val := s.values[s.top]
	s.top--
	return val, nil

}

func (s *stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return s.defaultValue, errors.New("stack empty")
	}

	s.RLock()
	defer s.RUnlock()

	val := s.values[s.top]
	return val, nil
}

func (s *stack[T]) Print() {
	s.RLock()
	defer s.RUnlock()

	for i := 0; i <= s.top; i++ {
		fmt.Printf(" %v ||", s.values[i])
	}
	fmt.Println()
}
