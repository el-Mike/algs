package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	firstNode := NewNode(1, nil)

	list := NewLinkedList(firstNode)

	list.Append(2)

	i := list.Size()

	assert.Equal(t, 2, i)
}
