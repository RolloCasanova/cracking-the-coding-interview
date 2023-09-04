package main

import (
	"os"
	"strconv"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

func main() {
	// read the stack elements from stdin and push them onto the stack
	// the first argument is a boolean indicating whether the stack should be sorted in ascending order
	// the rest of the arguments are the stack elements
	if len(os.Args) < 3 {
		panic("usage: go run main.go <is_ascending> <stack_elements...>")
	}

	s := utils.NewStack[int](len(os.Args[2:]))
	for _, v := range os.Args[2:] {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		s.Push(val)
	}

	isAscending := os.Args[1] == "true"

	s.Print("original stack:")

	s2 := SortStack(s, isAscending)
	s2.Print("sorted stack:")

}

// SortStack sorts the given stack in ascending order if isAscending is true, or descending order if isAscending is false
func SortStack(s utils.Stack[int], isAscending bool) utils.Stack[int] {
	s2 := utils.NewStack[int](s.Len())

	for !s.IsEmpty() {
		// pop the top element off the stack
		top := s.Pop()

		// while the top element of s2 is less than the top element of s
		// pop the top element off s2 and push it onto s
		for !s2.IsEmpty() && ((isAscending && s2.Peek() < top) || (!isAscending && s2.Peek() > top)) {
			val := s2.Pop()
			s.Push(val)
		}

		// push the top element of s onto s2
		s2.Push(top)
	}

	// pop the elements off s2 and push them onto s
	for !s2.IsEmpty() {
		val := s2.Pop()
		s.Push(val)
	}

	return s
}
