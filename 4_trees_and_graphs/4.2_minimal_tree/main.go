package main

import (
	"fmt"
	"os"
	"strconv"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func main() {
	// n denotes the limit of the given ascending order array, starting from 1
	// this means, n must be positive integer
	if len(os.Args) < 2 {
		panic("usage: go run main.go <n>")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("n must be an integer")
	}

	if n < 1 {
		panic("n must be greater than 0")
	}

	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	t := MinimalTree(nums)

	t.print()
}

// print bst in level order
// each level is separated by a new line
// every two nodes represent the children of the same parent
func (b *BST) print() {
	q := []*BST{b}

	for len(q) > 0 {
		var nextQ []*BST

		for _, n := range q {
			if n == nil {
				continue
			}

			fmt.Printf("(%d) ", n.Value)
			nextQ = append(nextQ, n.Left, n.Right)
		}

		fmt.Println()
		q = nextQ
	}
	fmt.Println()
}

func MinimalTree(nums []int) *BST {
	switch len(nums) {
	case 0:
		return nil
	case 1:
		return &BST{
			Value: nums[0],
		}

	case 2:
		return &BST{
			Value: nums[1],
			Left:  &BST{Value: nums[0]},
		}
	case 3:
		return &BST{
			Value: nums[1],
			Left:  &BST{Value: nums[0]},
			Right: &BST{Value: nums[2]},
		}
	}

	midIndex := len(nums) / 2

	return &BST{
		Value: nums[midIndex],
		Left:  MinimalTree(nums[0:midIndex]),
		Right: MinimalTree(nums[midIndex+1:]),
	}
}
