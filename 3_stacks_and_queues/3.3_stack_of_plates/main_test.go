package main

import (
	"reflect"
	"testing"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

func TestStackSet(t *testing.T) {
	type args struct {
		stackSize int
		numStacks int
		n         int
	}
	tests := []struct {
		name string
		args args
		want utils.StackSet
	}{
		{
			name: "stack set of 1 stack with 1 value, n = 0",
			args: args{
				stackSize: 1,
				numStacks: 1,
				n:         0,
			},
			want: func() utils.StackSet {
				return utils.NewStackSet(1, 1)
			}(),
		},
		{
			name: "stack set of 1 stack with 1 value, n = 1",
			args: args{
				stackSize: 1,
				numStacks: 1,
				n:         1,
			},
			want: func() utils.StackSet {
				ss := utils.NewStackSet(1, 1)
				ss.Push(1)

				return ss
			}(),
		},
		{
			name: "stack set of 3 stacks with 3 values, n = 9",
			args: args{
				stackSize: 3,
				numStacks: 3,
				n:         9,
			},
			want: func() utils.StackSet {
				ss := utils.NewStackSet(3, 3)
				for i := 1; i <= 9; i++ {
					ss.Push(i)
				}

				return ss
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StackSet(tt.args.stackSize, tt.args.numStacks, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StackSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
