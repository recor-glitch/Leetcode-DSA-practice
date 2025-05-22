package Array

import "sync"

type Heap struct {
	items []int
	mutex sync.Mutex
}

func (h *Heap) heapifyUp(index int) {
	parent := index / 2
	if index > 0 && h.items[index] > h.items[parent] {
		temp := h.items[index]
		h.items[index] = h.items[parent]
		h.items[parent] = temp

		h.heapifyUp(parent)
	}
}

func (h *Heap) insertElement(element int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.items = append(h.items, element)
	h.heapifyUp(len(h.items) - 1)
}
