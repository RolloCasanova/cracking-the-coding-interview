package main

import (
	"fmt"
	"os"

	"github.com/RolloCasanova/ctci-go/2_linked_lists/utils"
)

type node = utils.Node[string]

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <nodes...>")
	}

	ll := utils.ArrayToLinkedList(os.Args[1:])

	ll.Print("Linked list:")

	fmt.Println("Using reverse: ")
	if isPalindromeReverse(ll) {
		fmt.Print("is palindrome\n\n")
	} else {
		fmt.Print("is not palindrome\n\n")
	}

	ll = utils.ArrayToLinkedList(os.Args[1:])
	fmt.Println("Using a stack: ")
	if isPalindromeStack(ll) {
		fmt.Print("is palindrome\n\n")
	} else {
		fmt.Print("is not palindrome\n\n")
	}

	ll = utils.ArrayToLinkedList(os.Args[1:])
	fmt.Println("Using recursion: ")
	if isPalindromeRecursion(ll) {
		fmt.Print("is palindrome\n\n")
	} else {
		fmt.Print("is not palindrome\n\n")
	}
}

// isPalindromeReverse checks if a linked list is a palindrome by reversing it and comparing it to the original
func isPalindromeReverse(ll *node) bool {
	return ll.String() == ll.Reverse().String()
}

// isPalindromeStack checks if a linked list is a palindrome by using a stack
func isPalindromeStack(ll *node) bool {
	root := ll

	stack := []string{}

	for ll != nil {
		stack = append(stack, ll.Value)
		ll = ll.Next
	}

	ll = root

	for ll != nil {
		if stack[len(stack)-1] != ll.Value {
			return false
		}

		stack = stack[:len(stack)-1]
		ll = ll.Next
	}

	return true
}

// isPalindromeRecursion checks if a linked list is a palindrome by using recursion
func isPalindromeRecursion(ll *node) bool {
	return isPalindromeRecursionHelper(ll, ll.Len()).IsPalindrome
}

// recursionResult is the result of the isPalindromeRecursionHelper function
// it contains the node to check and whether it is a palindrome
type recursionResult struct {
	Node         *node
	IsPalindrome bool
}

// isPalindromeRecursionHelper is a helper function for isPalindromeRecursion
func isPalindromeRecursionHelper(ll *node, length int) recursionResult {
	if length == 0 {
		return recursionResult{
			Node:         ll,
			IsPalindrome: true,
		}
	}

	if length == 1 {
		return recursionResult{
			Node:         ll.Next,
			IsPalindrome: true,
		}
	}

	res := isPalindromeRecursionHelper(ll.Next, length-2)

	if !res.IsPalindrome || res.Node == nil {
		return res
	}

	res.IsPalindrome = ll.Value == res.Node.Value
	res.Node = res.Node.Next

	return res
}
