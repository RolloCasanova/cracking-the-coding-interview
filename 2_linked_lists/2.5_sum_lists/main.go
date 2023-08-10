package main

import (
	"os"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

type node = utils.Node[int]

func main() {
	// nodes must be separated by spaces and must be 1 digit long
	if len(os.Args) < 4 {
		panic("usage: go run ./main.go <ll1 - initialNode> <ll1 - nodes...> <+> <ll2 - initialNode> <ll2 - nodes...> ")
	}

	// separate the two linked lists by splitting the arguments at the "+"
	// get position of the "+" in the arguments; it only appears once, if not panic
	plusPosition := -1
	for i, arg := range os.Args {
		if arg == "+" {
			if plusPosition != -1 {
				panic("more than one \"+\" found in arguments")
			}
			plusPosition = i
		}
	}

	if plusPosition == -1 {
		panic("no \"+\" found in arguments")
	}

	// get the two linked lists from the arguments
	list1, err := utils.StringArrayToIntArray(os.Args[1:plusPosition])
	if err != nil {
		panic(err)
	}

	list2, err := utils.StringArrayToIntArray(os.Args[plusPosition+1:])
	if err != nil {
		panic(err)
	}

	// ensure all nodes in list1 are 1 digit long
	for _, num := range list1 {
		if num < 0 || num > 9 {
			panic("all nodes in list1 must be 1 digit long")
		}
	}

	// ensure all nodes in list2 are 1 digit long
	for _, num := range list2 {
		if num < 0 || num > 9 {
			panic("all nodes in list2 must be 1 digit long")
		}
	}

	ll1 := utils.ArrayToLinkedList[int](list1)
	ll2 := utils.ArrayToLinkedList[int](list2)

	ll1.Print("Linked list 1: ")
	ll2.Print("Linked list 2: ")

	sum := sumList(ll1, ll2)

	sum.Print("Sum (LSD on the left): ")

	sumForward := sumListForward(ll1, ll2)

	sumForward.Print("Sum (LSD on the right): ")

}

// sumList sums two linked lists, where each node contains a single digit
// the linked lists are in reverse order, so the first node is the least significant digit
func sumList(ll1, ll2 *node) *node {
	var (
		sum, result *node
		carry       bool
	)

	sum = &node{}
	result = sum

	for ll1 != nil || ll2 != nil {
		// if one of the linked lists is shorter than the other, pad it with 0s
		if ll1 == nil {
			ll1 = &node{}
		}

		if ll2 == nil {
			ll2 = &node{}
		}

		// sum the two nodes
		sum.Value = ll1.Value + ll2.Value

		// if there was a carry from the previous sum, add it to the current sum
		if carry {
			sum.Value++
		}

		// if the sum is greater than 10, set carry to true and subtract 10 from the sum
		if sum.Value >= 10 {
			carry = true
			sum.Value %= 10
		} else {
			carry = false
		}

		// move to the next nodes
		ll1 = ll1.Next
		ll2 = ll2.Next

		// if there are still nodes to sum, create a new node and move to it
		if ll1 != nil || ll2 != nil {
			sum.Next = &node{}
			sum = sum.Next
		}
	}

	if carry {
		sum.Next = &node{Value: 1}
	}

	return result
}

// sumListForward sums two linked lists, where each node contains a single digit
// the linked lists are in forward order, so the first node is the most significant digit
func sumListForward(ll1, ll2 *node) *node {
	// get the length of the two linked lists
	len1 := ll1.Len()
	len2 := ll2.Len()

	// if one of the linked lists is shorter than the other, pad it with 0s
	if len1 < len2 {
		ll1 = ll1.PadLinkedList(len2 - len1)
	} else if len2 < len1 {
		ll2 = ll2.PadLinkedList(len1 - len2)
	}

	// must sum the two linked lists in reverse order
	// so we need to reverse them first
	ll1 = ll1.Reverse()
	ll2 = ll2.Reverse()

	return sumList(ll1, ll2).Reverse()
}
