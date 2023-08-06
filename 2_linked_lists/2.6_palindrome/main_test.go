package main

import (
	"testing"
)

func Test_isPalindromeReverse(t *testing.T) {
	type args struct {
		ll *node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a single node is a palindrome",
			args: args{ll: &node{Value: "a"}},
			want: true,
		},
		{
			name: "a two-character palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "a"}}},
			want: true,
		},
		{
			name: "a three-character palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a"}}}},
			want: true,
		},
		{
			name: "a really long odd palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: &node{Value: "d", Next: &node{Value: "c", Next: &node{Value: "b", Next: &node{Value: "a"}}}}}}}},
			want: true,
		},
		{
			name: "a really long even palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: &node{Value: "d", Next: &node{Value: "d", Next: &node{Value: "c", Next: &node{Value: "b", Next: &node{Value: "a"}}}}}}}}},
			want: true,
		},
		{
			name: "a non-palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c"}}}},
			want: false,
		},
		{
			name: "almost a palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a", Next: &node{Value: "a"}}}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeReverse(tt.args.ll); got != tt.want {
				t.Errorf("isPalindromeReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeStack(t *testing.T) {
	type args struct {
		ll *node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a single node is a palindrome",
			args: args{ll: &node{Value: "a"}},
			want: true,
		},
		{
			name: "a two-character palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "a"}}},
			want: true,
		},
		{
			name: "a three-character palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a"}}}},
			want: true,
		},
		{
			name: "a really long odd palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: &node{Value: "d", Next: &node{Value: "c", Next: &node{Value: "b", Next: &node{Value: "a"}}}}}}}},
			want: true,
		},
		{
			name: "a really long even palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: &node{Value: "d", Next: &node{Value: "d", Next: &node{Value: "c", Next: &node{Value: "b", Next: &node{Value: "a"}}}}}}}}},
			want: true,
		},
		{
			name: "a non-palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c"}}}},
			want: false,
		},
		{
			name: "almost a palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a", Next: &node{Value: "a"}}}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeStack(tt.args.ll); got != tt.want {
				t.Errorf("isPalindromeStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeRecursion(t *testing.T) {
	type args struct {
		ll *node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a single node is a palindrome",
			args: args{ll: &node{Value: "a"}},
			want: true,
		},
		{
			name: "a two-character palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "a"}}},
			want: true,
		},
		{
			name: "a three-character palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a"}}}},
			want: true,
		},
		{
			name: "a really long odd palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: &node{Value: "d", Next: &node{Value: "c", Next: &node{Value: "b", Next: &node{Value: "a"}}}}}}}},
			want: true,
		},
		{
			name: "a really long even palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c", Next: &node{Value: "d", Next: &node{Value: "d", Next: &node{Value: "c", Next: &node{Value: "b", Next: &node{Value: "a"}}}}}}}}},
			want: true,
		},
		{
			name: "a non-palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "c"}}}},
			want: false,
		},
		{
			name: "almost a palindrome",
			args: args{ll: &node{Value: "a", Next: &node{Value: "b", Next: &node{Value: "a", Next: &node{Value: "a"}}}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeRecursion(tt.args.ll); got != tt.want {
				t.Errorf("isPalindromeRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
