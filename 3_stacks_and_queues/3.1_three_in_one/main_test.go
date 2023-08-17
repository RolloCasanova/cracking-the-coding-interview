package main

import (
	"reflect"
	"testing"
)

func Test_threeInOne(t *testing.T) {
	type args struct {
		stackSize int
		stack1    []string
		stack2    []string
		stack3    []string
	}
	tests := []struct {
		name string
		args args
		want *threeStack
	}{
		{
			name: "size 0 stacks",
			args: args{
				stackSize: 0,
				stack1:    []string{},
				stack2:    []string{},
				stack3:    []string{},
			},
			want: &threeStack{
				values:  []string{},
				starts:  [3]int{0, 0, 0},
				ends:    [3]int{-1, -1, -1},
				indexes: [3]int{0, 0, 0},
			},
		},
		{
			name: "size 1 stacks",
			args: args{
				stackSize: 1,
				stack1:    []string{"1"},
				stack2:    []string{"2"},
				stack3:    []string{"3"},
			},
			want: &threeStack{
				values:  []string{"1", "2", "3"},
				starts:  [3]int{0, 1, 2},
				ends:    [3]int{0, 1, 2},
				indexes: [3]int{1, 2, 3},
			},
		},
		{
			name: "big stacks",
			args: args{
				stackSize: 5,
				stack1:    []string{"1", "2", "3", "4", "5"},
				stack2:    []string{"6", "7", "8", "9", "10"},
				stack3:    []string{"11", "12", "13", "14", "15"},
			},
			want: &threeStack{
				values:  []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"},
				starts:  [3]int{0, 5, 10},
				ends:    [3]int{4, 9, 14},
				indexes: [3]int{5, 10, 15},
			},
		},
		{
			name: "uneven stacks",
			args: args{
				stackSize: 3,
				stack1:    []string{},
				stack2:    []string{"4"},
				stack3:    []string{"7", "8", "9"},
			},
			want: &threeStack{
				values:  []string{"", "", "", "4", "", "", "7", "8", "9"},
				starts:  [3]int{0, 3, 6},
				ends:    [3]int{2, 5, 8},
				indexes: [3]int{0, 4, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeInOne(tt.args.stackSize, tt.args.stack1, tt.args.stack2, tt.args.stack3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeInOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
