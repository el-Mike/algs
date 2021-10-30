package list

import (
	"fmt"
	"sync"
)

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
	} else {
		l.Lock()
		l.head = newNode
		l.Unlock()
	}
}

// GetAt - returns an element at specified position.
func (l *LinkedList) GetAt(position int) *Node {
	current := l.head

	for i := 0; i <= position; i += 1 {
		if current == nil {
			return nil
		}

		if i == position {
			return current
		}

		current = current.next
	}

	return nil
}

// InsertAt - adds an element at specified position.
func (l *LinkedList) InsertAt(value, position int) {
	newNode := NewNode(value, nil)

	current := l.head

	// i == 1 works as a offset - we want to get the Node before the
	// insert position, therefore we "subtract" one iteration.
	for i := 1; i < position; i += 1 {
		if current == nil {
			return
		}

		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
}

//  DeleteAt - removes an element at given position from LinkedList.
func (l *LinkedList) DeleteAt(position int) {
	current := l.head

	for i := 1; i < position; i += 1 {
		if current == nil {
			return
		}

		current = current.next
	}

	current.next = current.next.next
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

// Sort - sorts LinkedList by Nodes' values. Uses BubbleSort approach.
func (l *LinkedList) Sort() {
	current := l.head

	if current == nil {
		return
	}

	var index *Node

	for current != nil {
		index = current.next

		for index != nil {
			if current.data > index.data {
				current.Lock()
				index.Lock()

				current.data, index.data = index.data, current.data

				current.Unlock()
				index.Unlock()
			}

			index = index.next
		}

		current = current.next
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

// Print - prints LinkedList's Nodes values as an array.
func (l *LinkedList) Print() {
	var listData []int

	l.ForEach(func(node *Node) {
		listData = append(listData, node.data)
	})

	fmt.Printf("%v\n", listData)
}

func RunLinkedList() {
	list := NewLinkedList(nil)

	testData := GetTestDataCopy()

	for _, x := range testData {
		list.Append(x)
	}

	list.Print()
	list.Sort()
	list.Print()
}
