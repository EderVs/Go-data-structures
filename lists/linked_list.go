package lists

import (
	"fmt"
	"strings"
)

type LinkedListNode struct {
	Value interface{}
	next  *LinkedListNode
}

// IsEmpty returns if the list is empty.
func (node *LinkedListNode) Next() interface{} {
	return node.Next
}

type LinkedList struct {
	head   *LinkedListNode
	tail   *LinkedListNode
	lenght int
}

// IsEmpty returns if the list is empty.
func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

// Lenght returns list lenght.
func (list *LinkedList) Lenght() int {
	return list.lenght
}

// Head returns the head of the list.
func (list *LinkedList) Head() *LinkedListNode {
	return list.head
}

// Tail returns the tail of the list.
func (list *LinkedList) Tail() *LinkedListNode {
	return list.tail
}

// String representation of the linked list.
func (list *LinkedList) String() string {
	var b strings.Builder
	node := list.head
	for node != nil {
		b.WriteString(fmt.Sprintf("%v", node.Value))
		if node.next != nil {
			b.WriteString(" -> ")
		}
		node = node.next
	}
	return b.String()
}

// Insert inserts the value into the list and returns the LinkedListNode created.
func (list *LinkedList) Insert(value interface{}) *LinkedListNode {
	node := &LinkedListNode{Value: value}
	list.lenght++

	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}

	list.tail = node

	return node
}
