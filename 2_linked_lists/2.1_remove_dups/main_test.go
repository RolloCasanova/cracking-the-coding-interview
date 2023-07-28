package main

import (
	"reflect"
	"testing"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

func Test_removeDupsWithMap(t *testing.T) {
	type args struct {
		ll *utils.Node
	}
	tests := []struct {
		name string
		args args
		want *utils.Node
	}{
		{
			name: "empty linked list should return nil",
			args: args{ll: nil},
			want: nil,
		},
		{
			name: "linked list with no duplicates should return the same linked list",
			args: args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "c", Next: nil}}}},
			want: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "c", Next: nil}}},
		},
		{
			name: "linked list with duplicates should return a linked list with duplicates removed",
			args: args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "a", Next: nil}}}},
			want: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: nil}},
		},
		{
			name: "linked list with all duplicates should return a linked list with only one node",
			args: args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "a", Next: &utils.Node{Value: "a", Next: nil}}}},
			want: &utils.Node{Value: "a", Next: nil},
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
		ll *utils.Node
	}
	tests := []struct {
		name string
		args args
		want *utils.Node
	}{
		{
			name: "empty linked list should return nil",
			args: args{ll: nil},
			want: nil,
		},
		{
			name: "linked list with no duplicates should return the same linked list",
			args: args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "c", Next: nil}}}},
			want: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "c", Next: nil}}},
		},
		{
			name: "linked list with duplicates should return a linked list with duplicates removed",
			args: args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "a", Next: nil}}}},
			want: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: nil}},
		},
		{
			name: "linked list with all duplicates should return a linked list with only one node",
			args: args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "a", Next: &utils.Node{Value: "a", Next: nil}}}},
			want: &utils.Node{Value: "a", Next: nil},
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
