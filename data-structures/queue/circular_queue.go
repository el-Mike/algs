package queue

import "sync"

type CircularQueue struct {
	// Makes queue implementation thread-safe.
	sync.RWMutex

	data []int

	size int
	head int
	tail int
}

// NewCircularQueue - returns new CircularQueue instance.
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		data: make([]int, size),
		size: size,
		head: -1,
		tail: -1,
	}
}

// IsFull - returns true if CircularQueue is full, false otherwise.
func (cq *CircularQueue) IsFull() bool {
	cq.RLock()
	defer cq.RUnlock()

	return cq.isFull()
}

// isFull - helper function for avoiding recursive locks.
func (cq *CircularQueue) isFull() bool {
	if cq.head == 0 && cq.tail == (cq.size-1) {
		return true
	}

	if cq.head == (cq.tail - 1) {
		return true
	}

	return false
}

// IsEmpty - returns true if CircularQueue is empty, false otherwise.
func (cq *CircularQueue) IsEmpty() bool {
	cq.RLock()
	defer cq.RUnlock()

	return cq.isEmpty()
}

// isEmpty - helper function for avoiding recursive locks.
func (cq *CircularQueue) isEmpty() bool {
	return cq.head == -1
}

/// Enqueue - adds an element to the end of the CircularQueue. Note that since it's
// circular, end does not neccessarily means the last index.
func (cq *CircularQueue) Enqueue(value int) error {
	cq.Lock()
	defer cq.Unlock()

	if cq.isFull() {
		return NewQueueFullError()
	}

	if cq.head == -1 {
		cq.head = 0
	}

	cq.tail = (cq.tail + 1) % cq.size

	cq.data[cq.tail] = value

	return nil
}

// Dequeue - removes and returns an element from the front of the CircularQueue.
// Note that since it's circular, front does not neccessarily means the first index.
func (cq *CircularQueue) Dequeue() (int, error) {
	cq.Lock()
	defer cq.Unlock()

	if cq.isEmpty() {
		return 0, NewQueueEmptyError()
	}

	value := cq.data[cq.head]

	// If both head and tail are equal, that means there is no more elements in the queue.
	if cq.head == cq.tail {
		cq.head = -1
		cq.tail = -1
	} else {
		cq.head = (cq.head + 1) % cq.size
	}

	return value, nil
}
