package main

import (
	"reflect"
	"testing"
)

func Test_sumList(t *testing.T) {
	type args struct {
		ll1 *node
		ll2 *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "1 + 1 = 2",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 1},
			},
			want: &node{Value: 2},
		},
		{
			name: "1 + 9 = 10",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 9},
			},
			want: &node{Value: 0, Next: &node{Value: 1}},
		},
		{
			name: "1 + 99 = 100",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 9, Next: &node{Value: 9}},
			},
			want: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 1}}},
		},
		{
			name: "99 + 1 = 100",
			args: args{
				ll1: &node{Value: 9, Next: &node{Value: 9}},
				ll2: &node{Value: 1},
			},
			want: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 1}}},
		},
		{
			name: "1 + 999999 = 1000000",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9}}}}}},
			},
			want: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 1}}}}}}},
		},
		{
			name: "123 + 456 = 579",
			args: args{
				ll1: &node{Value: 3, Next: &node{Value: 2, Next: &node{Value: 1}}},
				ll2: &node{Value: 6, Next: &node{Value: 5, Next: &node{Value: 4}}},
			},
			want: &node{Value: 9, Next: &node{Value: 7, Next: &node{Value: 5}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumList(tt.args.ll1, tt.args.ll2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumListForward(t *testing.T) {
	type args struct {
		ll1 *node
		ll2 *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "1 + 1 = 2",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 1},
			},
			want: &node{Value: 2},
		},
		{
			name: "1 + 9 = 10",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 9},
			},
			want: &node{Value: 1, Next: &node{Value: 0}},
		},
		{
			name: "1 + 99 = 100",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 9, Next: &node{Value: 9}},
			},
			want: &node{Value: 1, Next: &node{Value: 0, Next: &node{Value: 0}}},
		},
		{
			name: "99 + 1 = 100",
			args: args{
				ll1: &node{Value: 9, Next: &node{Value: 9}},
				ll2: &node{Value: 1},
			},
			want: &node{Value: 1, Next: &node{Value: 0, Next: &node{Value: 0}}},
		},
		{
			name: "1 + 999999 = 1000000",
			args: args{
				ll1: &node{Value: 1},
				ll2: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9, Next: &node{Value: 9}}}}}},
			},
			want: &node{Value: 1, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0, Next: &node{Value: 0}}}}}}},
		},
		{
			name: "123 + 456 = 579",
			args: args{
				ll1: &node{Value: 3, Next: &node{Value: 2, Next: &node{Value: 1}}},
				ll2: &node{Value: 6, Next: &node{Value: 5, Next: &node{Value: 4}}},
			},
			want: &node{Value: 9, Next: &node{Value: 7, Next: &node{Value: 5}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumListForward(tt.args.ll1, tt.args.ll2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumListForward() = %v, want %v", got, tt.want)
			}
		})
	}
}
