package queues

import (
	"github.com/edervs/go-data-structures/lists"
)

// Queue is a simple implementation using LinkedList.
type Queue struct {
	lists.LinkedList
}

// Add adds a value to Queue.
func (q *Queue) Add(value interface{}) {
	q.Append(value)
}

// Insert adds a value to Queue.
func (q *Queue) Insert(value interface{}) {
	q.Append(value)
}

// Enqueue adds a value to Queue.
func (q *Queue) Enqueue(value interface{}) {
	q.Append(value)
}

// Dequeue pops element from the .
func (q *Queue) Dequeue() (*interface{}, bool) {
	return q.Pop()
}
