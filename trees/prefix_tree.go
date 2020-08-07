package trees

import (
	"fmt"
	"strings"
)

// PrefixNode is the Node of PrefixTree.
type PrefixNode struct {
	value    interface{}
	children map[interface{}]*PrefixNode
	isFinal  bool
	hasValue bool
}

func (node *PrefixNode) insertEachElement(value []interface{}, from int) {
	if len(value) == 0 {
		return
	}
	var element interface{}
	for i := from; i < len(value); i++ {
		element = value[i]
		newNode := createNode()
		newNode.setValue(element)
		node.children[element] = newNode
		node = newNode
	}
	node.isFinal = true
}

func (node *PrefixNode) getFinalChainValues() *[][]interface{} {
	chains := make([][]interface{}, 0)
	if len(node.children) > 0 {
		for _, child := range node.children {
			chains = append(chains, *child.getFinalChainValues()...)
		}
	}
	for i := range chains {
		chains[i] = append([]interface{}{node.value}, chains[i]...)
	}
	if node.isFinal {
		chains = append(chains, []interface{}{node.value})
	}
	return &chains
}

func (node *PrefixNode) setValue(value interface{}) {
	node.value = value
	node.hasValue = true
}

// isLeaf returns True if it doesn't have any children.
func (node *PrefixNode) isLeaf() bool {
	return len(node.children) == 0
}

// string returns the string representation of a Node
func (node *PrefixNode) string() string {
	if node == nil {
		return ""
	}
	var b strings.Builder
	if node.hasValue {
		b.WriteString(fmt.Sprintf("%v", node.value))
	}
	b.WriteString("(")
	numChildren := len(node.children)
	i := 0
	for _, child := range node.children {
		b.WriteString(child.string())
		if i < numChildren-1 {
			b.WriteString(", ")
		}
		i++
	}
	b.WriteString(")")

	return b.String()
}

// searchLastNode returns the last node with the prefix and the index in the prefix.
func (node *PrefixNode) searchLastNode(prefix []interface{}) (*PrefixNode, int) {
	i := 0
	n := len(prefix)
	for !node.isLeaf() && i < n {
		_, found := node.children[prefix[i]]
		if !found {
			break
		}
		node = node.children[prefix[i]]
		i++
	}
	return node, i - 1
}

func createNode() *PrefixNode {
	node := new(PrefixNode)
	node.children = make(map[interface{}]*PrefixNode)
	return node
}

// PrefixTree is the implementation of a Trie.
type PrefixTree struct {
	root   *PrefixNode
	length int
}

// String returns a string representation of PrefixTree.
func (pT *PrefixTree) String() string {
	return pT.root.string()
}

// Length returns the length of the PrefixTree.
func (pT *PrefixTree) Length() int {
	return pT.length
}

// searchLastNode returns the last node with the prefix.
func (pT *PrefixTree) searchLastNode(prefix []interface{}) (*PrefixNode, int) {
	if pT.root == nil {
		return nil, -1
	}
	return pT.root.searchLastNode(prefix)
}

// Insert adds a value to the PrefixTree.
func (pT *PrefixTree) Insert(value []interface{}) {
	defer func() { pT.length++ }()
	if pT.root == nil {
		node := createNode()
		node.insertEachElement(value, 0)
		pT.root = node
		return
	}
	node, from := pT.searchLastNode(value)
	node.insertEachElement(value, from+1)
}

// GetValues returns all the values stored.
func (pT *PrefixTree) GetValues() [][]interface{} {
	values := make([][]interface{}, 0)
	if pT.root == nil {
		return values
	}
	for _, child := range pT.root.children {
		values = append(values, *child.getFinalChainValues()...)
	}
	return values
}

func addPrefixToValues(prefix []interface{}, values *[][]interface{}) {
	for i := range *values {
		(*values)[i] = append(prefix, (*values)[i]...)
	}
}

// SearchPossibles returns the values that has as some prefix in the prefix the value given.
func (pT *PrefixTree) SearchPossibles(prefix []interface{}) [][]interface{} {
	values := make([][]interface{}, 0)
	lastNode, to := pT.searchLastNode(prefix)
	if to == -1 {
		return values
	}
	values = *lastNode.getFinalChainValues()
	prefixInTrie := prefix[:to]
	addPrefixToValues(prefixInTrie, &values)
	return values
}

// Search returns the values that has as prefix the value given.
func (pT *PrefixTree) Search(prefix []interface{}) [][]interface{} {
	values := make([][]interface{}, 0)
	lastNode, to := pT.searchLastNode(prefix)
	if to+1 != len(prefix) {
		return values
	}
	values = *lastNode.getFinalChainValues()
	prefixInTrie := prefix[:to]
	addPrefixToValues(prefixInTrie, &values)
	return values
}
