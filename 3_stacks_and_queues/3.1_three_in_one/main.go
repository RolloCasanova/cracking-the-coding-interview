package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type threeStack struct {
	values  []string
	starts  [3]int
	ends    [3]int
	indexes [3]int
}

func newThreeInOne(size int, stack1, stack2, stack3 []string) *threeStack {
	return &threeStack{
		values:  make([]string, 3*size),
		starts:  [3]int{0, size, 2 * size},
		ends:    [3]int{size - 1, 2*size - 1, 3*size - 1},
		indexes: [3]int{0, size, 2 * size},
	}
}

func main() {
	// stack size denotes the size of each single stack (the array will be 3x this size)
	// stack values are the values to be pushed onto the stack
	if len(os.Args) < 7 {
		panic("usage: go run main.go <stack size> <stack1 values...> <+> <stack2 values...> <+> <stack3 values...>")
	}

	// ensure that the stack size is an integer
	stackSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("stack size must be an integer")
	}

	firstPlusIndex, secondPlusIndex := -1, -1

	// ensure there are exactly two "+" separators in the arguments
	for i, arg := range os.Args {
		if arg == "+" {
			if firstPlusIndex == -1 {
				firstPlusIndex = i
			} else if secondPlusIndex == -1 {
				secondPlusIndex = i
			} else {
				panic("too many \"+\" separators")
			}
		}
	}

	// ensure none of the stacks is greater than the stack size
	stack1 := os.Args[2:firstPlusIndex]
	stack2 := os.Args[firstPlusIndex+1 : secondPlusIndex]
	stack3 := os.Args[secondPlusIndex+1:]

	if len(stack1) > stackSize {
		panic("number of elements in first stack exceeds stack size")
	}

	if len(stack2) > stackSize {
		panic("number of elements in second stack exceeds stack size")
	}

	if len(stack3) > stackSize {
		panic("number of elements in third stack exceeds stack size")
	}

	// create a new threeStack
	ts := threeInOne(stackSize, stack1, stack2, stack3)

	// start reading instructions from stdin to the stacks until EOF
	// the instructions are of the form:
	// 	<stack number> <operation> <value>
	// e.g.
	// 0 push M
	// 1 pop
	// 2 peek
	// 0 isEmpty
	// 0 print

	var stack int
	var operation, value string

	for {
		// some operations don't require a value
		// so we only need to read two values

		fmt.Println("enter stack number and operation (push, pop, peek, isEmpty, print), or exit to quit:")
		_, err := fmt.Scan(&stack, &operation)
		if err != nil {
			break
		}

		// if the operation is push, we need to read the value
		if operation == "push" {
			fmt.Println("enter value to push: ")
			_, err := fmt.Scan(&value)
			if err != nil {
				panic("push operation requires a value")
			}
		}

		switch operation {
		case "push":
			ts.push(stack, value)
		case "pop":
			ts.pop(stack)
		case "peek":
			ts.peek(stack)
		case "isEmpty":
			ts.isEmpty(stack)
		case "print":
			ts.print(stack)
		default:
			fmt.Println("invalid operation \" " + operation + "\"")
		}

		fmt.Println()
	}
}

func threeInOne(stackSize int, stack1, stack2, stack3 []string) *threeStack {
	// create a new threeStack
	ts := newThreeInOne(stackSize, stack1, stack2, stack3)

	// push the values for each stack into the threeStack
	for _, v := range stack1 {
		ts.push(0, v)
	}

	for _, v := range stack2 {
		ts.push(1, v)
	}

	for _, v := range stack3 {
		ts.push(2, v)
	}

	return ts
}

func (ts *threeStack) push(stack int, value string) {
	if ts.indexes[stack] > ts.ends[stack] {
		panic("stack overflow in stack " + strconv.Itoa(stack) + " while trying to push \"" + value + "\"")
	}

	fmt.Println("pushing \"" + value + "\" onto stack " + strconv.Itoa(stack))
	ts.values[ts.indexes[stack]] = value
	ts.indexes[stack]++
}

func (ts *threeStack) pop(stack int) string {
	ts.indexes[stack]--
	if ts.indexes[stack] < ts.starts[stack] {
		panic("stack underflow in stack " + strconv.Itoa(stack) + " while trying to pop value")
	}

	val := ts.values[ts.indexes[stack]]
	fmt.Println("popping \"" + val + "\" from stack " + strconv.Itoa(stack))

	ts.values[ts.indexes[stack]] = ""

	return val
}

func (ts *threeStack) peek(stack int) string {
	var val string
	if ts.indexes[stack] == ts.starts[stack] {
		val = "empty stack"
	} else {
		val = ts.values[ts.indexes[stack]-1]
	}

	fmt.Println("peeking at top of stack", strconv.Itoa(stack)+":", val)

	return val
}

func (ts *threeStack) isEmpty(stack int) bool {
	if ts.indexes[stack] == ts.starts[stack] {
		fmt.Println("stack " + strconv.Itoa(stack) + " is empty")
	} else {
		fmt.Println("stack " + strconv.Itoa(stack) + " is not empty")
	}

	return ts.indexes[stack] == ts.starts[stack]
}

func (ts *threeStack) print(stack int) {
	if ts.indexes[stack] == ts.starts[stack] {
		fmt.Println("stack " + strconv.Itoa(stack) + " is empty")

		return
	}

	fmt.Println("stack", strconv.Itoa(stack), ":", strings.Join(ts.values[ts.starts[stack]:ts.indexes[stack]], ", "))
}
