package Array

import "sync"

type Stack[T any] struct {
	items []T
	mutex sync.Mutex
}

func NewEmptyStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
	}
}

func (s *Stack[T]) IsEmptyStack() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Push(item T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	index := len(s.items) - 1
	element := s.items[index]
	s.items = s.items[:index]

	return element, true
}

func (s *Stack[T]) Top() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.items[len(s.items)-1]
}
