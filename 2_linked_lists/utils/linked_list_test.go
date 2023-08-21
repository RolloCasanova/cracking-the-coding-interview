package utils

import (
	"reflect"
	"testing"
)

func TestNewNode_String(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *Node[string]
	}{
		{
			name: "new node with string value",
			args: args{
				v: "A",
			},
			want: &Node[string]{
				Value: "A",
				Next:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode_Int(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		{
			name: "new node with int value",
			args: args{
				v: 1,
			},
			want: &Node[int]{
				Value: 1,
				Next:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToLinkedList_String(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want *Node[string]
	}{
		{
			name: "array to linked list no values",
			args: args{
				values: []string{},
			},
			want: nil,
		},
		{
			name: "array to linked list with string values",
			args: args{
				values: []string{"A", "B", "C"},
			},
			want: &Node[string]{
				Value: "A",
				Next: &Node[string]{
					Value: "B",
					Next: &Node[string]{
						Value: "C",
						Next:  nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToLinkedList(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToLinkedList_Int(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		{
			name: "array to linked list no values",
			args: args{
				values: []int{},
			},
			want: nil,
		},
		{
			name: "array to linked list with string values",
			args: args{
				values: []int{1, 2, 3},
			},
			want: &Node[int]{
				Value: 1,
				Next: &Node[int]{
					Value: 2,
					Next: &Node[int]{
						Value: 3,
						Next:  nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToLinkedList(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrint_String(t *testing.T) {
	type args struct {
		header string
	}
	type fields struct {
		ll *Node[string]
	}
	tests := []struct {
		name   string
		args   args
		fields fields
	}{
		{
			name: "print linked list with string values",
			args: args{
				header: "linked List: ",
			},
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			ll.Print(tt.args.header)
		})
	}
}

func TestPrint_Int(t *testing.T) {
	type args struct {
		header string
	}
	type fields struct {
		ll *Node[int]
	}
	tests := []struct {
		name   string
		args   args
		fields fields
	}{
		{
			name: "print linked list with string values",
			args: args{
				header: "linked List: ",
			},
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			ll.Print(tt.args.header)
		})
	}
}

func TestLen_String(t *testing.T) {
	type fields struct {
		ll *Node[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "linked list with 0 values",
			fields: fields{
				ll: nil,
			},
			want: 0,
		},
		{
			name: "linked list with 3 values",
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLen_Int(t *testing.T) {
	type fields struct {
		ll *Node[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "linked list with 0 values",
			fields: fields{
				ll: nil,
			},
			want: 0,
		},
		{
			name: "linked list with 3 values",
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLinkedList_String(t *testing.T) {
	type args struct {
		n int
	}
	type fields struct {
		ll *Node[string]
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		want   *Node[string]
	}{
		{
			name: "pad nil linked list with 0 nodes",
			args: args{
				n: 0,
			},
			fields: fields{
				ll: nil,
			},
			want: nil,
		},
		{
			name: "pad nil linked list with 1 node",
			args: args{
				n: 1,
			},
			fields: fields{
				ll: nil,
			},
			want: &Node[string]{
				Value: "",
				Next:  nil,
			},
		},
		{
			name: "pad nil linked list with 3 nodes",
			args: args{
				n: 3,
			},
			fields: fields{
				ll: nil,
			},
			want: &Node[string]{
				Value: "",
				Next: &Node[string]{
					Value: "",
					Next: &Node[string]{
						Value: "",
						Next:  nil,
					},
				},
			},
		},
		{
			name: "pad linked list with 0 nodes",
			args: args{
				n: 0,
			},
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
			want: &Node[string]{
				Value: "A",
				Next: &Node[string]{
					Value: "B",
					Next: &Node[string]{
						Value: "C",
						Next:  nil,
					},
				},
			},
		},
		{
			name: "pad linked list with 1 node",
			args: args{
				n: 1,
			},
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
			want: &Node[string]{
				Value: "",
				Next: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
		},
		{
			name: "pad linked list with 3 nodes",
			args: args{
				n: 3,
			},
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
			want: &Node[string]{
				Value: "",
				Next: &Node[string]{
					Value: "",
					Next: &Node[string]{
						Value: "",
						Next: &Node[string]{
							Value: "A",
							Next: &Node[string]{
								Value: "B",
								Next: &Node[string]{
									Value: "C",
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.PadLinkedList(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PadLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLinkedList_Int(t *testing.T) {
	type args struct {
		n int
	}
	type fields struct {
		ll *Node[int]
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		want   *Node[int]
	}{
		{
			name: "pad nil linked list with 0 nodes",
			args: args{
				n: 0,
			},
			fields: fields{
				ll: nil,
			},
			want: nil,
		},
		{
			name: "pad nil linked list with 1 node",
			args: args{
				n: 1,
			},
			fields: fields{
				ll: nil,
			},
			want: &Node[int]{
				Value: 0,
				Next:  nil,
			},
		},
		{
			name: "pad nil linked list with 3 nodes",
			args: args{
				n: 3,
			},
			fields: fields{
				ll: nil,
			},
			want: &Node[int]{
				Value: 0,
				Next: &Node[int]{
					Value: 0,
					Next: &Node[int]{
						Value: 0,
						Next:  nil,
					},
				},
			},
		},
		{
			name: "pad linked list with 0 nodes",
			args: args{
				n: 0,
			},
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: &Node[int]{
				Value: 1,
				Next: &Node[int]{
					Value: 2,
					Next: &Node[int]{
						Value: 3,
						Next:  nil,
					},
				},
			},
		},
		{
			name: "pad linked list with 1 node",
			args: args{
				n: 1,
			},
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: &Node[int]{
				Value: 0,
				Next: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
		},
		{
			name: "pad linked list with 3 nodes",
			args: args{
				n: 3,
			},
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: &Node[int]{
				Value: 0,
				Next: &Node[int]{
					Value: 0,
					Next: &Node[int]{
						Value: 0,
						Next: &Node[int]{
							Value: 1,
							Next: &Node[int]{
								Value: 2,
								Next: &Node[int]{
									Value: 3,
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.PadLinkedList(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PadLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse_String(t *testing.T) {
	type fields struct {
		ll *Node[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node[string]
	}{
		{
			name: "reverse linked list with 0 nodes",
			fields: fields{
				ll: nil,
			},
			want: nil,
		},
		{
			name: "reverse linked list with 1 node",
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next:  nil,
				},
			},
			want: &Node[string]{
				Value: "A",
				Next:  nil,
			},
		},
		{
			name: "reverse linked list with 3 nodes",
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
			want: &Node[string]{
				Value: "C",
				Next: &Node[string]{
					Value: "B",
					Next: &Node[string]{
						Value: "A",
						Next:  nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse_Int(t *testing.T) {
	type fields struct {
		ll *Node[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node[int]
	}{
		{
			name: "reverse linked list with 0 nodes",
			fields: fields{
				ll: nil,
			},
			want: nil,
		},
		{
			name: "reverse linked list with 1 node",
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next:  nil,
				},
			},
			want: &Node[int]{
				Value: 1,
				Next:  nil,
			},
		},
		{
			name: "reverse linked list with 3 nodes",
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: &Node[int]{
				Value: 3,
				Next: &Node[int]{
					Value: 2,
					Next: &Node[int]{
						Value: 1,
						Next:  nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_String(t *testing.T) {
	type fields struct {
		ll *Node[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "string linked list with 0 nodes",
			fields: fields{
				ll: nil,
			},
			want: "",
		},
		{
			name: "string linked list with 1 node",
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next:  nil,
				},
			},
			want: "A",
		},
		{
			name: "string linked list with 3 nodes",
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
			want: "ABC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.Flat(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Int(t *testing.T) {
	type fields struct {
		ll *Node[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "string linked list with 0 nodes",
			fields: fields{
				ll: nil,
			},
			want: "",
		},
		{
			name: "string linked list with 1 node",
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next:  nil,
				},
			},
			want: "1",
		},
		{
			name: "string linked list with 3 nodes",
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll := tt.fields.ll
			if got := ll.Flat(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIntersectingLists_String(t *testing.T) {
	type args struct {
		ll1 *Node[string]
		ll2 *Node[string]
		ll3 *Node[string]
	}
	tests := []struct {
		name string
		args args
		want []*Node[string]
	}{
		{
			name: "ll1 and ll2 are nil",
			args: args{
				ll1: nil,
				ll2: nil,
				ll3: &Node[string]{
					Value: "C",
					Next:  nil,
				},
			},
			want: []*Node[string]{nil, nil},
		},
		{
			name: "ll1 is nil",
			args: args{
				ll1: nil,
				ll2: &Node[string]{
					Value: "B",
					Next:  nil,
				},
				ll3: &Node[string]{
					Value: "C",
					Next:  nil,
				},
			},
			want: []*Node[string]{nil, {Value: "B", Next: nil}},
		},
		{
			name: "ll2 is nil",
			args: args{
				ll1: &Node[string]{
					Value: "A",
					Next:  nil,
				},
				ll2: nil,
				ll3: &Node[string]{
					Value: "C",
					Next:  nil,
				},
			},
			want: []*Node[string]{{Value: "A", Next: nil}, nil},
		},
		{
			name: "intersection is nil",
			args: args{
				ll1: &Node[string]{
					Value: "A",
					Next:  nil,
				},
				ll2: &Node[string]{
					Value: "B",
					Next:  nil,
				},
				ll3: nil,
			},
			want: []*Node[string]{{Value: "A", Next: nil}, {Value: "B", Next: nil}},
		},
		{
			name: "single node lists and intersection",
			args: args{
				ll1: &Node[string]{
					Value: "A",
					Next:  nil,
				},
				ll2: &Node[string]{
					Value: "B",
					Next:  nil,
				},
				ll3: &Node[string]{
					Value: "C",
					Next:  nil,
				},
			},
			want: func() []*Node[string] {
				ll3 := &Node[string]{
					Value: "C",
					Next:  nil,
				}

				ll1 := &Node[string]{
					Value: "A",
					Next:  ll3,
				}
				ll2 := &Node[string]{
					Value: "B",
					Next:  ll3,
				}
				return []*Node[string]{ll1, ll2}
			}(),
		},
		{
			name: "multiple nodes lists and intersection",
			args: args{
				ll1: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next:  nil,
					},
				},
				ll2: &Node[string]{
					Value: "C",
					Next: &Node[string]{
						Value: "D",
						Next:  nil,
					},
				},
				ll3: &Node[string]{
					Value: "E",
					Next: &Node[string]{
						Value: "F",
						Next:  nil,
					},
				},
			},
			want: func() []*Node[string] {
				ll3 := &Node[string]{
					Value: "E",
					Next: &Node[string]{
						Value: "F",
						Next:  nil,
					},
				}

				ll1 := &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next:  ll3,
					},
				}
				ll2 := &Node[string]{
					Value: "C",
					Next: &Node[string]{
						Value: "D",
						Next:  ll3,
					},
				}

				return []*Node[string]{ll1, ll2}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll1 := tt.args.ll1
			ll2 := tt.args.ll2
			ll3 := tt.args.ll3
			if got1, got2 := CreateIntersectingLists(ll1, ll2, ll3); !reflect.DeepEqual(got1, tt.want[0]) || !reflect.DeepEqual(got2, tt.want[1]) {
				t.Errorf("CreateIntersectingLists() = %v, %v, want %v, %v", got1, got2, tt.want[0], tt.want[1])
			}
		})
	}
}

func TestCreateIntersectingLists_Int(t *testing.T) {
	type args struct {
		ll1 *Node[int]
		ll2 *Node[int]
		ll3 *Node[int]
	}
	tests := []struct {
		name string
		args args
		want []*Node[int]
	}{
		{
			name: "ll1 and ll2 are nil",
			args: args{
				ll1: nil,
				ll2: nil,
				ll3: &Node[int]{
					Value: 3,
					Next:  nil,
				},
			},
			want: []*Node[int]{nil, nil},
		},
		{
			name: "ll1 is nil",
			args: args{
				ll1: nil,
				ll2: &Node[int]{
					Value: 2,
					Next:  nil,
				},
				ll3: &Node[int]{
					Value: 3,
					Next:  nil,
				},
			},
			want: []*Node[int]{nil, {Value: 2, Next: nil}},
		},
		{
			name: "ll2 is nil",
			args: args{
				ll1: &Node[int]{
					Value: 1,
					Next:  nil,
				},
				ll2: nil,
				ll3: &Node[int]{
					Value: 3,
					Next:  nil,
				},
			},
			want: []*Node[int]{{Value: 1, Next: nil}, nil},
		},
		{
			name: "intersection is nil",
			args: args{
				ll1: &Node[int]{
					Value: 1,
					Next:  nil,
				},
				ll2: &Node[int]{
					Value: 2,
					Next:  nil,
				},
				ll3: nil,
			},
			want: []*Node[int]{{Value: 1, Next: nil}, {Value: 2, Next: nil}},
		},
		{
			name: "single node lists and intersection",
			args: args{
				ll1: &Node[int]{
					Value: 1,
					Next:  nil,
				},
				ll2: &Node[int]{
					Value: 2,
					Next:  nil,
				},
				ll3: &Node[int]{
					Value: 3,
					Next:  nil,
				},
			},
			want: func() []*Node[int] {
				ll3 := &Node[int]{
					Value: 3,
					Next:  nil,
				}

				ll1 := &Node[int]{
					Value: 1,
					Next:  ll3,
				}
				ll2 := &Node[int]{
					Value: 2,
					Next:  ll3,
				}
				return []*Node[int]{ll1, ll2}
			}(),
		},
		{
			name: "multiple nodes lists and intersection",
			args: args{
				ll1: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next:  nil,
					},
				},
				ll2: &Node[int]{
					Value: 3,
					Next: &Node[int]{
						Value: 4,
						Next:  nil,
					},
				},
				ll3: &Node[int]{
					Value: 5,
					Next: &Node[int]{
						Value: 6,
						Next:  nil,
					},
				},
			},
			want: func() []*Node[int] {
				ll3 := &Node[int]{
					Value: 5,
					Next: &Node[int]{
						Value: 6,
						Next:  nil,
					},
				}

				ll1 := &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next:  ll3,
					},
				}
				ll2 := &Node[int]{
					Value: 3,
					Next: &Node[int]{
						Value: 4,
						Next:  ll3,
					},
				}

				return []*Node[int]{ll1, ll2}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			ll1 := tt.args.ll1
			ll2 := tt.args.ll2
			ll3 := tt.args.ll3
			if got1, got2 := CreateIntersectingLists(ll1, ll2, ll3); !reflect.DeepEqual(got1, tt.want[0]) || !reflect.DeepEqual(got2, tt.want[1]) {
				t.Errorf("CreateIntersectingLists() = %v, %v, want %v, %v", got1, got2, tt.want[0], tt.want[1])
			}
		})
	}
}

func TestCreateLoop_String(t *testing.T) {
	type args struct {
		v string
	}
	type fields struct {
		ll *Node[string]
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		want   *Node[string]
	}{
		{
			name: "create loop in linked list with 0 nodes",
			args: args{
				v: "A",
			},
			fields: fields{
				ll: nil,
			},
			want: nil,
		},
		{
			name: "create loop in linked list with 1 node",
			args: args{
				v: "A",
			},
			fields: fields{
				ll: &Node[string]{
					Value: "A",
				},
			},
			want: func() *Node[string] {
				ll := &Node[string]{
					Value: "A",
				}
				ll.Next = ll

				return ll
			}(),
		},
		{
			name: "create loop in linked list with 3 nodes",
			args: args{
				v: "B",
			},
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				},
			},
			want: func() *Node[string] {
				ll := &Node[string]{
					Value: "A",
					Next: &Node[string]{
						Value: "B",
						Next: &Node[string]{
							Value: "C",
							Next:  nil,
						},
					},
				}
				// create loop from C to B
				ll.Next.Next.Next = ll.Next

				return ll
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			v := tt.args.v
			ll := tt.fields.ll

			// we are creating a loop, no need to deep equal check for the value
			// just check if the loop is created by checking the next nodes for a fixed ammount of times
			// if the loop is created, the next node will never be nil

			got := ll.CreateLoop(v)

			if tt.want == nil {
				if got != nil {
					t.Errorf("CreateLoop() = %v, want %v", got, tt.want)
				}
				return
			}

			for i := 0; i < 1000; i++ {
				if got == nil {
					t.Errorf("CreateLoop() = %v, want %v", got, tt.want)
				}
				got = got.Next
			}
		})
	}
}

func TestCreateLoop_Int(t *testing.T) {
	type args struct {
		v int
	}
	type fields struct {
		ll *Node[int]
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		want   *Node[int]
	}{
		{
			name: "create loop in linked list with 0 nodes",
			args: args{
				v: 1,
			},
			fields: fields{
				ll: nil,
			},
			want: nil,
		},
		{
			name: "create loop in linked list with 1 node",
			args: args{
				v: 1,
			},
			fields: fields{
				ll: &Node[int]{
					Value: 1,
				},
			},
			want: func() *Node[int] {
				ll := &Node[int]{
					Value: 1,
				}
				ll.Next = ll

				return ll
			}(),
		},
		{
			name: "create loop in linked list with 3 nodes",
			args: args{
				v: 1,
			},
			fields: fields{
				ll: &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				},
			},
			want: func() *Node[int] {
				ll := &Node[int]{
					Value: 1,
					Next: &Node[int]{
						Value: 2,
						Next: &Node[int]{
							Value: 3,
							Next:  nil,
						},
					},
				}
				// create loop from 3 to 1
				ll.Next.Next.Next = ll

				return ll
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			v := tt.args.v
			ll := tt.fields.ll

			// we are creating a loop, no need to deep equal check for the value
			// just check if the loop is created by checking the next nodes for a fixed ammount of times
			// if the loop is created, the next node will never be nil

			got := ll.CreateLoop(v)

			if tt.want == nil {
				if got != nil {
					t.Errorf("CreateLoop() = %v, want %v", got, tt.want)
				}
				return
			}

			for i := 0; i < 1000; i++ {
				if got == nil {
					t.Errorf("CreateLoop() = %v, want %v", got, tt.want)
				}
				got = got.Next
			}
		})
	}
}

func TestPrintLoop_String(t *testing.T) {
	type args struct {
		header string
	}
	type fields struct {
		ll *Node[string]
	}
	tests := []struct {
		name   string
		args   args
		fields fields
	}{
		{
			name: "print loop in linked list with 0 nodes",
			args: args{
				header: "linked list: ",
			},
			fields: fields{
				ll: nil,
			},
		},
		{
			name: "print loop in linked list with 1 node",
			args: args{
				header: "linked list: ",
			},
			fields: fields{
				ll: &Node[string]{
					Value: "A",
					Next:  nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			header := tt.args.header
			ll := tt.fields.ll
			ll.PrintLoop(header)
		})
	}
}
