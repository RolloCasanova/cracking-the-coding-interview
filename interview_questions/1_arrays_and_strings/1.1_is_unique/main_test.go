package main

import "testing"

func Test_isUnique(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "empty string should be unique",
			args: args{s: ""},
			want: true,
		},
		{
			name: "string with unique characters should be unique",
			args: args{s: "abc"},
			want: true,
		},
		{
			name: "string with all characters should be unique",
			args: args{s: "abcdefghijklmnopqrstuvwxyz"},
			want: true,
		},
		{
			name:    "string with non-unique characters should not be unique",
			args:    args{s: "abca"},
			wantErr: true,
		},
		{
			name:    "string with more than 26 characters should not be unique",
			args:    args{s: "abcdefghijklmnopqrstuvwxyza"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isUnique(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("isUnique() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
