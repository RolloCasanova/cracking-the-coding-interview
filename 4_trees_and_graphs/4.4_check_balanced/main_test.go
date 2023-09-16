package main

import (
	"testing"

	"github.com/RolloCasanova/ctci-go/4_trees_and_graphs/utils/bst"
)

func TestCheckBalanced(t *testing.T) {
	type args struct {
		bt     *bst.BST
		height int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no nodes",
			args: args{
				bt:     nil,
				height: 0,
			},
			want: 0,
		},
		{
			name: "one node",
			args: args{
				bt:     bst.New(1),
				height: 0,
			},
			want: 1,
		},
		{
			name: "perfect tree of 3 nodes",
			args: args{
				bt:     bst.New(2, 1, 3),
				height: 0,
			},
			want: 2,
		},
		{
			name: "perfect tree of 7 nodes",
			args: args{
				bt:     bst.New(4, 2, 6, 1, 3, 5, 7),
				height: 0,
			},
			want: 3,
		},
		{
			name: "full tree except for one leaf",
			args: args{
				bt:     bst.New(4, 2, 6, 1, 3, 5),
				height: 0,
			},
			want: 3,
		},
		{
			name: "two node difference between left and right",
			args: args{
				bt:     bst.New(4, 2, 6, 1, 3, 5, 7, 8, 9),
				height: 0,
			},
			want: -1,
		},
		{
			name: "unbalanced tree - skewed left",
			args: args{
				bt: bst.New(4, 3, 2, 1),
			},
			want: -1,
		},
		{
			name: "unbalanced tree - skewed right",
			args: args{
				bt: bst.New(1, 2, 3, 4),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckBalanced(tt.args.bt, tt.args.height); got != tt.want {
				t.Errorf("CheckBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "negative",
			args: args{
				n: -1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.n); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
