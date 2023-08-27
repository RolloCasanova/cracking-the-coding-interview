package main

import (
	"reflect"
	"testing"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

func TestQueueViaStacks(t *testing.T) {
	type args struct {
		s1   stack
		s2   stack
		size int
	}
	tests := []struct {
		name string
		args args
		want *queueWithStacks
	}{
		{
			name: "nil stacks should return nil queueWithStacks",
			args: args{s1: nil, s2: nil, size: 0},
			want: nil,
		},
		{
			name: "size <= 0 should return nil queueWithStacks",
			args: args{s1: utils.NewStack[string](1), s2: utils.NewStack[string](1), size: 0},
			want: nil,
		},
		{
			name: "size > 0 should return queueWithStacks",
			args: args{s1: utils.NewStack[string](1), s2: utils.NewStack[string](1), size: 1},
			want: &queueWithStacks{s1: utils.NewStack[string](1), s2: utils.NewStack[string](1), size: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueueViaStacks(tt.args.s1, tt.args.s2, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueueViaStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queueWithStacks_enqueue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		qvs  *queueWithStacks
		args args
	}{
		{
			name: "enqueuing with stack s1 full should not panic",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](1)
				s2 := utils.NewStack[string](1)
				s1.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 1}
			}(),
			args: args{value: "B"},
		},
		{
			name: "enqueuing with stack s2 full should not panic",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](1)
				s2 := utils.NewStack[string](1)
				s2.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 1}
			}(),
			args: args{value: "B"},
		},
		{
			name: "enqueuing by not popping from s2 to s1 should succeed",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s1.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
			args: args{value: "B"},
		},
		{
			name: "enqueuing by popping from s2 to s1 should succeed",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s2.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
			args: args{value: "B"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.qvs.enqueue(tt.args.value)
		})
	}
}

func Test_queueWithStacks_dequeue(t *testing.T) {
	tests := []struct {
		name string
		qvs  *queueWithStacks
	}{
		{
			name: "dequeuing with both stacks empty should not panic",
			qvs:  &queueWithStacks{s1: utils.NewStack[string](1), s2: utils.NewStack[string](1), size: 1},
		},
		{
			name: "dequeuing with stack s1 empty should succeed",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s2.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
		},
		{
			name: "dequeuing with stack s2 empty should succeed",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s1.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.qvs.dequeue()
		})
	}
}

func Test_queueWithStacks_peek(t *testing.T) {
	tests := []struct {
		name string
		qvs  *queueWithStacks
	}{
		{
			name: "peeking with both stacks empty should not panic",
			qvs:  &queueWithStacks{s1: utils.NewStack[string](1), s2: utils.NewStack[string](1), size: 1},
		},
		{
			name: "peeking with stack s1 empty should succeed",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s2.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
		},
		{
			name: "peeking with stack s2 empty should succeed",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s1.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.qvs.peek()
		})
	}
}

func Test_queueWithStacks_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		qvs  *queueWithStacks
		want bool
	}{
		{
			name: "both stacks empty should return true",
			qvs:  &queueWithStacks{s1: utils.NewStack[string](1), s2: utils.NewStack[string](1), size: 1},
			want: true,
		},
		{
			name: "stack s1 empty should return false",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s2.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
			want: false,
		},
		{
			name: "stack s2 empty should return false",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s1.Push("A")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.qvs.isEmpty(); got != tt.want {
				t.Errorf("queueWithStacks.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queueWithStacks_print(t *testing.T) {
	tests := []struct {
		name string
		qvs  *queueWithStacks
	}{
		{
			name: "both stacks empty should print nothing",
			qvs:  &queueWithStacks{s1: utils.NewStack[string](1), s2: utils.NewStack[string](1), size: 1},
		},
		{
			name: "stack s1 empty should print stack s2",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s2.Push("A")
				s2.Push("B")
				s2.Push("C")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
		},
		{
			name: "stack s2 empty should pop values s1 to s2 and print stack s2",
			qvs: func() *queueWithStacks {
				s1 := utils.NewStack[string](3)
				s2 := utils.NewStack[string](3)
				s1.Push("A")
				s1.Push("B")
				s1.Push("C")

				return &queueWithStacks{s1: s1, s2: s2, size: 3}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.qvs.print()
		})
	}
}
