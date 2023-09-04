package main

import (
	"testing"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

func TestSortStack(t *testing.T) {
	type args struct {
		s           utils.Stack[int]
		isAscending bool
	}
	tests := []struct {
		name string
		args args
		want utils.Stack[int]
	}{
		{
			name: "ascending order",
			args: args{
				s: func() utils.Stack[int] {
					s := utils.NewStack[int](5)
					s.Push(5)
					s.Push(4)
					s.Push(3)
					s.Push(2)
					s.Push(1)
					return s
				}(),
				isAscending: true,
			},
			want: func() utils.Stack[int] {
				s := utils.NewStack[int](5)
				s.Push(1)
				s.Push(2)
				s.Push(3)
				s.Push(4)
				s.Push(5)
				return s
			}(),
		},
		{
			name: "descending order",
			args: args{
				s: func() utils.Stack[int] {
					s := utils.NewStack[int](5)
					s.Push(1)
					s.Push(2)
					s.Push(3)
					s.Push(4)
					s.Push(5)
					return s
				}(),
				isAscending: false,
			},
			want: func() utils.Stack[int] {
				s := utils.NewStack[int](5)
				s.Push(5)
				s.Push(4)
				s.Push(3)
				s.Push(2)
				s.Push(1)
				return s
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortStack(tt.args.s, tt.args.isAscending)
			for !got.IsEmpty() {
				gotVal := got.Pop()
				wantVal := tt.want.Pop()
				if gotVal != wantVal {
					t.Errorf("SortStack() = %v, want %v", gotVal, wantVal)
				}
			}
		})
	}

}
