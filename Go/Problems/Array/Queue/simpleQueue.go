package Array

import "sync"

type Queue[T any] struct {
	items []T
	mutex sync.Mutex
}

func NewEmptyQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: []T{},
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	element := q.items[0]
	q.items = q.items[1:]

	return element, true
}

func (q *Queue[T]) Front() T {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		var zero T
		return zero
	}

	return q.items[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
