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

// PadLinkedList pads the linked list with n nil nodes to the left of the linked list ll
func (ll *Node[T]) PadLinkedList(n int) *Node[T] {
	for i := 0; i < n; i++ {
		newNode := &Node[T]{Next: ll}
		ll = newNode
	}

	return ll
}

// Reverse reverses the linked list while preserving ll as the root
// it returns the new root of the reversed linked list
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

// CreateIntersectingLists returns two linked lists that intersect at the third linked list
// it returns two linked lists that might or not intersect
//
// Example:
// ll1: 1 -> 2 -> 3 -> nil
// ll2: 4 -> 5 -> nil
// ll3: 9 -> nil
//
// ll1, ll2 = CreateIntersectingLists(ll1, ll2, ll3) will return:
// ll1: 1 -> 2 -> 3 -> 9 -> nil
// ll2: 4 -> 5 -> 9 -> nil
func CreateIntersectingLists[T NodeVal](ll1, ll2, ll3 *Node[T]) (*Node[T], *Node[T]) {
	root1, root2 := ll1, ll2

	for ll1.Next != nil {
		ll1 = ll1.Next
	}

	for ll2.Next != nil {
		ll2 = ll2.Next
	}

	ll1.Next = ll3
	ll2.Next = ll3

	return root1, root2
}
