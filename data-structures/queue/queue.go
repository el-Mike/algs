package queue

import (
	"errors"
	"sync"
)

const (
	QUEUE_EMPTY_ERROR = "Queue is empty!"
)

type Queue struct {
	// Makes queue implementation thread-safe.
	sync.RWMutex

	data []int
}

// NewQueue - returns new Queue instance.
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// Enqueue - ads an element to the end of the Queue.
func (q *Queue) Enqueue(value int) {
	q.Lock()
	defer q.Unlock()

	q.data = append(q.data, value)
}

// Dequeue - removes and returns an element from the front of the Queue.
func (q *Queue) Dequeue() (int, error) {
	q.Lock()
	defer q.Unlock()

	l := len(q.data)

	if l == 0 {
		return 0, errors.New(QUEUE_EMPTY_ERROR)
	}

	value := q.data[0]
	q.data = q.data[1:]

	return value, nil
}

// IsEmpty - returns true if Queue is empty, false otherwise.
func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.data) == 0
}

func (q *Queue) Peek() (int, error) {
	q.RLock()
	defer q.RUnlock()

	l := len(q.data)

	if l == 0 {
		return 0, errors.New(QUEUE_EMPTY_ERROR)
	}

	return q.data[0], nil
}
