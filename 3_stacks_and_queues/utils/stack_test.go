package utils

import (
	"reflect"
	"testing"
)

func TestNewStack_String(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *stack[string]
	}{
		{
			name: "new stack with capacity 0",
			args: args{
				capacity: 0,
			},
			want: &stack[string]{
				values: make([]string, 0),
			},
		},
		{
			name: "new stack with capacity 1",
			args: args{
				capacity: 1,
			},
			want: &stack[string]{
				values: make([]string, 0, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack[string](tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStack_Int(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *stack[int]
	}{
		{
			name: "new stack with capacity 0",
			args: args{
				capacity: 0,
			},
			want: &stack[int]{
				values: make([]int, 0),
			},
		},
		{
			name: "new stack with capacity 1",
			args: args{
				capacity: 1,
			},
			want: &stack[int]{
				values: make([]int, 0, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack[int](tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPush_String(t *testing.T) {
	tests := []struct {
		name  string
		stack *stack[string]
		value string
		want  *stack[string]
	}{
		{
			name: "push value to empty stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)

				return s
			}(),
			value: "1",
			want: &stack[string]{
				values: []string{"1"},
			},
		},
		{
			name: "push value with half full stack",
			stack: func() *stack[string] {
				s := NewStack[string](3)
				s.Push("1")

				return s
			}(),
			value: "2",
			want: &stack[string]{
				values: []string{"1", "2"},
			},
		},
		{
			name: "push value with non-full stack",
			stack: func() *stack[string] {
				s := NewStack[string](3)
				s.Push("1")
				s.Push("2")

				return s
			}(),
			value: "3",
			want: &stack[string]{
				values: []string{"1", "2", "3"},
			},
		},
		{
			name: "push value with full stack (will not push)",
			stack: func() *stack[string] {
				s := NewStack[string](3)
				s.Push("1")
				s.Push("2")
				s.Push("3")

				return s
			}(),
			value: "4",
			want: &stack[string]{
				values: []string{"1", "2", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Push(tt.value)
			if !reflect.DeepEqual(tt.stack, tt.want) {
				t.Errorf("Push() = %v, want %v", tt.stack, tt.want)
			}
		})
	}
}

func TestPush_Int(t *testing.T) {
	tests := []struct {
		name  string
		stack *stack[int]
		value int
		want  *stack[int]
	}{
		{
			name: "push value to empty stack",
			stack: func() *stack[int] {
				s := NewStack[int](1)

				return s
			}(),
			value: 1,
			want: &stack[int]{
				values: []int{1},
			},
		},
		{
			name: "push value with half full stack",
			stack: func() *stack[int] {
				s := NewStack[int](3)
				s.Push(1)

				return s
			}(),
			value: 2,
			want: &stack[int]{
				values: []int{1, 2},
			},
		},
		{
			name: "push value with non-full stack",
			stack: func() *stack[int] {
				s := NewStack[int](3)
				s.Push(1)
				s.Push(2)

				return s
			}(),
			value: 3,
			want: &stack[int]{
				values: []int{1, 2, 3},
			},
		},
		{
			name: "push value with full stack (will not push)",
			stack: func() *stack[int] {
				s := NewStack[int](3)
				s.Push(1)
				s.Push(2)
				s.Push(3)

				return s
			}(),
			value: 4,
			want: &stack[int]{
				values: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Push(tt.value)
			if !reflect.DeepEqual(tt.stack, tt.want) {
				t.Errorf("Push() = %v, want %v", tt.stack, tt.want)
			}
		})
	}
}

func TestPop_String(t *testing.T) {
	tests := []struct {
		name  string
		stack *stack[string]
		want  string
	}{
		{
			name: "pop value from 1 element stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)
				s.Push("1")

				return s
			}(),
			want: "1",
		},
		{
			name: "pop value from multiple elements stack",
			stack: func() *stack[string] {
				s := NewStack[string](3)
				s.Push("1")
				s.Push("2")
				s.Push("3")

				return s
			}(),
			want: "3",
		},
		{
			name: "pop value from empty stack", // will return zero value
			stack: func() *stack[string] {
				s := NewStack[string](3)

				return s
			}(),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPop_Int(t *testing.T) {
	tests := []struct {
		name  string
		stack *stack[int]
		want  int
	}{
		{
			name: "pop value from 1 element stack",
			stack: func() *stack[int] {
				s := NewStack[int](1)
				s.Push(1)

				return s
			}(),
			want: 1,
		},
		{
			name: "pop value from multiple elements stack",
			stack: func() *stack[int] {
				s := NewStack[int](3)
				s.Push(1)
				s.Push(2)
				s.Push(3)

				return s
			}(),
			want: 3,
		},
		{
			name: "pop value from empty stack", // will return zero value
			stack: func() *stack[int] {
				s := NewStack[int](3)

				return s
			}(),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeek_String(t *testing.T) {
	tests := []struct {
		name  string
		stack *stack[string]
		want  string
	}{
		{
			name: "peek value from 1 element stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)
				s.Push("1")

				return s
			}(),
			want: "1",
		},
		{
			name: "peek value from multiple elements stack",
			stack: func() *stack[string] {
				s := NewStack[string](3)
				s.Push("1")
				s.Push("2")
				s.Push("3")

				return s
			}(),
			want: "3",
		},
		{
			name: "peek value from empty stack", // will return zero value
			stack: func() *stack[string] {
				s := NewStack[string](3)

				return s
			}(),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Peek(); got != tt.want {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeek_Int(t *testing.T) {
	tests := []struct {
		name  string
		stack *stack[int]
		want  int
	}{
		{
			name: "peek value from 1 element stack",
			stack: func() *stack[int] {
				s := NewStack[int](1)
				s.Push(1)

				return s
			}(),
			want: 1,
		},
		{
			name: "peek value from multiple elements stack",
			stack: func() *stack[int] {
				s := NewStack[int](3)
				s.Push(1)
				s.Push(2)
				s.Push(3)

				return s
			}(),
			want: 3,
		},
		{
			name: "peek value from empty stack", // will return zero value
			stack: func() *stack[int] {
				s := NewStack[int](3)

				return s
			}(),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Peek(); got != tt.want {
				t.Errorf("PeekInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty_String(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack[string]
		want  bool
	}{
		{
			name: "empty stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)

				return s
			}(),
			want: true,
		},
		{
			name: "non-empty stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)
				s.Push("1")

				return s
			}(),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty_Int(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack[int]
		want  bool
	}{
		{
			name: "empty stack",
			stack: func() *stack[int] {
				s := NewStack[int](1)

				return s
			}(),
			want: true,
		},
		{
			name: "non-empty stack",
			stack: func() *stack[int] {
				s := NewStack[int](1)
				s.Push(1)

				return s
			}(),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFull_String(t *testing.T) {
	tests := []struct {
		name     string
		stack    *stack[string]
		capacity int
		want     bool
	}{
		{
			name: "full stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)
				s.Push("1")

				return s
			}(),
			capacity: 1,
			want:     true,
		},
		{
			name: "non-full stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)

				return s
			}(),
			capacity: 1,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.IsFull(tt.capacity); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFull_Int(t *testing.T) {
	tests := []struct {
		name     string
		stack    *stack[int]
		capacity int
		want     bool
	}{
		{
			name: "full stack",
			stack: func() *stack[int] {
				s := NewStack[int](1)
				s.Push(1)

				return s
			}(),
			capacity: 1,
			want:     true,
		},
		{
			name: "non-full stack",
			stack: func() *stack[int] {
				s := NewStack[int](1)

				return s
			}(),
			capacity: 1,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.IsFull(tt.capacity); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrint_String(t *testing.T) {
	tests := []struct {
		name   string
		stack  *stack[string]
		header string
	}{
		{
			name: "print empty stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)

				return s
			}(),
			header: "stack: ",
		},
		{
			name: "print non-empty stack",
			stack: func() *stack[string] {
				s := NewStack[string](1)
				s.Push("1")

				return s
			}(),
			header: "stack: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.stack.Print(tt.header)
		})
	}
}
