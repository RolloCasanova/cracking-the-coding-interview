package bst

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *BST
	}{
		{
			name: "no values",
			args: args{
				values: []int{},
			},
			want: nil,
		},
		{
			name: "one value",
			args: args{
				values: []int{1},
			},
			want: &BST{Value: 1},
		},
		{
			name: "two values",
			args: args{
				values: []int{1, 2},
			},
			want: &BST{
				Value: 1,
				Right: &BST{Value: 2},
			},
		},
		{
			name: "three values",
			args: args{

				values: []int{2, 1, 3},
			},
			want: &BST{
				Value: 2,
				Left:  &BST{Value: 1},
				Right: &BST{Value: 3},
			},
		},
		{
			name: "perfect tree of 7 values",
			args: args{
				values: []int{4, 2, 6, 1, 3, 5, 7},
			},
			want: &BST{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Insert(t *testing.T) {
	type fields struct {
		Value int
		Left  *BST
		Right *BST
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BST
	}{
		{
			name: "insert into nil tree",
			args: args{
				value: 1,
			},
			want: &BST{Value: 1},
		},
		{
			name: "insert into left",
			fields: fields{
				Value: 1,
			},
			args: args{
				value: 0,
			},
			want: &BST{
				Value: 1,
				Left:  &BST{Value: 0},
			},
		},
		{
			name: "insert into right",
			fields: fields{
				Value: 1,
			},
			args: args{
				value: 2,
			},
			want: &BST{
				Value: 1,
				Right: &BST{Value: 2},
			},
		},
		{
			name: "insert into left and right",
			fields: fields{
				Value: 3,
				Left:  &BST{Value: 1},
			},
			args: args{
				value: 2,
			},
			want: &BST{
				Value: 3,
				Left: &BST{
					Value: 1,
					Right: &BST{Value: 2},
				},
			},
		},
		{
			name: "insert into right and left",
			fields: fields{
				Value: 1,
				Right: &BST{Value: 3},
			},
			args: args{
				value: 2,
			},
			want: &BST{
				Value: 1,
				Right: &BST{
					Value: 3,
					Left:  &BST{Value: 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tr *BST
			if tt.fields != (fields{}) {
				tr = &BST{
					Value: tt.fields.Value,
					Left:  tt.fields.Left,
					Right: tt.fields.Right,
				}
			}
			if got := tr.Insert(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BST.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
