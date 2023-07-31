package main

import (
	"testing"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

var ll *utils.Node

func Test_deleteMiddleNode(t *testing.T) {
	type args struct {
		node *utils.Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "delete middle node",
			args: args{
				node: func() *utils.Node {
					ll = utils.ArrayToLinkedList([]string{"a", "b", "c"})
					return ll.Next
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll.Print("Linked list: ")
			deleteMiddleNode(tt.args.node)
			ll.Print("Linked list with middle node deleted: ")
		})
	}
}
