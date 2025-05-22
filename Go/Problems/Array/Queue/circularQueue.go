package Array

import "sync"

type CircularQueue[T any] struct {
	items    []T
	head     int
	tail     int
	size     int
	capacity int
	mutex    sync.Mutex
}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	return &CircularQueue[T]{
		items:    make([]T, capacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: capacity,
	}
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *CircularQueue[T]) Enqueue(item T) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.size == q.capacity {
		return false
	}
	

	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
	q.size++
	return true
}

func (q *CircularQueue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.size == 0 {
		var zero T
		return zero, false
	}

	item := q.items[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return item, true
}

func (q *CircularQueue[T]) Front() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.size == 0 {
		var zero T
		return zero, false
	}

	return q.items[q.head], true
}
