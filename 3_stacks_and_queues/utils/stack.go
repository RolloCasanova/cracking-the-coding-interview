package utils

import "fmt"

// StackValue is a constraint for the type of values that can be stored in a stack
type StackValue interface {
	~int | ~string
}

// Stack is an interface for a stack data structure
type Stack[T StackValue] interface {
	Push(T)
	Pop() T
	Peek() T
	IsEmpty() bool
}

// stack represents a stack data structure implemented with an array
type stack[T StackValue] struct {
	values []T
}

// stackWithMin represents a stack data structure implemented with an array that contains a return the minimum value

// NewStack returns a new stack with the given capacity
func NewStack[T StackValue](capacity int) Stack[T] {
	return stack[T]{
		values: make([]T, 0, capacity),
	}
}

// Push adds a value to the top of the stack
func (s stack[T]) Push(value T) {
	if len(s.values) == cap(s.values) {
		// gracefully handle stack overflow
		fmt.Println("stack overflow")
		return
	}

	s.values = append(s.values, value)
}

// Pop returns the top value of the stack by removing it
func (s stack[T]) Pop() (value T) {
	if len(s.values) == 0 {
		// gracefully handle stack underflow
		fmt.Println("empty stack")
		return
	}

	value = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return
}

// Peek returns the top value of the stack without removing it
func (s stack[T]) Peek() (value T) {
	if len(s.values) == 0 {
		// gracefully handle stack underflow
		fmt.Println("empty stack")
		return
	}

	return s.values[len(s.values)-1]
}

// IsEmpty returns whether the stack is empty or not
func (s stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
