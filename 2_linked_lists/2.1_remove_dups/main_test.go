package main

import (
	"reflect"
	"testing"
)

func Test_removeDupsWithMap(t *testing.T) {
	type args struct {
		ll *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "empty linked list should return nil",
			args: args{ll: nil},
			want: nil,
		},
		{
			name: "linked list with no duplicates should return the same linked list",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: nil}}}},
			want: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: nil}}},
		},
		{
			name: "linked list with duplicates should return a linked list with duplicates removed",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a", Next: nil}}}},
			want: &node{Value: "a", Next: &node{Value: "b", Next: nil}},
		},
		{
			name: "linked list with all duplicates should return a linked list with only one node",
			args: args{ll: &node{Value: "a", Next: &node{Value: "a", Next: &node{Value: "a", Next: nil}}}},
			want: &node{Value: "a", Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDupsWithMap(tt.args.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDupsWithMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDupsWithDoublePointer(t *testing.T) {
	type args struct {
		ll *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "empty linked list should return nil",
			args: args{ll: nil},
			want: nil,
		},
		{
			name: "linked list with no duplicates should return the same linked list",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: nil}}}},
			want: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: nil}}},
		},
		{
			name: "linked list with duplicates should return a linked list with duplicates removed",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a", Next: nil}}}},
			want: &node{Value: "a", Next: &node{Value: "b", Next: nil}},
		},
		{
			name: "linked list with all duplicates should return a linked list with only one node",
			args: args{ll: &node{Value: "a", Next: &node{Value: "a", Next: &node{Value: "a", Next: nil}}}},
			want: &node{Value: "a", Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDupsWithDoublePointer(tt.args.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDupsWithDoublePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
