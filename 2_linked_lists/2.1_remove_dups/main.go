package main

import (
	"fmt"
	"os"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <val1> <val2> ... <valN>")
	}

	ll := utils.ArrayToLinkedList(os.Args[1:])

	fmt.Println("Original linked list: ")
	utils.PrintLinkedList(ll)

	fmt.Println("Linked list with duplicates remove (map approach): ")
	utils.PrintLinkedList(removeDupsWithMap(ll))

	fmt.Println("Linked list with duplicates remove (double pointer approach): ")
	utils.PrintLinkedList(removeDupsWithDoublePointer(ll))
}

func removeDupsWithMap(ll *utils.Node) *utils.Node {
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

func removeDupsWithDoublePointer(ll *utils.Node) *utils.Node {
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
