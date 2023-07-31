package main

import (
	"os"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

func main() {
	if len(os.Args) < 4 {
		panic("usage: go run main.go <initialNode> <nodeToDelete> <innerNodes...> <lastNode>")
	}

	ll := utils.ArrayToLinkedList(os.Args[1:])

	ll.Print("Linked list: ")

	middleNode := ll.Next

	deleteMiddleNode(middleNode)

	ll.Print("Linked list with middle node deleted: ")
}

func deleteMiddleNode(node *utils.Node) {
	// node is guaranteed to be a middle node, so it can't be the last node
	// no need to check for node.Next == nil

	// simply copy the value and the next reference from the next node to the current node
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}
