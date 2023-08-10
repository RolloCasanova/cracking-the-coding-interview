package main

import (
	"testing"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

func Test_intersection(t *testing.T) {
	type args struct {
		ll1 *node
		ll2 *node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same length, intersecting",
			args: func(ll1, ll2, ll3 *node) args {
				ll1, ll2 = utils.CreateIntersectingLists(ll1, ll2, ll3)

				return args{ll1, ll2}
			}(utils.ArrayToLinkedList([]string{"A", "B", "C"}), utils.ArrayToLinkedList([]string{"D", "E", "F"}), utils.ArrayToLinkedList([]string{"X", "Y", "Z"})),
			want: true,
		},
		{
			name: "ll1 longer, intersecting",
			args: func(ll1, ll2, ll3 *node) args {
				ll1, ll2 = utils.CreateIntersectingLists(ll1, ll2, ll3)

				return args{ll1, ll2}
			}(utils.ArrayToLinkedList([]string{"A", "B", "C", "D"}), utils.ArrayToLinkedList([]string{"E", "F", "G"}), utils.ArrayToLinkedList([]string{"X", "Y", "Z"})),
			want: true,
		},
		{
			name: "ll2 longer, intersecting",
			args: func(ll1, ll2, ll3 *node) args {
				ll1, ll2 = utils.CreateIntersectingLists(ll1, ll2, ll3)

				return args{ll1, ll2}
			}(utils.ArrayToLinkedList([]string{"A", "B", "C"}), utils.ArrayToLinkedList([]string{"D", "E", "F", "G"}), utils.ArrayToLinkedList([]string{"X", "Y", "Z"})),
			want: true,
		},
		{
			name: "same length, not intersecting",
			args: args{
				ll1: utils.ArrayToLinkedList([]string{"A", "B", "C"}),
				ll2: utils.ArrayToLinkedList([]string{"D", "E", "F"}),
			},
			want: false,
		},
		{
			name: "ll1 longer, not intersecting",
			args: args{
				ll1: utils.ArrayToLinkedList([]string{"A", "B", "C", "D"}),
				ll2: utils.ArrayToLinkedList([]string{"E", "F", "G"}),
			},
			want: false,
		},
		{
			name: "ll2 longer, not intersecting",
			args: args{
				ll1: utils.ArrayToLinkedList([]string{"A", "B", "C"}),
				ll2: utils.ArrayToLinkedList([]string{"D", "E", "F", "G"}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.ll1, tt.args.ll2); got != tt.want {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
