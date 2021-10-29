package stack

import (
	"errors"
	"sync"
)

const (
	STACK_EMPTY_ERROR = "Stack is empty!"
)

type Stack struct {
	// Makes stack implementation thread-safe.
	sync.RWMutex

	data []int
}

// NewStack - returns new Stack instance.
func NewStack() *Stack {
	return &Stack{
		data: []int{},
	}
}

// Push - push value to the top.
func (s *Stack) Push(value int) {
	s.Lock()
	defer s.Unlock()

	s.data = append(s.data, value)
}

// Pop - pop value from the top.
func (s *Stack) Pop() (int, error) {
	s.Lock()
	defer s.Unlock()

	l := len(s.data)

	if l == 0 {
		return 0, errors.New(STACK_EMPTY_ERROR)
	}

	value := s.data[l-1]
	s.data = s.data[:(l - 1)]

	return value, nil
}

// IsEmpty - returns true if Stack is empty, false otherwise.
func (s *Stack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	return len(s.data) == 0
}

// Peek - returns the top value from the Stack without removin it.
func (s *Stack) Peek() (int, error) {
	s.RLock()
	defer s.RUnlock()

	l := len(s.data)

	if l == 0 {
		return 0, errors.New(STACK_EMPTY_ERROR)
	}

	return s.data[l-1], nil
}
