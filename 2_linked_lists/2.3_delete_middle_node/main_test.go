package main

import (
	"testing"

	"github.com/RolloCasanova/ctci-go/2_linked_lists/utils"
)

var ll *node

func Test_deleteMiddleNode(t *testing.T) {
	type args struct {
		node *node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "delete middle node",
			args: args{
				node: func() *node {
					ll = utils.ArrayToLinkedList[string]([]string{"a", "b", "c"})
					return ll.Next
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			deleteMiddleNode(tt.args.node)
		})
	}
}
