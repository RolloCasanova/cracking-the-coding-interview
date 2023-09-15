package main

import (
	"fmt"
	"os"
	"strconv"

	ll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func main() {
	// root is mandatory
	// nodes will be used to generate a binary search tree
	if len(os.Args) < 2 {
		panic("usage: go run main.go <root_value> <nodes>...")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("root must be a number")
	}

	t := &BST{Value: n}

	for i := 2; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			panic("root must be a number")
		}

		t = t.insert(n)
	}

	lists := ListOfDepths(t)

	for i, l := range lists {
		fmt.Print("level: ", i, "\t")
		for _, v := range l.Values() {
			fmt.Printf("%d ", v.(int))
		}
		fmt.Println()
	}
}

// ListOfDepths returns an array of linked list for each level of the tree
func ListOfDepths(t *BST) []ll.List {
	lists := make([]ll.List, 0)

	q := []*BST{t}
	for len(q) > 0 {
		nextQ := []*BST{}

		list := ll.List{}
		for _, node := range q {
			if node == nil {
				continue
			}
			list.Add(node.Value)
			nextQ = append(nextQ, node.Left, node.Right)
		}

		if len(list.Values()) != 0 {
			lists = append(lists, list)
		}

		q = nextQ
	}

	return lists
}

func (t *BST) insert(n int) *BST {
	if t == nil {
		return &BST{Value: n}
	}

	if n <= t.Value {
		t.Left = t.Left.insert(n)
		return t
	}

	t.Right = t.Right.insert(n)

	return t
}
