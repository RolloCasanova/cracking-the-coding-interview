package main

import "testing"

func Test_urlify(t *testing.T) {
	type args struct {
		s          string
		trueLength int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "an empty string should return an empty string",
			args: args{s: "", trueLength: 0},
			want: "",
		},
		{
			name: "a single space should return %20",
			args: args{s: " ", trueLength: 1},
			want: "%20",
		},
		{
			name: "a string with no spaces should return the same string",
			args: args{s: "abc", trueLength: 3},
			want: "abc",
		},
		{
			name: "a string with spaces should return the same string with %20 instead of spaces",
			args: args{s: "a b c", trueLength: 5},
			want: "a%20b%20c",
		},
		{
			name: "a string with trailing spaces at the end should return the same string with %20 instead of spaces, cut at the true length",
			args: args{s: "a b c    ", trueLength: 5},
			want: "a%20b%20c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlify(tt.args.s, tt.args.trueLength); got != tt.want {
				t.Errorf("urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}
