package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	size := 2
	queue := NewCircularQueue(size)

	queue.Enqueue(1)
	queue.Enqueue(2)

	assert.True(t, queue.IsFull())

	assert.Error(t, queue.Enqueue(3))

	queue.Dequeue()
	queue.Dequeue()

	_, err := queue.Dequeue()

	assert.Error(t, err)
}
