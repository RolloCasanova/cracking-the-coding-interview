package main

import (
	"testing"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

func Test_returnKthToLast(t *testing.T) {
	type args struct {
		ll *utils.Node
		k  int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "a linked list with one node and k = 1 should return the value of the node",
			args:    args{ll: &utils.Node{Value: "a", Next: nil}, k: 1},
			want:    "a",
			wantErr: false,
		},
		{
			name:    "a linked list with one node and k = 2 should return an error",
			args:    args{ll: &utils.Node{Value: "a", Next: nil}, k: 2},
			want:    "",
			wantErr: true,
		},
		{
			name:    "a linked list with two nodes and k = 1 should return the value of the last node",
			args:    args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: nil}}, k: 1},
			want:    "b",
			wantErr: false,
		},
		{
			name:    "a linked list with two nodes and k = 2 should return the value of the first node",
			args:    args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: nil}}, k: 2},
			want:    "a",
			wantErr: false,
		},
		{
			name:    "a linked list with three nodes and k = 1 should return the value of the last node",
			args:    args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "c", Next: nil}}}, k: 1},
			want:    "c",
			wantErr: false,
		},
		{
			name:    "a linked list with n nodes and k = n should return the value of the first node",
			args:    args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: &utils.Node{Value: "c", Next: nil}}}, k: 3},
			want:    "a",
			wantErr: false,
		},
		{
			name:    "a negative k should return an error",
			args:    args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: nil}}, k: -1},
			want:    "",
			wantErr: true,
		},
		{
			name:    "a value of k = 0 should return an error",
			args:    args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: nil}}, k: 0},
			want:    "",
			wantErr: true,
		},
		{
			name:    "a value of k greater than the length of the list should return an error",
			args:    args{ll: &utils.Node{Value: "a", Next: &utils.Node{Value: "b", Next: nil}}, k: 3},
			want:    "",
			wantErr: true,
		},
		{
			name:    "an empty linked list should return an error",
			args:    args{ll: nil, k: 1},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := returnKthToLast(tt.args.ll, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("returnKthToLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("returnKthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
