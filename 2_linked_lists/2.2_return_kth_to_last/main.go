package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RolloCasanova/cracking-the-coding-interview/2_linked_lists/utils"
)

func main() {
	if len(os.Args) < 3 {
		panic("usage: main.go <k> <nodes...>")
	}

	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("k must be a number")
	}

	ll := utils.ArrayToLinkedList(os.Args[2:])

	ll.Print("Linked list: ")

	fmt.Printf("K: %d\n\n", k)

	kthToLast, err := returnKthToLast(ll, k)
	if err != nil {
		fmt.Printf("an error occurred: %s\n", err.Error())
	} else {
		fmt.Printf("Kth to last element: %s\n", kthToLast)
	}

}

func returnKthToLast(ll *utils.Node, k int) (string, error) {
	if ll == nil {
		return "", fmt.Errorf("linked list is nil")
	}

	if k <= 0 {
		return "", fmt.Errorf("k must be greater than 0")
	}

	runner := ll

	// advance runner k positions
	for i := 0; i < k; i++ {
		// if runner is nil, k is greater than the length of the list
		if runner == nil {
			return "", fmt.Errorf("k is greater than the length of the list")
		}

		runner = runner.Next
	}

	// if runner is nil, k is equal to the length of the list, so first element of ll is the one that should be returned
	if runner == nil {
		return ll.Value, nil
	}

	// advance both pointers until runner reaches the last element of the list
	for runner.Next != nil {
		ll = ll.Next
		runner = runner.Next
	}

	// next element is the kth to last
	ll = ll.Next

	return ll.Value, nil
}
