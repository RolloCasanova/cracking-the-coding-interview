package main

import (
	"fmt"
	"os"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

type stack = utils.Stack[string]

func main() {
	if len(os.Args) < 2 {
		panic("Usage: go run main.go <values...>")
	}

	// create two stacks
	s1 := utils.NewStack[string](len(os.Args) - 1)
	s2 := utils.NewStack[string](len(os.Args) - 1)

	// push values onto s1
	for i := 1; i < len(os.Args); i++ {
		s1.Push(os.Args[i])
	}

	qvs := QueueViaStacks(s1, s2, len(os.Args)-1)

	var operation, value string
	// read operations from stdin and perform them on the queueWithStacks instance
	for {
		// some operations don't require a value
		// so we only need to read two values

		fmt.Println("enter operation (enqueue, dequeue, peek, isEmpty, print), or exit to quit:")
		_, err := fmt.Scan(&operation)
		if err != nil {
			break
		}

		// if the operation is enqueue, we need to read the value
		if operation == "enqueue" {
			fmt.Println("enter value to enqueue: ")
			_, err := fmt.Scan(&value)
			if err != nil {
				panic("enqueue operation requires a value")
			}
		}

		switch operation {
		case "enqueue":
			qvs.enqueue(value)
		case "dequeue":
			qvs.dequeue()
		case "peek":
			qvs.peek()
		case "isEmpty":
			if qvs.isEmpty() {
				fmt.Println("queueWithStacks is empty")
			} else {
				fmt.Println("queueWithStacks is not empty")
			}
		case "print":
			qvs.print()
		default:
			fmt.Println("invalid operation \" " + operation + "\"")
		}

		fmt.Println()
	}
}

type queueWithStacks struct {
	s1   stack
	s2   stack
	size int
}

// QueueViaStacks creates a new queueWithStacks instance
// s1 is the stack that will be used for enqueue operations by pushing values onto it
// s2 is the stack that will be used for dequeue operations by popping all values from s1 and pushing them onto s2, and then popping the top value from s2
// s1 and s2 should be the same size
func QueueViaStacks(s1, s2 stack, size int) *queueWithStacks {
	if s1 == nil || s2 == nil {
		fmt.Println("s1 and s2 must not be nil")

		return nil
	}

	if size <= 0 {
		fmt.Println("size must be greater than 0")

		return nil
	}

	return &queueWithStacks{
		s1:   s1,
		s2:   s2,
		size: size,
	}
}

// enqueue pushes a value onto s1 (the stack used for enqueue operations)
// if either s1 or s2 is full, enqueue operation fails
// pop all values from s2 and push them onto s1 one by one
func (qvs *queueWithStacks) enqueue(value string) {
	if qvs.s1.IsFull(qvs.size) || qvs.s2.IsFull(qvs.size) {
		fmt.Println("enqueue operation failed: queueWithStacks is full")

		return
	}

	// pop all values from s2 and push them onto s1
	if qvs.s1.IsEmpty() {
		for !qvs.s2.IsEmpty() {
			qvs.s1.Push(qvs.s2.Pop())
		}
	}

	// push value onto s1
	qvs.s1.Push(value)
}

// dequeue pops the top value from s2 (the stack used for dequeue operations)
// if s2 is empty, pop all values from s1 and push them onto s2
// if s1 and s2 are both empty, dequeue operation fails
func (qvs *queueWithStacks) dequeue() {
	if qvs.s1.IsEmpty() && qvs.s2.IsEmpty() {
		fmt.Println("dequeue operation failed: queueWithStacks is empty")

		return
	}

	// pop all values from s1 and push them onto s2
	if qvs.s2.IsEmpty() {
		for !qvs.s1.IsEmpty() {
			qvs.s2.Push(qvs.s1.Pop())
		}
	}

	qvs.s2.Pop()
}

// peek prints the top value from s2 (the stack used for dequeue operations)
// if s2 is empty, pop all values from s1 and push them onto s2
// if s1 and s2 are both empty, peek operation fails
func (qvs *queueWithStacks) peek() {
	if qvs.s1.IsEmpty() && qvs.s2.IsEmpty() {
		fmt.Println("peek operation failed: queueWithStacks is empty")

		return
	}

	// pop all values from s1 and push them onto s2
	if qvs.s2.IsEmpty() {
		for !qvs.s1.IsEmpty() {
			qvs.s2.Push(qvs.s1.Pop())
		}
	}

	fmt.Println(qvs.s2.Peek())
}

// isEmpty checks if queueWithStack is empty by validating both s1 and s2 are empty
func (qvs *queueWithStacks) isEmpty() bool {
	return qvs.s1.IsEmpty() && qvs.s2.IsEmpty()
}

// print prints the values in the queueWithStacks instance in the order they would be dequeued
// if s2 is empty, pop all values from s1 and push them onto s2
func (qvs *queueWithStacks) print() {
	// pop all values from s1 and push them onto s2
	if qvs.s2.IsEmpty() {
		for !qvs.s1.IsEmpty() {
			qvs.s2.Push(qvs.s1.Pop())
		}
	}

	qvs.s2.Print("queueWithStacks: ")
}
