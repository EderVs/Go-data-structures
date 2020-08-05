package lists

import (
	"fmt"
	"strings"
)

// LinkedListNode is a node in LinkedList.
type LinkedListNode struct {
	Value interface{}
	next  *LinkedListNode
}

// Next returns next node.
func (node *LinkedListNode) Next() interface{} {
	return node.Next
}

// LinkedList is a simple implemention of the data structure.
type LinkedList struct {
	head   *LinkedListNode
	tail   *LinkedListNode
	length int
}

// IsEmpty returns if the list is empty.
func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

// Length returns list length.
func (list *LinkedList) Length() int {
	return list.length
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

// Append inserts the value into the end of the list.
func (list *LinkedList) Append(value interface{}) {
	node := &LinkedListNode{Value: value}
	list.length++

	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}

	list.tail = node
}

// Insert inserts the value into the start of the list.
func (list *LinkedList) Insert(value interface{}) {
	node := &LinkedListNode{Value: value}
	list.length++

	node.next = list.head
	if list.tail == nil {
		list.tail = node
	}
	list.head = node
}

// Pop deletes the first value if there is one.
func (list *LinkedList) Pop() (*interface{}, bool) {
	if list.head == nil {
		return nil, false
	}
	if list.head == list.tail {
		list.tail = nil
	}
	node := list.head
	list.head = node.next
	list.length--
	return &node.Value, true
}

// GetIthNode gets the ith node.
func (list *LinkedList) GetIthNode(i int) (*LinkedListNode, bool) {
	if i < 0 {
		return nil, false
	}
	node := list.head
	for j := 0; j < i && node != nil; j++ {
		node = node.next
	}
	if node == nil {
		return nil, false
	}
	return node, true
}

// Get gets the ith node.
func (list *LinkedList) Get(i int) (*interface{}, bool) {
	if i == 0 {
		return &list.head.Value, true
	}
	node, found := list.GetIthNode(i)
	if !found {
		return nil, false
	}
	return &node.Value, true
}

// PopIth gets the ith element in the list.
func (list *LinkedList) PopIth(i int) (*interface{}, bool) {
	if i < 0 {
		return nil, false
	}
	if i == 0 {
		return list.Pop()
	}
	pNode, found := list.GetIthNode(i - 1)
	if !found || pNode.next == nil {
		return nil, false
	}
	node := pNode.next
	pNode.next = pNode.next.next
	if list.tail == node {
		list.tail = pNode
	}
	list.length--
	return &node.Value, true
}

// GetNodeByValue look for the value and returns the first node that contains it.
func (list *LinkedList) GetNodeByValue(value interface{}) (*LinkedListNode, bool) {
	node := list.head
	for node != nil {
		if node.Value == value {
			return node, true
		}
		node = node.next
	}
	return nil, false
}

// Contains checks if the value is in the list.
func (list *LinkedList) Contains(value interface{}) bool {
	_, found := list.GetNodeByValue(value)
	return found
}
