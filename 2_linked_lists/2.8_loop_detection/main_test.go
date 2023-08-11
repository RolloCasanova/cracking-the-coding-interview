package main

import (
	"reflect"
	"testing"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

func auxGenerateCircularLinkedList(v []string, loopNodeVal string) (*node, *node) {
	ll := utils.ArrayToLinkedList(v)
	wantNode := ll.CreateLoop(loopNodeVal)

	return ll, wantNode
}

func Test_loopDetection(t *testing.T) {
	type args struct {
		ll *node
	}
	type test struct {
		name string
		args args
		want *node
	}
	tests := []test{
		func() test {
			ll, want := auxGenerateCircularLinkedList([]string{"A"}, "A")
			return test{
				name: "looped single node linked list",
				args: args{ll: ll},
				want: want,
			}
		}(),
		func() test {
			ll, want := auxGenerateCircularLinkedList([]string{"A", "B", "C", "D", "E"}, "A")
			return test{
				name: "looped linked list, loop at the beginning",
				args: args{ll: ll},
				want: want,
			}
		}(),
		func() test {
			ll, want := auxGenerateCircularLinkedList([]string{"A", "B", "C", "D", "E"}, "C")
			return test{
				name: "looped linked list, loop at the middle",
				args: args{ll: ll},
				want: want,
			}
		}(),
		func() test {
			ll, want := auxGenerateCircularLinkedList([]string{"A", "B", "C", "D", "E"}, "E")
			return test{
				name: "looped linked list, loop at the end",
				args: args{ll: ll},
				want: want,
			}
		}(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loopDetection(tt.args.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loopDetection() = %v, want %v", got, tt.want)
			}
		})
	}
}
