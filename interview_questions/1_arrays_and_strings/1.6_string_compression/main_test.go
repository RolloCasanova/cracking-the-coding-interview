package main

import "testing"

func Test_stringCompression(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "an empty string should return an empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "a string with one character should return the same string",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "a string with two equal characters should return the same string",
			args: args{s: "aa"},
			want: "aa",
		},
		{
			name: "a string with more than two same characters should return the compressed string",
			args: args{s: "aaa"},
			want: "a3",
		},
		{
			name: "a string with two different characters should return the same string",
			args: args{s: "ab"},
			want: "ab",
		},
		{
			name: "a string with three different characters should return the same string",
			args: args{s: "abc"},
			want: "abc",
		},
		{
			name: "a string with two same characters and one different should return the same string",
			args: args{s: "aab"},
			want: "aab",
		},
		{
			name: "a string with two same characters and other two same characters should return the same string",
			args: args{s: "aabb"},
			want: "aabb",
		},
		{
			name: "a really long string of the same character should return the compressed string",
			args: args{s: "aaaaaaaaaa"},
			want: "a10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringCompression(tt.args.s); got != tt.want {
				t.Errorf("stringCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
