package main

import (
	"reflect"
	"testing"
)

func TestMinimalTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *BST
	}{
		{
			name: "n = 0",
			args: args{nums: []int{}},
			want: nil,
		},
		{
			name: "n = 1",
			args: args{nums: []int{1}},
			want: &BST{Value: 1},
		},
		{
			name: "n = 2",
			args: args{nums: []int{1, 2}},
			want: &BST{
				Value: 2,
				Left:  &BST{Value: 1},
			},
		},
		{
			name: "n = 3",
			args: args{nums: []int{1, 2, 3}},
			want: &BST{
				Value: 2,
				Left:  &BST{Value: 1},
				Right: &BST{Value: 3},
			},
		},
		{
			name: "n = 7", // a perfect binary tree
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}},
			want: &BST{
				Value: 4,
				Left: &BST{
					Value: 2,
					Left: &BST{
						Value: 1,
					},
					Right: &BST{
						Value: 3,
					},
				},
				Right: &BST{
					Value: 6,
					Left: &BST{
						Value: 5,
					},
					Right: &BST{
						Value: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinimalTree(tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimalTree() = %v, want %v", got, tt.want)
			}

			got.print()
		})
	}
}
