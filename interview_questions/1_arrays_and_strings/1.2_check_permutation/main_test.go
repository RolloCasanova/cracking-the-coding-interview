package main

import "testing"

func Test_isPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "empty strings are permutations",
			args:    args{s1: "", s2: ""},
			want:    true,
			wantErr: false,
		},
		{
			name:    "same single char strings are permutations",
			args:    args{s1: "a", s2: "a"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "same strings are permutations",
			args:    args{s1: "abc", s2: "abc"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "same strings with different order are permutations",
			args:    args{s1: "abc", s2: "cba"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "strings with different lengths are not permutations",
			args:    args{s1: "abc", s2: "ab"},
			want:    false,
			wantErr: true,
		},
		{
			name:    "string that differ by one character are not permutations",
			args:    args{s1: "abc", s2: "abd"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isPermutation(tt.args.s1, tt.args.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("isPermutation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
