package main

import (
	"reflect"
	"testing"

	ll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

func TestListOfDepths(t *testing.T) {
	type args struct {
		t *BST
	}
	tests := []struct {
		name string
		args args
		want []ll.List
	}{
		{
			name: "a tree with 1 node",
			args: args{
				t: &BST{Value: 1},
			},
			want: func() []ll.List {
				lls := make([]ll.List, 1)
				lls[0] = ll.List{}
				lls[0].Append(1)

				return lls
			}(),
		},
		{
			name: "a tree with 2 nodes",
			args: args{
				t: &BST{
					Value: 1,
					Left:  &BST{Value: 2},
				},
			},
			want: func() []ll.List {
				lls := make([]ll.List, 2)
				lls[0] = ll.List{}
				lls[0].Append(1)

				lls[1] = ll.List{}
				lls[1].Append(2)

				return lls
			}(),
		},
		{
			name: "a tree with 3 nodes (2 levels)",
			args: args{
				t: &BST{
					Value: 1,
					Left:  &BST{Value: 2},
					Right: &BST{Value: 3},
				},
			},
			want: func() []ll.List {
				lls := make([]ll.List, 2)
				lls[0] = ll.List{}
				lls[0].Append(1)

				lls[1] = ll.List{}
				lls[1].Append(2)
				lls[1].Append(3)

				return lls
			}(),
		},
		{
			name: "a perfect tree with 7 nodes (3 levels)",
			args: args{
				t: &BST{
					Value: 4,
					Left: &BST{
						Value: 2,
						Left:  &BST{Value: 1},
						Right: &BST{Value: 3},
					},
					Right: &BST{
						Value: 6,
						Left:  &BST{Value: 5},
						Right: &BST{Value: 7},
					},
				},
			},
			want: func() []ll.List {
				lls := make([]ll.List, 3)
				lls[0] = ll.List{}
				lls[0].Append(4)

				lls[1] = ll.List{}
				lls[1].Append(2)
				lls[1].Append(6)

				lls[2] = ll.List{}
				lls[2].Append(1)
				lls[2].Append(3)
				lls[2].Append(5)
				lls[2].Append(7)

				return lls
			}(),
		},
		{
			name: "a perfect tree with 15 nodes (4 levels)",
			args: args{
				t: &BST{
					Value: 8,
					Left: &BST{
						Value: 4,
						Left: &BST{
							Value: 2,
							Left:  &BST{Value: 1},
							Right: &BST{Value: 3},
						},
						Right: &BST{
							Value: 6,
							Left:  &BST{Value: 5},
							Right: &BST{Value: 7},
						},
					},
					Right: &BST{
						Value: 12,
						Left: &BST{
							Value: 10,
							Left:  &BST{Value: 9},
							Right: &BST{Value: 11},
						},
						Right: &BST{
							Value: 14,
							Left:  &BST{Value: 13},
							Right: &BST{Value: 15},
						},
					},
				},
			},
			want: func() []ll.List {
				lls := make([]ll.List, 4)
				lls[0] = ll.List{}
				lls[0].Append(8)

				lls[1] = ll.List{}
				lls[1].Append(4)
				lls[1].Append(12)

				lls[2] = ll.List{}
				lls[2].Append(2)
				lls[2].Append(6)
				lls[2].Append(10)
				lls[2].Append(14)

				lls[3] = ll.List{}
				lls[3].Append(1)
				lls[3].Append(3)
				lls[3].Append(5)
				lls[3].Append(7)
				lls[3].Append(9)
				lls[3].Append(11)
				lls[3].Append(13)
				lls[3].Append(15)

				return lls
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListOfDepths(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListOfDepths() = %v, want %v", got, tt.want)
			}
		})
	}
}
