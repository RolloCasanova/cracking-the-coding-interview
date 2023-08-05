package utils

import "fmt"

type NodeVal interface {
	~int | ~string
}

type Node[T NodeVal] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T NodeVal](v T) *Node[T] {
	return &Node[T]{Value: v}
}

// ArrayToLinkedList create a linked list from an array of T (int or string)
func ArrayToLinkedList[T NodeVal](values []T) *Node[T] {
	if len(values) == 0 {
		return nil
	}

	root := &Node[T]{
		Value: values[0],
	}

	ll := root

	for i := 1; i < len(values); i++ {
		ll.Next = NewNode[T](values[i])
		ll = ll.Next
	}

	return root
}

// Print prints the linked list with a given header (optional)
// it also prints the arrows and nil at the end for readability
func (ll *Node[T]) Print(header string) {
	if header != "" {
		fmt.Println(header)
	}

	for ll != nil {
		fmt.Print(ll.Value, " -> ")
		ll = ll.Next
	}

	fmt.Print("nil\n\n")
}

// Len returns the number of nodes in the linked list
func (ll *Node[T]) Len() int {
	var count int

	for ll != nil {
		count++
		ll = ll.Next
	}

	return count
}

// PadLinkedList pads the linked list with zero values until it reaches the given length
func (ll *Node[T]) PadLinkedList(length int) {
	for i := 0; i < length-1; i++ {
		node := &Node[T]{}
		node.Next = ll
		ll = node
	}
}

// Reverse reverses the linked list while preserving ll as the root
// it returns the new root
func (ll *Node[T]) Reverse() *Node[T] {
	var prev *Node[T]

	for ll != nil {
		next := ll.Next
		ll.Next = prev
		prev = ll
		ll = next
	}

	return prev
}

// String returns a string representation of the linked list (e.g. 1 -> 2 -> 3 -> nil is "123")
func (ll *Node[T]) String() string {
	var s string

	for ll != nil {
		s += fmt.Sprintf("%v", ll.Value)
		ll = ll.Next
	}

	return s
}
