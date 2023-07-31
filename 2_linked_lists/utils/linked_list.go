package utils

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

func NewNode(v string) *Node {
	return &Node{Value: v}
}

func ArrayToLinkedList(values []string) *Node {
	if len(values) == 0 {
		return nil
	}

	root := NewNode(values[0])
	ll := root

	for i := 1; i < len(values); i++ {
		ll.Next = NewNode(values[i])
		ll = ll.Next
	}

	return root
}

func (ll *Node) Print(header string) {
	if header != "" {
		fmt.Println(header)
	}

	for ll != nil {
		fmt.Print(ll.Value, " -> ")
		ll = ll.Next
	}
	fmt.Print("nil\n\n")
}
