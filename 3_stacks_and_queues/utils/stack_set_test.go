package utils

import (
	"reflect"
	"testing"
)

func TestNewStackSet(t *testing.T) {
	type args struct {
		stackSize int
		numStacks int
	}
	tests := []struct {
		name string
		args args
		want StackSet
	}{
		{
			name: "new stack set with stack size 1 and 1 stack",
			args: args{
				stackSize: 1,
				numStacks: 1,
			},
			want: &stackSet{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
		},
		{
			name: "new stack set with stack size 1 and 2 stacks",
			args: args{
				stackSize: 1,
				numStacks: 2,
			},
			want: &stackSet{
				stacks: []*stack[int]{
					NewStack[int](1),
					NewStack[int](1),
				},
				stackSize: 1,
			},
		},
		{
			name: "new stack set with stack size 5 and 3 stacks",
			args: args{
				stackSize: 5,
				numStacks: 3,
			},
			want: &stackSet{
				stacks: []*stack[int]{
					NewStack[int](5),
					NewStack[int](5),
					NewStack[int](5),
				},
				stackSize: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStackSet(tt.args.stackSize, tt.args.numStacks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackSet_Push(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "push value onto empty stack set",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			args: args{
				value: 1,
			},
		},
		{
			name: "push value onto non-full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)

					return stacks
				}(),
				stackSize: 1,
			},
			args: args{
				value: 3,
			},
		},
		{
			name: "push value onto full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)
					stacks[2].Push(3)

					return stacks
				}(),
				stackSize: 1,
			},
			args: args{
				value: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			ss.Push(tt.args.value)
		})
	}
}

func Test_stackSet_Pop(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue int
	}{
		{
			name: "pop value from empty stack set",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			wantValue: 0,
		},
		{
			name: "pop value from non-full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)

					return stacks
				}(),
				stackSize: 1,
			},
			wantValue: 2,
		},
		{
			name: "pop value from full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)
					stacks[2].Push(3)

					return stacks
				}(),
				stackSize: 1,
			},
			wantValue: 3,
		},
		{
			name: "pop value from full stack set after popping from specfic stack",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)
					stacks[2].Push(3) // full stack set

					stacks[1].Pop() // popAt 1 action

					return stacks
				}(),
				stackSize: 1,
			},
			wantValue: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			if gotValue := ss.Pop(); gotValue != tt.wantValue {
				t.Errorf("stackSet.Pop() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func Test_stackSet_Peek(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue int
	}{
		{
			name: "peek value from empty stack set",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			wantValue: 0,
		},
		{
			name: "peek value from non-full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)

					return stacks
				}(),
				stackSize: 1,
			},
			wantValue: 2,
		},
		{
			name: "peek value from full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)
					stacks[2].Push(3)

					return stacks
				}(),
				stackSize: 1,
			},
			wantValue: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			if gotValue := ss.Peek(); gotValue != tt.wantValue {
				t.Errorf("stackSet.Peek() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func Test_stackSet_IsEmpty(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty stack set",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			want: true,
		},
		{
			name: "non-full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)

					return stacks
				}(),
				stackSize: 1,
			},
			want: false,
		},
		{
			name: "full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)
					stacks[2].Push(3)

					return stacks
				}(),
				stackSize: 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			if got := ss.IsEmpty(); got != tt.want {
				t.Errorf("stackSet.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackSet_IsFull(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	type args struct {
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "empty stack set",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			args: args{
				capacity: 1,
			},
			want: false,
		},
		{
			name: "non-full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)

					return stacks
				}(),
				stackSize: 1,
			},
			args: args{
				capacity: 1,
			},
			want: false,
		},
		{
			name: "full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}
					stacks[0].Push(1)
					stacks[1].Push(2)
					stacks[2].Push(3)

					return stacks
				}(),
				stackSize: 1,
			},
			args: args{
				capacity: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			if got := ss.IsFull(tt.args.capacity); got != tt.want {
				t.Errorf("stackSet.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackSet_Print(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	type args struct {
		header string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "print empty stack set",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			args: args{
				header: "stack set: ",
			},
		},
		{
			name: "print non-full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}

					return stacks
				}(),
				stackSize: 1,
			},
			args: args{
				header: "stack set: ",
			},
		},
		{
			name: "print full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
						stacks[i].Push(i + 1)
					}

					return stacks
				}(),
				stackSize: 1,
			},
			args: args{
				header: "stack set: ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			ss.Print(tt.args.header)
		})
	}
}

func Test_stackSet_Len(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "stack set of size 1",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			want: 1,
		},
		{
			name: "stack set of size 5",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](5),
				},
				stackSize: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			if got := ss.Len(); got != tt.want {
				t.Errorf("stackSet.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackSet_PopAt(t *testing.T) {
	type fields struct {
		stacks    []*stack[int]
		stackSize int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue int
		wantErr   bool
	}{
		{
			name: "pop value from empty stack set",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			args: args{
				index: 0,
			},
			wantValue: 0,
			wantErr:   true,
		},
		{
			name: "pop value from stack set with non-existing index",
			fields: fields{
				stacks: []*stack[int]{
					NewStack[int](1),
				},
				stackSize: 1,
			},
			args: args{
				index: 10,
			},
			wantValue: 0,
			wantErr:   true,
		},
		{
			name: "pop value from non-full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
					}

					stacks[0].Push(1)
					stacks[1].Push(2)

					return stacks
				}(),
			},
			args: args{
				index: 1,
			},
			wantValue: 2,
			wantErr:   false,
		},
		{
			name: "pop value from full stack set",
			fields: fields{
				stacks: func() []*stack[int] {
					stacks := make([]*stack[int], 3)
					for i := range stacks {
						stacks[i] = NewStack[int](1)
						stacks[i].Push(i + 1)
					}

					return stacks
				}(),
			},
			args: args{
				index: 2,
			},
			wantValue: 3,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &stackSet{
				stacks:    tt.fields.stacks,
				stackSize: tt.fields.stackSize,
			}
			gotValue, err := ss.PopAt(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("stackSet.PopAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("stackSet.PopAt() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
