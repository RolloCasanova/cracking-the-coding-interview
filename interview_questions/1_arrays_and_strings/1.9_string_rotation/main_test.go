package main

import "testing"

func Test_stringRotation(t *testing.T) {
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
			name: "two empty string should return true",
			args: args{s1: "", s2: ""},
			want: true,
		},
		{
			name: "same strings should return true",
			args: args{s1: "abcde", s2: "abcde"},
			want: true,
		},
		{
			name: "a rotated string should return true",
			args: args{s1: "abcde", s2: "cdeab"},
			want: true,
		},
		{
			name: "a non-rotated string should return false",
			args: args{s1: "abcde", s2: "edcba"},
			want: false,
		},
		{
			name: "different length strings should return false",
			args: args{s1: "abcde", s2: "cdeabc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringRotation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("stringRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
