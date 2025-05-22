package Array

import (
	"sync"
)

type QueueItem[T any] struct {
	value  T
	weight int
}

type PriorityQueue[T any] struct {
	items []QueueItem[T]
	mutex sync.Mutex
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: []QueueItem[T]{},
	}
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *PriorityQueue[T]) Enqueue(item *QueueItem[T]) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, *item)
}

