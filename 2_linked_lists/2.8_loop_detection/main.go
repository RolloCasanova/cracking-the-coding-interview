package main

import (
	"os"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

type node = utils.Node[string]

func main() {
	if len(os.Args) < 4 {
		panic("usage: go run main.go <initialNode> <nodes...> <+> <loopingNode>")
	}

	// ensure "+" exists in the arguments only once
	plusIndex := -1
	for i, arg := range os.Args {
		if arg == "+" {
			if plusIndex != -1 {
				panic("\"+\" can only be used once")
			}

			plusIndex = i
		}
	}

	if plusIndex == -1 {
		panic("\"+\" is required")
	}

	// ensure there's only one looping node just after "+"
	if plusIndex+1 >= len(os.Args) {
		panic("looping node is required")
	}

	// if there are more nodes after the looping node, panic
	if plusIndex+2 < len(os.Args) {
		panic("there should be only one looping node")
	}

	// ensure linked list nodes are all different
	seen := make(map[string]bool)
	for _, arg := range os.Args[1:plusIndex] {
		if seen[arg] {
			panic("linked list nodes must be all different")
		}

		seen[arg] = true
	}

	// ensure looping node is in the linked list
	if !seen[os.Args[plusIndex+1]] {
		panic("looping node must be in the linked list")
	}

	ll := utils.ArrayToLinkedList(os.Args[1:plusIndex])
	ll.CreateLoop(os.Args[plusIndex+1])
	ll.PrintLoop("Linked List: ")

	loopNode := loopDetection(ll)
	if loopNode != nil {
		println("loop detected at node", loopNode.Value)
	} else {
		println("no loop detected")
	}
}

// loopDetection returns the node at the beginning of the loop in the linked list
// if no loop is found, it returns nil
func loopDetection(ll *node) *node {
	slow, fast := ll, ll

	// there's a loop guaranteed, so we don't need to check for fast.Next != nil
	slow = slow.Next
	fast = fast.Next.Next

	// soon or later, slow and fast will meet
	for slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// fast is going to be k nodes away from the beginning of the loop
	// and also k nodes away from the beginning of the linked list
	// so if we move slow to the beginning of the linked list and move both one node at a time
	// they will meet at the loop node
	slow = ll

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
