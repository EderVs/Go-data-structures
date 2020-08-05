package stacks

import (
	"github.com/edervs/go-data-structures/lists"
)

// Stack is a simple implmentation using LinkedList.
type Stack struct {
	lists.LinkedList
}

// Add adds value to stack.
func (s *Stack) Add(value interface{}) {
	s.Insert(value)
}

// Append adds value to stack.
func (s *Stack) Append(value interface{}) {
	s.Insert(value)
}
