package main

import (
	"os"

	"github.com/RolloCasanova/ctci-go/2_linked_lists/utils"
)

type node = utils.Node[string]

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <nodes...>")
	}

	ll := utils.ArrayToLinkedList[string](os.Args[1:])

	ll.Print("Original linked list: ")

	removeDupsWithMap(ll)
	ll.Print("Linked list with duplicates remove (map approach): ")

	removeDupsWithDoublePointer(ll)
	ll.Print("Linked list with duplicates remove (double pointer approach): ")
}

// removeDupsWithMap removes duplicates from a linked list using a map approach
func removeDupsWithMap(ll *node) *node {
	if ll == nil {
		return nil
	}

	root := ll
	m := make(map[string]bool)

	m[ll.Value] = true
	for ll.Next != nil {
		if m[ll.Next.Value] {
			ll.Next = ll.Next.Next
		} else {
			m[ll.Next.Value] = true
			ll = ll.Next
		}
	}

	return root
}

// removeDupsWithDoublePointer removes duplicates from a linked list using a double pointer approach
func removeDupsWithDoublePointer(ll *node) *node {
	if ll == nil {
		return nil
	}

	root := ll

	// use two pointers to iterate through the linked list and remove duplicates
	for ll != nil {
		runner := ll

		for runner.Next != nil {
			if runner.Next.Value == ll.Value {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}

		ll = ll.Next
	}

	return root
}
