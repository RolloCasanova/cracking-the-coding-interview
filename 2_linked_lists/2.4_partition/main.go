package main

import (
	"os"
	"strconv"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

type node = utils.Node[int]

func main() {
	if len(os.Args) < 3 {
		panic("usage: go run main.go <partition> <initialNode> <nodes...>")
	}

	partition, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("partition must be an integer")
	}

	ints, err := utils.StringArrayToIntArray(os.Args[2:])
	if err != nil {
		panic(err)
	}

	ll := utils.ArrayToLinkedList[int](ints)

	ll.Print("Linked list: ")

	ll = partitionLinkedList(ll, partition)

	ll.Print("Linked list partitioned around " + os.Args[1] + ": ")
}

// partitionLinkedList rearranges the linked list around the value of the partition
// on the left side of the partition, all nodes should have a value less than the partition
// on the right side of the partition, all nodes should have a value greater than or equal to the partition
func partitionLinkedList(ll *node, partition int) *node {
	// use two linked links approach; left and right, to store values to the left and the right of the partition value
	// at the end of the iteration we will merge both linked lists

	// left linked list
	var leftHead, leftTail *node

	// right linked list
	var rightHead, rightTail *node

	for ll != nil {
		if ll.Value < partition {
			if leftHead == nil {
				leftHead = ll
				leftTail = ll
			} else {
				leftTail.Next = ll
				leftTail = leftTail.Next
			}
		} else {
			if rightHead == nil {
				rightHead = ll
				rightTail = ll
			} else {
				rightTail.Next = ll
				rightTail = rightTail.Next
			}
		}

		ll = ll.Next
	}

	// if there are no nodes on the left side of the partition, return the right linked list
	if leftHead == nil {
		return rightHead
	}

	// if there are no nodes on the right side of the partition, return the left linked list
	if rightHead == nil {
		return leftHead
	}

	// right tail might be pointing to a certain node that might be part of the left linked list
	// we need to make sure that the right tail is pointing to nil
	rightTail.Next = nil

	// merge both linked lists
	leftTail.Next = rightHead

	// return the head of the merged linked list
	return leftHead
}
