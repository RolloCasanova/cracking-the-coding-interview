package main

import "testing"

func Test_oneAway(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "two empty strings are zero edits away",
			args: args{s1: "", s2: ""},
			want: true,
		},
		{
			name: "an empty string and a one length non-empty string are one edit away",
			args: args{s1: "", s2: "a"},
			want: true,
		},
		{
			name: "two distinct one-character strings are one edit away",
			args: args{s1: "a", s2: "b"},
			want: true,
		},
		{
			name: "two equal strings are zero edits away",
			args: args{s1: "abcde", s2: "abcde"},
			want: true,
		},
		{
			name: "two same length strings with one character difference are one edit away",
			args: args{s1: "abcde", s2: "ab de"},
			want: true,
		},
		{
			name: "two one-length difference strings with a different character are one edit away",
			args: args{s1: "abcde", s2: "abc de"},
			want: true,
		},
		{
			name: "two strings with two or more characters difference are more than one edit away",
			args: args{s1: "abc", s2: "abcde"},
			want: false,
		},
		{
			name: "two same length strings with more than one character difference are more than one edit away",
			args: args{s1: "abcde", s2: "abcyz"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneAway(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("oneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
