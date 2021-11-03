package list

import "sync"

type Node struct {
	// Makes linked list implementation thread-safe.
	sync.RWMutex

	Data interface{}
	Next *Node
}

func NewNode(data interface{}, next *Node) *Node {
	return &Node{
		Data: data,
		Next: next,
	}
}
