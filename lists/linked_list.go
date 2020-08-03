package lists

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

type LinkedList struct {
	Head   *LinkedListNode
	Tail   *LinkedListNode
	lenght int
}

// IsEmpty returns if the list is empty.
func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

// GetLenght returns list lenght.
func (list *LinkedList) GetLenght() int {
	return list.lenght
}

// Insert inserts the value into the list and returns the LinkedListNode created.
func (list *LinkedList) Insert(value interface{}) *LinkedListNode {
	node := &LinkedListNode{Value: value}
	list.lenght++

	if list.Head == nil {
		list.Head = node
	} else {
		list.Tail.Next = node
	}

	list.Tail = node

	return node
}
