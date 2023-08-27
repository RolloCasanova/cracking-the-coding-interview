package utils

import (
	"reflect"
	"testing"
)

func TestNewQueue_String(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *queue[string]
	}{
		{
			name: "create 0-length queue",
			args: args{
				capacity: 0,
			},
			want: &queue[string]{
				values: []string{},
				first:  0,
				last:   0,
			},
		},
		{
			name: "create 1-length queue",
			args: args{
				capacity: 1,
			},
			want: &queue[string]{
				values: []string{},
				first:  0,
				last:   0,
			},
		},
		{
			name: "create queue with multiple capacity",
			args: args{
				capacity: 3,
			},
			want: &queue[string]{
				values: []string{},
				first:  0,
				last:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue[string](tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewQueue_Int(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *queue[int]
	}{
		{
			name: "create 0-length queue",
			args: args{
				capacity: 0,
			},
			want: &queue[int]{
				values: []int{},
				first:  0,
				last:   0,
			},
		},
		{
			name: "create 1-length queue",
			args: args{
				capacity: 1,
			},
			want: &queue[int]{
				values: []int{},
				first:  0,
				last:   0,
			},
		},
		{
			name: "create queue with multiple capacity",
			args: args{
				capacity: 3,
			},
			want: &queue[int]{
				values: []int{},
				first:  0,
				last:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue[int](tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnqueue_String(t *testing.T) {
	type args struct {
		value string
	}
	type fields struct {
		q *queue[string]
	}
	tests := []struct {
		name   string
		args   args
		fiedls fields
		want   *queue[string]
	}{
		{
			name: "enqueing into a nil queue",
			args: args{
				value: "A",
			},
			fiedls: fields{
				q: nil,
			},
			want: nil,
		},
		{
			name: "enqueing into an empty queue",
			args: args{
				value: "A",
			},
			fiedls: fields{
				q: NewQueue[string](3),
			},
			want: func() *queue[string] {
				q := NewQueue[string](3)
				q.values = append(q.values, "A")
				q.first = 1

				return q
			}(),
		},
		{
			name: "enqueing into a queue with 1 element",
			args: args{
				value: "B",
			},
			fiedls: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A")
					q.first = 1

					return q
				}(),
			},
			want: func() *queue[string] {
				q := NewQueue[string](3)
				q.values = append(q.values, "A", "B")
				q.first = 2

				return q
			}(),
		},
		{
			name: "enqueing to fill a queue",
			args: args{
				value: "C",
			},
			fiedls: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B")
					q.first = 2

					return q
				}(),
			},
			want: func() *queue[string] {
				q := NewQueue[string](3)
				q.values = append(q.values, "A", "B", "C")
				q.first = 3

				return q
			}(),
		},
		{
			name: "enqueing to a full queue", // should not change the queue
			args: args{
				value: "D",
			},
			fiedls: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B", "C")
					q.first = 3

					return q
				}(),
			},
			want: func() *queue[string] {
				q := NewQueue[string](3)
				q.values = append(q.values, "A", "B", "C")
				q.first = 3

				return q
			}(),
		},
		{
			name: "enqueing to fill the queue and shift the values to the left",
			args: args{
				value: "D",
			},
			fiedls: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B", "C")
					q.first = 3

					// q.values = []string{"A", "B", "C"}
					// q.last               ^
					// q.first                        ^

					// dequeue "A" by moving the last index one position to the right
					q.last++

					// this will result with the same elements in the array, except the first index will be 1 (in "B"'s position)
					// q.values = []string{"A", "B", "C"}
					// q.last                    ^
					// q.first                        ^

					// after enqueing "D", the queue should be shifted to the left starting from first intex to make room for "D"
					// q.values = []string{"B", "C", "D"}
					// q.last               ^
					// q.first                        ^

					return q
				}(),
			},
			want: func() *queue[string] {
				q := NewQueue[string](3)
				q.values = append(q.values, "B", "C", "D")
				q.first = 3

				return q
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fiedls.q
			q.Enqueue(tt.args.value)
			if !reflect.DeepEqual(q, tt.want) {
				t.Errorf("Enqueue() = %v, want %v", q, tt.want)
			}
		})
	}
}

func TestEnqueue_Int(t *testing.T) {
	type args struct {
		value int
	}
	type fields struct {
		q *queue[int]
	}
	tests := []struct {
		name   string
		args   args
		fiedls fields
		want   *queue[int]
	}{
		{
			name: "enqueing into a nil queue",
			args: args{
				value: 1,
			},
			fiedls: fields{
				q: nil,
			},
			want: nil,
		},
		{
			name: "enqueing into an empty queue",
			args: args{
				value: 1,
			},
			fiedls: fields{
				q: NewQueue[int](3),
			},
			want: func() *queue[int] {
				q := NewQueue[int](3)
				q.values = append(q.values, 1)
				q.first = 1

				return q
			}(),
		},
		{
			name: "enqueing into a queue with 1 element",
			args: args{
				value: 2,
			},
			fiedls: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1)
					q.first = 1

					return q
				}(),
			},
			want: func() *queue[int] {
				q := NewQueue[int](3)
				q.values = append(q.values, 1, 2)
				q.first = 2

				return q
			}(),
		},
		{
			name: "enqueing to fill a queue",
			args: args{
				value: 3,
			},
			fiedls: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2)
					q.first = 2

					return q
				}(),
			},
			want: func() *queue[int] {
				q := NewQueue[int](3)
				q.values = append(q.values, 1, 2, 3)
				q.first = 3

				return q
			}(),
		},
		{
			name: "enqueing to a full queue", // should not change the queue
			args: args{
				value: 4,
			},
			fiedls: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2, 3)
					q.first = 3

					return q
				}(),
			},
			want: func() *queue[int] {
				q := NewQueue[int](3)
				q.values = append(q.values, 1, 2, 3)
				q.first = 3

				return q
			}(),
		},
		{
			name: "enqueing to fill the queue and shift the values to the left",
			args: args{
				value: 4,
			},
			fiedls: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2, 3)
					q.first = 3

					// q.values = []int{1, 2, 3}
					// q.last           ^
					// q.first                ^

					// dequeue 1 by moving the last index one position to the right
					q.last++

					// this will result with the same elements in the array, except the first index will be 1 (in 2's position)
					// q.values = []int{1, 2, 3}
					// q.last              ^
					// q.first                ^

					// after enqueing 4, the queue should be shifted to the left starting from first intex to make room for 4
					// q.values = []int{2, 3, 4}
					// q.last           ^
					// q.first                ^

					return q
				}(),
			},
			want: func() *queue[int] {
				q := NewQueue[int](3)
				q.values = append(q.values, 2, 3, 4)
				q.first = 3

				return q
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fiedls.q
			q.Enqueue(tt.args.value)
			if !reflect.DeepEqual(q, tt.want) {
				t.Errorf("Enqueue() = %v, want %v", q, tt.want)
			}
		})
	}
}

func TestDequeue_String(t *testing.T) {
	type fields struct {
		q *queue[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "dequeue from a nil queue",
			fields: fields{
				q: nil,
			},
			want: "",
		},
		{
			name: "dequeue from an empty queue",
			fields: fields{
				q: NewQueue[string](3),
			},
			want: "",
		},
		{
			name: "dequeue from a queue with 1 element",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A")
					q.first = 1

					return q
				}(),
			},
			want: "A",
		},
		{
			name: "dequeue from a full queue",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B", "C")
					q.first = 3

					return q
				}(),
			},
			want: "A",
		},
		{
			name: "dequeue from a full queue after dequeueing the first element",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B", "C")
					q.first = 3

					// simulate dequeueing "A" by moving the last index one position to the right
					q.last++

					return q
				}(),
			},
			want: "B",
		},
		{
			name: "dequeue from a full queue after dequeueing the first 2 elements in order to shift the values to the left",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B", "C")
					q.first = 3

					// simulate dequeueing "A" by moving the last index one position to the right
					q.last++

					// simulate dequeueing "B" by moving the last index one position to the right
					q.last++

					// simulate dequeueing "C" by moving the last index one position to the right
					q.last++

					// q.start and q.last are now equal, which means the queue is empty and these values will be reset back to 0

					return q
				}(),
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fields.q
			if got := q.Dequeue(); got != tt.want {
				t.Errorf("Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDequeue_Int(t *testing.T) {
	type fields struct {
		q *queue[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "dequeue from a nil queue",
			fields: fields{
				q: nil,
			},
			want: 0,
		},
		{
			name: "dequeue from an empty queue",
			fields: fields{
				q: NewQueue[int](3),
			},
			want: 0,
		},
		{
			name: "dequeue from a queue with 1 element",
			fields: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1)
					q.first = 1

					return q
				}(),
			},
			want: 1,
		},
		{
			name: "dequeue from a full queue",
			fields: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2, 3)
					q.first = 3

					return q
				}(),
			},
			want: 1,
		},
		{
			name: "dequeue from a full queue after dequeueing the first element",
			fields: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2, 3)
					q.first = 3

					// simulate dequeueing 1 by moving the last index one position to the right
					q.last++

					return q
				}(),
			},
			want: 2,
		},
		{
			name: "dequeue from a full queue after dequeueing the first 2 elements in order to shift the values to the left",
			fields: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2, 3)
					q.first = 3

					// simulate dequeueing 1 by moving the last index one position to the right
					q.last++

					// simulate dequeueing 2 by moving the last index one position to the right
					q.last++

					// simulate dequeueing 3 by moving the last index one position to the right
					q.last++

					// q.start and q.last are now equal, which means the queue is empty and these values will be reset back to 0

					return q
				}(),
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fields.q
			if got := q.Dequeue(); got != tt.want {
				t.Errorf("Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekQueue_String(t *testing.T) {
	type fields struct {
		q *queue[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "peek from a nil queue",
			fields: fields{
				q: nil,
			},
			want: "",
		},
		{
			name: "peek from an empty queue",
			fields: fields{
				q: NewQueue[string](3),
			},
			want: "",
		},
		{
			name: "peek from a queue with 1 element",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A")
					q.first = 1

					return q
				}(),
			},
			want: "A",
		},
		{
			name: "peek from a full queue",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B", "C")
					q.first = 3

					return q
				}(),
			},
			want: "A",
		},
		{
			name: "peek from a full queue after dequeueing the first element",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A", "B", "C")
					q.first = 3

					// simulate dequeueing "A" by moving the last index one position to the right
					q.last++

					return q
				}(),
			},
			want: "B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fields.q
			if got := q.Peek(); got != tt.want {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekQueue_Int(t *testing.T) {
	type fields struct {
		q *queue[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "peek from a nil queue",
			fields: fields{
				q: nil,
			},
			want: 0,
		},
		{
			name: "peek from an empty queue",
			fields: fields{
				q: NewQueue[int](3),
			},
			want: 0,
		},
		{
			name: "peek from a queue with 1 element",
			fields: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1)
					q.first = 1

					return q
				}(),
			},
			want: 1,
		},
		{
			name: "peek from a full queue",
			fields: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2, 3)
					q.first = 3

					return q
				}(),
			},
			want: 1,
		},
		{
			name: "peek from a full queue after dequeueing the first element",
			fields: fields{
				q: func() *queue[int] {
					q := NewQueue[int](3)
					q.values = append(q.values, 1, 2, 3)
					q.first = 3

					// simulate dequeueing 1 by moving the last index one position to the right
					q.last++

					return q
				}(),
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fields.q
			if got := q.Peek(); got != tt.want {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmptyQueue_String(t *testing.T) {
	type fields struct {
		q *queue[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "a nil queue is empty",
			fields: fields{
				q: nil,
			},
			want: true,
		},
		{
			name: "an empty queue is empty",
			fields: fields{
				q: NewQueue[string](3),
			},
			want: true,
		},
		{
			name: "a queue with 1 element is not empty",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A")
					q.first = 1

					return q
				}(),
			},
			want: false,
		},
		{
			name: "a full queue is not empty",
			fields: fields{
				q: func() *queue[string] {
					q := NewQueue[string](3)
					q.values = append(q.values, "A")
					q.first = 1

					return q
				}(),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fields.q
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
