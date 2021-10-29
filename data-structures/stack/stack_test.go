package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)

	assert.False(t, stack.IsEmpty())

	value, err := stack.Pop()
	if err != nil {
		return
	}

	assert.Equal(t, 2, value)
	stack.Pop()
	assert.True(t, stack.IsEmpty())
}
