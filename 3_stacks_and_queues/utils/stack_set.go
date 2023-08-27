package utils

import "fmt"

// StackSet is the interface for a stack set
type StackSet interface {
	Stack[int]
	PopAt(int) (int, error)
}

// stackSet represents a set of stacks
type stackSet struct {
	stacks    []*stack[int]
	stackSize int
}

// NewStackSet returns a new stack set
func NewStackSet(stackSize, numStacks int) *stackSet {
	stacks := make([]*stack[int], numStacks)

	for i := 0; i < numStacks; i++ {
		stacks[i] = NewStack[int](stackSize)
	}

	return &stackSet{
		stacks:    stacks,
		stackSize: stackSize,
	}
}

// Push pushes a value onto the stack set
func (ss *stackSet) Push(value int) {
	if ss.IsFull(ss.stackSize) {
		// gracefully handle stackset overflow
		fmt.Println("stack set is full")

		return
	}

	// find the first stack available in the set to push the value
	for i := range ss.stacks {
		if !ss.stacks[i].IsFull(ss.stackSize) {
			ss.stacks[i].Push(value)

			return
		}
	}
}

// Pop removes the top value of the stack set and returns it
func (ss *stackSet) Pop() (value int) {
	if ss.IsEmpty() {
		// gracefully handle stackset underflow
		fmt.Println("stack set is empty")

		return
	}

	// find the last stack available in the set to pop the value
	for i := len(ss.stacks) - 1; i >= 0; i-- {
		if !ss.stacks[i].IsEmpty() {
			return ss.stacks[i].Pop()
		}
	}

	// this should never be reached, but it is here to satisfy the compiler
	return
}

// Peek returns the top value of the stack set without removing it
func (ss *stackSet) Peek() (value int) {
	if ss.IsEmpty() {
		// gracefully handle stackset underflow
		fmt.Println("stack set is empty")

		return
	}

	// find the last stack available in the set to peek the value
	for i := len(ss.stacks) - 1; i >= 0; i-- {
		if !ss.stacks[i].IsEmpty() {
			return ss.stacks[i].Peek()
		}
	}

	// this should never be reached, but it is here to satisfy the compiler
	return
}

// IsEmpty returns if the stack set is empty by checking if the first stack is also empty
func (ss *stackSet) IsEmpty() bool {
	// ensure all stacks are empty
	for i := range ss.stacks {
		if !ss.stacks[i].IsEmpty() {
			return false
		}
	}

	return true
}

// IsFull returns if the stack set is full by checking if the last stack is also full
func (ss *stackSet) IsFull(capacity int) bool {
	// ensure all stacks are full
	for i := range ss.stacks {
		if !ss.stacks[i].IsFull(capacity) {
			return false
		}
	}

	return true
}

// Print prints the stack set values
func (ss *stackSet) Print(header string) {
	if ss.IsEmpty() {
		fmt.Println(header, "empty stack set")

		return
	}

	if header != "" {
		fmt.Println(header)
	}

	for i := range ss.stacks {
		ss.stacks[i].Print(fmt.Sprintf("stack %d: ", i))
	}
}

// Len returns the size of the stack set
func (ss *stackSet) Len() int {
	return ss.stackSize
}

// PopAt removes the top value of the stack at the given index and returns it
func (ss *stackSet) PopAt(index int) (value int, err error) {
	if index < 0 || index >= len(ss.stacks) {
		return 0, fmt.Errorf("invalid index %d: unable to pop", index)
	}

	if ss.stacks[index].IsEmpty() {
		return 0, fmt.Errorf("stack %d is empty: unable to pop", index)
	}

	return ss.stacks[index].Pop(), nil
}
