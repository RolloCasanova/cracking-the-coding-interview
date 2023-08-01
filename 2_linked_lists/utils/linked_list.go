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

// create a linked list from an array of T (int or string)
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
