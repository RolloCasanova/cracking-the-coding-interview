package main

import (
	"testing"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

func TestStackMin(t *testing.T) {
	type args struct {
		sm utils.StackMin
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "stack with 1 value",
			args: func() args {
				sm := utils.NewStackMin(1)
				sm.Push(1)

				return args{sm: sm}
			}(),
			want: 1,
		},
		{
			name: "stack with values in ascending order",
			args: func() args {
				sm := utils.NewStackMin(5)
				sm.Push(1)
				sm.Push(2)
				sm.Push(3)
				sm.Push(4)
				sm.Push(5)

				return args{sm: sm}
			}(),
			want: 1,
		},
		{
			name: "stack with values in descending order",
			args: func() args {
				sm := utils.NewStackMin(5)
				sm.Push(5)
				sm.Push(4)
				sm.Push(3)
				sm.Push(2)
				sm.Push(1)

				return args{sm: sm}
			}(),
			want: 1,
		},
		{
			name: "stack with values in random order",
			args: func() args {
				sm := utils.NewStackMin(5)
				sm.Push(3)
				sm.Push(5)
				sm.Push(1)
				sm.Push(4)
				sm.Push(2)

				return args{sm: sm}
			}(),
			want: 1,
		},
		{
			name: "stack with duplicate values",
			args: func() args {
				sm := utils.NewStackMin(5)
				sm.Push(1)
				sm.Push(2)
				sm.Push(2)
				sm.Push(3)
				sm.Push(4)

				return args{sm: sm}
			}(),
			want: 1,
		},
		{
			name: "stack with negative values",
			args: func() args {
				sm := utils.NewStackMin(5)
				sm.Push(-1)
				sm.Push(-2)
				sm.Push(-3)
				sm.Push(-4)
				sm.Push(-5)

				return args{sm: sm}
			}(),
			want: -5,
		},
		{
			name: "stack with negative and positive values",
			args: func() args {
				sm := utils.NewStackMin(5)
				sm.Push(-1)
				sm.Push(2)
				sm.Push(-3)
				sm.Push(4)
				sm.Push(-5)

				return args{sm: sm}
			}(),
			want: -5,
		},
		{
			name: "stack with popped values",
			args: func() args {
				sm := utils.NewStackMin(5)
				sm.Push(5)
				sm.Push(4)
				sm.Push(3)
				sm.Push(2)
				sm.Push(1)
				sm.Pop()
				sm.Pop()
				sm.Pop()
				return args{sm: sm}
			}(),
			want: 4,
		},
		{
			name: "stack with mix of pushed and popped values",
			args: func() args {
				sm := utils.NewStackMin(5)
				for i := -5; i < 0; i++ {
					sm.Push(i)
				}
				for i := 0; i < 3; i++ {
					sm.Pop()
				}

				sm.Push(-6)
				return args{sm: sm}
			}(),
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StackMin(tt.args.sm); got != tt.want {
				t.Errorf("StackMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
