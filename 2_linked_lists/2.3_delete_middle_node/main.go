package main

import (
	"os"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

type node = utils.Node[string]

func main() {
	if len(os.Args) < 4 {
		panic("usage: go run main.go <initialNode> <nodeToDelete> <innerNodes...> <lastNode>")
	}

	ll := utils.ArrayToLinkedList[string](os.Args[1:])

	ll.Print("Linked list: ")

	middleNode := ll.Next

	deleteMiddleNode(middleNode)

	ll.Print("Linked list with middle node deleted: ")
}

// deleteMiddleNode deletes a node in the middle (not the first nor the last one) of a linked list
// for simplicity we are always deleting the second node in a linked list of at least 3 nodes
func deleteMiddleNode(node *node) {
	// node is guaranteed to be a middle node, so it can't be the last node
	// no need to check for node.Next == nil

	// simply copy the value and the next reference from the next node to the current node
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}
