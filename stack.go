package goqueue

import (
	"sync"
)

// Stack data structure
type Stack struct {
	items []interface{}
	mutex sync.Mutex
}

// Instantiate a new stack
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Add a new element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = append(s.items, item)
}

// Remove and return the top element of the stack
func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return item
}

// Return true if the stack is empty
func (s *Stack) IsEmpty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.items) == 0
}

// Return the number of elements in the stack
func (s *Stack) Count() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.items)
}