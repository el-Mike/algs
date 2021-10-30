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

func TestLinkedList_InsertAt(t *testing.T) {
	firstNode := NewNode(1, nil)

	list := NewLinkedList(firstNode)

	list.Append(2)
	list.Append(3)
	list.Append(4)

	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 3, list.GetAt(2).data)

	list.InsertAt(5, 2)

	assert.Equal(t, 5, list.Size())
	assert.Equal(t, 5, list.GetAt(2).data)
}

func TestLinkedList_DeleteAt(t *testing.T) {
	firstNode := NewNode(1, nil)

	list := NewLinkedList(firstNode)

	list.Append(2)
	list.Append(3)
	list.Append(4)

	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 3, list.GetAt(2).data)

	list.DeleteAt(2)

	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 4, list.GetAt(2).data)
}
