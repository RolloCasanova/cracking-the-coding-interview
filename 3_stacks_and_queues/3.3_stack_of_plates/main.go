package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

func main() {
	// stack size is the maximum number of elements in each stack
	// number of stacks is the number of stacks in the set
	// n represent the number of elements to add to the stack set, and will panic if n > stack size * number of stacks
	if len(os.Args) < 4 {
		panic("usage: go run main.go <stack size> <number of stacks> <n>")
	}

	// convert the arguments to integers
	stackSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("stack size must be an integer")
	}

	numStacks, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("number of stacks must be an integer")
	}

	n, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic("n must be an integer")
	}

	if n > stackSize*numStacks {
		panic("n must be less than stack size * number of stacks")
	}

	// create a new stack set
	ss := StackSet(stackSize, numStacks, n)

	ss.Print("stack set: ")

	// start reading instructions from stdin to the stackset until EOF
	// the instructions are of the form:
	// <operation> <value>
	// e.g.
	// push 1
	// pop
	// peek
	// isEmpty
	// isFull
	// print
	// popAt 1

	var stack, value int
	var operation string

	for {
		// some operations don't require a value
		// so we only need to read two values

		fmt.Println("enter operation (push val, pop, peek, isEmpty, isFull, print, popAt stack):")
		_, err := fmt.Scan(&operation)
		if err != nil {
			break
		}

		// if the operation is push, we need to read the value
		if operation == "push" {
			_, err := fmt.Scan(&value)
			if err != nil {
				panic("push operation requires a value")
			}
		}

		// if the operation is popAt, we need to read the stack number
		if operation == "popAt" {
			_, err := fmt.Scan(&stack)
			if err != nil {
				panic("popAt operation requires a stack number")
			}
		}

		switch operation {
		case "push":
			ss.Push(value)
		case "pop":
			ss.Pop()
		case "peek":
			ss.Peek()
		case "isEmpty":
			if ss.IsEmpty() {
				fmt.Println("stack set is empty")
			} else {
				fmt.Println("stack set is not empty")
			}
		case "isFull":
			if ss.IsFull(ss.Len()) {
				fmt.Println("stack set is full")
			} else {
				fmt.Println("stack set is not full")
			}
		case "print":
			ss.Print("stack set: ")
		case "popAt":
			if val, err := ss.PopAt(stack); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("popped value:", val, "from stack:", stack)
			}
		default:
			fmt.Println("invalid operation \" " + operation + "\"")
		}

		fmt.Println()
	}

}

func StackSet(stackSize, numStacks, n int) utils.StackSet {
	stackSet := utils.NewStackSet(stackSize, numStacks)

	// add n elements to the stack set
	for i := 0; i < n; i++ {
		stackSet.Push(i + 1)
	}

	return stackSet
}
