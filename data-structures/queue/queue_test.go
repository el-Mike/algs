package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)

	assert.False(t, queue.IsEmpty())

	value, err := queue.Dequeue()
	if err != nil {
		return
	}

	assert.Equal(t, 1, value)
	queue.Dequeue()

	assert.True(t, queue.IsEmpty())
}
