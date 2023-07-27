package main

import (
	"testing"
)

func Test_isPalindromePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "an empty string should be a palindrome permutation",
			args: args{s: ""},
			want: true,
		},
		{
			name: "a string with an empty space should be a palindrome permutation",
			args: args{s: " "},
			want: true,
		},
		{
			name: "a string with only one character should be a palindrome permutation",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "a string with only one character and an empty space should be a palindrome permutation",
			args: args{s: " a"},
			want: true,
		},
		{
			name: "a string with all-even number of characters should be a palindrome permutation",
			args: args{s: "abba"},
			want: true,
		},
		{
			name: "a string with all-even number of characters and an empty space should be a palindrome permutation",
			args: args{s: "ab ba"},
			want: true,
		},
		{
			name: "a string with all-even number of characters and only one odd number of characters should be a palindrome permutation",
			args: args{s: "abbac"},
			want: true,
		},
		{
			name: "a string with more that two odd number of characters should not be a palindrome permutation",
			args: args{s: "abbc"},
			want: false,
		},
		{
			name: "a string with more that two odd number of characters and an empty space should not be a palindrome permutation",
			args: args{s: "ab bc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromePermutation(tt.args.s); got != tt.want {
				t.Errorf("isPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
