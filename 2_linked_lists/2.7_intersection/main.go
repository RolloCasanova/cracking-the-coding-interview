package main

import (
	"fmt"
	"os"

	"github.com/RolloCasanova/ctci-go/2_linked_lists/utils"
)

type node = utils.Node[string]

func main() {
	if len(os.Args) < 3 {
		panic("usage: go run main.go <ll1 nodes...> <-> <ll2 nodes...> <+> <intersection nodes...>")
	}

	// ensure - exists in the arguments only once
	minusIndex := -1
	for i, arg := range os.Args {
		if arg == "-" {
			if minusIndex != -1 {
				panic("\"-\" can only be used once")
			}

			minusIndex = i
		}
	}

	if minusIndex == -1 {
		panic("\"-\" is required")
	}

	// ensure + exists in the arguments only once
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

	// ensure minus comes before plus
	if minusIndex > plusIndex {
		panic("\"-\" must come before \"+\"")
	}

	ll1, ll2 := utils.CreateIntersectingLists(
		utils.ArrayToLinkedList(os.Args[1:minusIndex]),
		utils.ArrayToLinkedList(os.Args[minusIndex+1:plusIndex]),
		utils.ArrayToLinkedList(os.Args[plusIndex+1:]),
	)

	ll1.Print("Linked List 1")
	ll2.Print("Linked List 2")

	if intersection(ll1, ll2) {
		fmt.Print("intersection found\n\n")
	} else {
		fmt.Print("intersection not found\n\n")
	}

}

// intersection returns true if ll1 and ll2 intersect at any not-nil node (by reference)
// by leftpad-ing the shorter list with nils and comparing the lists, if a node is found to be equal,
// that means both lists intersect at that node
func intersection(ll1, ll2 *node) bool {
	if ll1.Len() > ll2.Len() {
		ll2 = ll2.PadLinkedList(ll1.Len() - ll2.Len())
	} else if ll2.Len() > ll1.Len() {
		ll1 = ll1.PadLinkedList(ll2.Len() - ll1.Len())
	}

	for ll1 != nil && ll2 != nil {
		if ll1 == ll2 {
			return true
		}

		ll1 = ll1.Next
		ll2 = ll2.Next
	}

	return false
}
