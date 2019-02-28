// Package bst Binary Search Tree
//
// Depth-first search: in-order, pre-order, post-order
// Breadth-first search: level-order
package bst

import (
	"errors"
	"fmt"

	deque "github.com/bsnux/go-cookbook/DequeGeneric"
)

// Node represents a tree node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Tree reprensents the root node of the tree
type Tree struct {
	Root *Node
}

// Insert a new node
func (n *Node) Insert(value int) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	// `switch` is usually faster than `if-else`
	switch {
	// value already inserted do nothing
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return nil
		}
		return n.Left.Insert(value)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return nil
		}
		return n.Right.Insert(value)
	}
	return nil
}

// Insert element in tree
func (t *Tree) Insert(value int) error {

	if t.Root == nil {
		t.Root = &Node{Value: value}
		return nil
	}

	return t.Root.Insert(value)
}

// Traverse prints sorted values traversing the tree in-order
func (t *Tree) Traverse(n *Node) {
	if n == nil {
		return
	}
	t.Traverse(n.Left)
	fmt.Println(n.Value)
	t.Traverse(n.Right)
}

// TraversePos prints values in post-order
func (t *Tree) TraversePos(n *Node) {
	if n == nil {
		return
	}
	t.TraversePos(n.Left)
	t.TraversePos(n.Right)
	fmt.Println(n.Value)
}

// TraversePre prints values in pre-order
func (t *Tree) TraversePre(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Value)
	t.TraversePre(n.Left)
	t.TraversePre(n.Right)
}

// TraverseLevel prints values level-order
func (t *Tree) TraverseLevel(n *Node) {
	d := deque.NewDeque()

	d.PushRight(n)

	for !d.IsEmpty() {
		item := d.PopLeft()
		fmt.Printf("%d, ", item.(*Node).Value)
		if item.(*Node).Left != nil {
			d.PushRight(item.(*Node).Left)
		}
		if item.(*Node).Right != nil {
			d.PushRight(item.(*Node).Right)
		}
	}
	fmt.Println("")
}

// Exists find a node a return true if node exists in tree
func (n *Node) Exists(value int) bool {
	if n == nil {
		return false
	}
	switch {
	case value == n.Value:
		return true
	case value < n.Value:
		return n.Left.Exists(value)
	default:
		return n.Right.Exists(value)
	}
}

// Exists find if value exists in tree
func (t *Tree) Exists(value int) bool {
	if t.Root == nil {
		return false
	}
	return t.Root.Exists(value)
}
