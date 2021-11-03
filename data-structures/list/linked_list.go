package list

import (
	"fmt"
	"sync"
)

type ForEachCallback func(node *Node)
type FindCallback func(node *Node) bool
type Comparator func(a, b *Node) bool

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
		cb(current)

		current = current.Next
	}
}

// Find - iterates over the LinkedList, and returns the Node that passed callback returns
// true for. If no such Node is found, nil is returned.
func (l *LinkedList) Find(cb FindCallback) *Node {
	current := l.head

	for current != nil {
		found := cb(current)

		if found {
			return current
		}

		current = current.Next
	}

	return nil
}

// Prepend - adds an element to the beginning of LinkedList, making it the new head.
func (l *LinkedList) Prepend(data interface{}) {
	l.Lock()
	defer l.Unlock()

	newNode := NewNode(data, l.head)

	l.head = newNode
}

// Append - adds an element at the end of LinkedList.
func (l *LinkedList) Append(data interface{}) {
	newNode := NewNode(data, nil)

	lastNode := l.Find(func(node *Node) bool {
		return node.Next == nil
	})

	if lastNode != nil {
		lastNode.Lock()
		lastNode.Next = newNode
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

		current = current.Next
	}

	return nil
}

// InsertAt - adds an element at specified position.
func (l *LinkedList) InsertAt(data interface{}, position int) {
	newNode := NewNode(data, nil)

	current := l.head

	// i == 1 works as a offset - we want to get the Node before the
	// insert position, therefore we "subtract" one iteration.
	for i := 1; i < position; i += 1 {
		if current == nil {
			return
		}

		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
}

//  DeleteAt - removes an element at given position from LinkedList.
func (l *LinkedList) DeleteAt(position int) {
	current := l.head

	for i := 1; i < position; i += 1 {
		if current == nil {
			return
		}

		current = current.Next
	}

	current.Next = current.Next.Next
}

// RemoveFirst - removes first element from LinkedList.
func (l *LinkedList) RemoveFirst() {
	l.Lock()
	defer l.Unlock()

	if l.head != nil {
		l.head = l.head.Next
	}
}

// RemoveLast - removes last element from LinkedList.
func (l *LinkedList) RemoveLast() {
	l.Lock()
	defer l.Unlock()

	secondLastNode := l.Find(func(node *Node) bool {
		return node.Next.Next == nil
	})

	if secondLastNode != nil {
		secondLastNode.Lock()

		secondLastNode.Next = nil

		secondLastNode.Unlock()
	}
}

// Sort - sorts LinkedList by Nodes' values. Uses BubbleSort approach.
func (l *LinkedList) Sort(comparator Comparator) {
	current := l.head

	if current == nil {
		return
	}

	var index *Node

	for current != nil {
		index = current.Next

		for index != nil {
			if comparator(current, index) {
				current.Lock()
				index.Lock()

				current.Data, index.Data = index.Data, current.Data

				current.Unlock()
				index.Unlock()
			}

			index = index.Next
		}

		current = current.Next
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
		listData = append(listData, node.Data.(int))
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
	list.Sort(func(a, b *Node) bool {
		return a.Data.(int) > b.Data.(int)
	})
	list.Print()
}
