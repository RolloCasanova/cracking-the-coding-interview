package utils

import (
	"reflect"
	"testing"
)

func TestNewStackMin(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *stackMin
	}{
		{
			name: "new stack with capacity 0",
			args: args{
				capacity: 0,
			},
			want: &stackMin{
				values: make([]StackMinValue, 0),
			},
		},
		{
			name: "new stack with capacity 1",
			args: args{
				capacity: 1,
			},
			want: &stackMin{
				values: make([]StackMinValue, 0, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStackMin(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackMin_Push(t *testing.T) {
	type fields struct {
		values []StackMinValue
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
			name: "push to empty stack",
			fields: fields{
				values: []StackMinValue{},
			},
			args: args{
				value: 1,
			},
		},
		{
			name: "push non-min value to stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			args: args{
				value: 2,
			},
		},
		{
			name: "push min value to stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 2,
						min:   2,
					},
				},
			},
			args: args{
				value: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			sm.Push(tt.args.value)
		})
	}
}

func Test_stackMin_Pop(t *testing.T) {
	type fields struct {
		values []StackMinValue
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "pop from empty stack",
			fields: fields{
				values: []StackMinValue{},
			},
			want: 0,
		},
		{
			name: "pop from non-empty stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			if got := sm.Pop(); got != tt.want {
				t.Errorf("stackMin.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackMin_Peek(t *testing.T) {
	type fields struct {
		values []StackMinValue
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "peek from empty stack",
			fields: fields{
				values: []StackMinValue{},
			},
			want: 0,
		},
		{
			name: "peek from non-empty stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			if got := sm.Peek(); got != tt.want {
				t.Errorf("stackMin.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackMin_IsEmpty(t *testing.T) {
	type fields struct {
		values []StackMinValue
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty stack",
			fields: fields{
				values: []StackMinValue{},
			},
			want: true,
		},
		{
			name: "non-empty stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			if got := sm.IsEmpty(); got != tt.want {
				t.Errorf("stackMin.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackMin_IsFull(t *testing.T) {
	type fields struct {
		values []StackMinValue
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
			name: "full stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			args: args{
				capacity: 1,
			},
			want: true,
		},
		{
			name: "non-full stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			args: args{
				capacity: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			if got := sm.IsFull(tt.args.capacity); got != tt.want {
				t.Errorf("stackMin.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackMin_Print(t *testing.T) {
	type fields struct {
		values []StackMinValue
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
			name: "print empty stack",
			fields: fields{
				values: []StackMinValue{},
			},
			args: args{
				header: "stack: ",
			},
		},
		{
			name: "print non-empty stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			args: args{
				header: "stack: ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			sm.Print(tt.args.header)
		})
	}
}

func Test_stackMin_Min(t *testing.T) {
	type fields struct {
		values []StackMinValue
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "min from empty stack",
			fields: fields{
				values: []StackMinValue{},
			},
			want: 0,
		},
		{
			name: "min from non-empty stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
					{
						value: 3,
						min:   1,
					},
					{
						value: 2,
						min:   1,
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			if got := sm.Min(); got != tt.want {
				t.Errorf("stackMin.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackMin_Len(t *testing.T) {
	type fields struct {
		values []StackMinValue
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "length of empty stack",
			fields: fields{
				values: []StackMinValue{},
			},
			want: 0,
		},
		{
			name: "length of non-empty stack",
			fields: fields{
				values: []StackMinValue{
					{
						value: 1,
						min:   1,
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stackMin{
				values: tt.fields.values,
			}
			if got := sm.Len(); got != tt.want {
				t.Errorf("stackMin.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
