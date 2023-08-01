package main

import (
	"reflect"
	"testing"
)

func Test_partitionLinkedList(t *testing.T) {
	type args struct {
		ll        *node
		partition int
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "a partition value greater than all nodes should return the same linked list",
			args: args{ll: &node{Value: 1, Next: &node{Value: 2, Next: &node{Value: 3, Next: nil}}}, partition: 4},
			want: &node{Value: 1, Next: &node{Value: 2, Next: &node{Value: 3, Next: nil}}},
		},
		{
			name: "a partition value less than all nodes should return the same linked list",
			args: args{ll: &node{Value: 1, Next: &node{Value: 2, Next: &node{Value: 3, Next: nil}}}, partition: 0},
			want: &node{Value: 1, Next: &node{Value: 2, Next: &node{Value: 3, Next: nil}}},
		},
		{
			name: "a partition value in between nodes should return a linked list partitioned around the partition value",
			args: args{ll: &node{Value: 5, Next: &node{Value: 1, Next: &node{Value: 3, Next: nil}}}, partition: 4},
			want: &node{Value: 1, Next: &node{Value: 3, Next: &node{Value: 5, Next: nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLinkedList(tt.args.ll, tt.args.partition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
