package list

import "sync"

type ForEachCallback func(node *Node)
type FindCallback func(node *Node) bool

type Node struct {
	// Makes linked list implementation thread-safe.
	sync.RWMutex

	data int
	next *Node
}

func NewNode(value int, next *Node) *Node {
	return &Node{
		data: value,
		next: next,
	}
}

type LinkedList struct {
	// Makes linked list implementation thread-safe.
	sync.RWMutex

	head *Node
}

func NewLinkedList(node *Node) *LinkedList {
	return &LinkedList{
		head: node,
	}
}

// ForEach - iterates over the LinkedList, calling passed callback for every node.
func (l *LinkedList) ForEach(cb ForEachCallback) {
	if l.head == nil {
		return
	}

	current := l.head

	for current != nil {
		current.Lock()

		cb(current)

		current.Unlock()

		current = current.next
	}
}

// Find - iterates over the LinkedList, and returns the Node that passed callback returns
// true for. If no such Node is found, nil is returned.
func (l *LinkedList) Find(cb FindCallback) *Node {
	current := l.head

	for current != nil {
		current.Lock()

		found := cb(current)

		current.Unlock()

		if found {
			return current
		}

		current = current.next
	}

	return nil
}

// Prepend - adds an element to the beginning of LinkedList, making it the new head.
func (l *LinkedList) Prepend(value int) {
	l.Lock()
	defer l.Unlock()

	newNode := NewNode(value, l.head)

	l.head = newNode
}

// Append - adds an element at the end of LinkedList.
func (l *LinkedList) Append(value int) {
	newNode := NewNode(value, nil)

	lastNode := l.Find(func(node *Node) bool {
		return node.next == nil
	})

	if lastNode != nil {
		lastNode.Lock()

		lastNode.next = newNode

		lastNode.Unlock()
	}
}

// RemoveFirst - removes first element from LinkedList.
func (l *LinkedList) RemoveFirst() {
	l.Lock()
	defer l.Unlock()

	if l.head != nil {
		l.head = l.head.next
	}
}

// RemoveLast - removes last element from LinkedList.
func (l *LinkedList) RemoveLast() {
	l.Lock()
	defer l.Unlock()

	secondLastNode := l.Find(func(node *Node) bool {
		return node.next.next == nil
	})

	if secondLastNode != nil {
		secondLastNode.Lock()

		secondLastNode.next = nil

		secondLastNode.Unlock()
	}
}

// Size - returns the length of the list.
func (l *LinkedList) Size() int {
	i := 0

	l.ForEach(func(_ *Node) {
		i++
	})

	return i
}
