package goqueue

import (
	"sync"
)

// Queue data structure
type Queue struct {
	items []interface{}
	mutex sync.Mutex
}

// Instantiate a new queue
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Add an element to the end of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
}

// Remove and return the element at the front of the queue
func (q *Queue) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Return false if items are in the queue
func (q *Queue) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.items) == 0
}

// Return the number of elements in the queue
func (q *Queue) Count() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.items)
}
