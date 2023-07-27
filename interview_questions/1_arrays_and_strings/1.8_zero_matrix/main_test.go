package main

import (
	"reflect"
	"testing"
)

func Test_zeroMatrix(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "all 1's in a matrix should not change",
			args: args{
				matrix: [][]string{
					{"1", "1", "1"},
					{"1", "1", "1"},
					{"1", "1", "1"},
				},
			},
			want: [][]string{
				{"1", "1", "1"},
				{"1", "1", "1"},
				{"1", "1", "1"},
			},
		},
		{
			name: "all 0's in matrix should not change",
			args: args{
				matrix: [][]string{
					{"0", "0", "0"},
					{"0", "0", "0"},
					{"0", "0", "0"},
				},
			},
			want: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			},
		},
		{
			name: "a whole row of 0's will result in a whole 0's matrix",
			args: args{
				matrix: [][]string{
					{"1", "1", "1"},
					{"0", "0", "0"},
					{"1", "1", "1"},
				},
			},
			want: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			},
		},
		{
			name: "a whole column of 0's will result in a whole 0's matrix",
			args: args{
				matrix: [][]string{
					{"1", "0", "1"},
					{"1", "0", "1"},
					{"1", "0", "1"},
				},
			},
			want: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			},
		},
		{
			name: "a single 0 in a all 1's matrix will result in a row and column of 0's",
			args: args{
				matrix: [][]string{
					{"1", "1", "1"},
					{"1", "0", "1"},
					{"1", "1", "1"},
				},
			},
			want: [][]string{
				{"1", "0", "1"},
				{"0", "0", "0"},
				{"1", "0", "1"},
			},
		},
		{
			name: "a 0's diagonal in a all 1's matrix will result in a whole 0's matrix",
			args: args{
				matrix: [][]string{
					{"0", "1", "1", "1", "1", "1"},
					{"1", "0", "1", "1", "1", "1"},
					{"1", "1", "0", "1", "1", "1"},
					{"1", "1", "1", "0", "1", "1"},
					{"1", "1", "1", "1", "0", "1"},
					{"1", "1", "1", "1", "1", "0"},
				},
			},
			want: [][]string{
				{"0", "0", "0", "0", "0", "0"},
				{"0", "0", "0", "0", "0", "0"},
				{"0", "0", "0", "0", "0", "0"},
				{"0", "0", "0", "0", "0", "0"},
				{"0", "0", "0", "0", "0", "0"},
				{"0", "0", "0", "0", "0", "0"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
