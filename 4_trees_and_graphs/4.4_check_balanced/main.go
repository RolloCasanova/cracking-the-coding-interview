package main

import (
	"os"
	"strconv"

	"github.com/RolloCasanova/ctci-go/4_trees_and_graphs/utils/bst"
)

type BT struct {
	Value int
	Left  *BT
	Right *BT
}

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <root> <nodes>...")
	}

	// convert array of strings to array of ints
	values := make([]int, 0, len(os.Args[1:]))
	for _, value := range os.Args[1:] {
		n, err := strconv.Atoi(value)
		if err != nil {
			panic("values must be integers")
		}

		values = append(values, n)
	}

	bt := bst.New(values...)

	if CheckBalanced(bt, 0) == -1 {
		println("tree is not balanced")
	} else {
		println("tree is balanced")
	}
}

// CheckBalanced checks if a binary tree is balanced and returns its height while doing so
// if the tree is not balanced, it returns -1, otherwise it returns the height of the tree
func CheckBalanced(bt *bst.BST, height int) int {
	if bt == nil {
		return height
	}

	leftHeight := CheckBalanced(bt.Left, height+1)
	rightHeight := CheckBalanced(bt.Right, height+1)

	if leftHeight == -1 || rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
