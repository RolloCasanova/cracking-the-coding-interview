package utils

import "fmt"

// QueueType is a type that can be used in a queue
type QueueType interface {
	~int | ~string
}

// Queue is the interface for a queue data structure
type Queue[T QueueType] interface {
	Enqueue(v T)
	Dequeue() T
	Peek() T
	IsEmpty() bool
}

// queue is a queue data structure
type queue[T QueueType] struct {
	values []T
	last   int
	first  int
}

// NewQueue creates a new queue data structure
// first represents the index of the first value in the enque (first in first out)
// last represents the index of the last value in the queue
// first will always be greater than or equal to last
// if first == last, the queue is empty
// if first == len(values), the queue is full
func NewQueue[T QueueType](capacity int) *queue[T] {
	return &queue[T]{
		values: make([]T, 0, capacity),
		first:  0,
		last:   0,
	}
}

// Enqueue adds a value to the end of the queue
func (q *queue[T]) Enqueue(value T) {
	if q == nil {
		fmt.Println("queue is nil")

		return
	}

	if q.first == cap(q.values) {
		if q.last == 0 {
			// queue is full
			fmt.Println("queue is full")

			return
		}

		// shift values to the left
		for i := q.last; i < q.first; i++ {
			q.values[i-q.last] = q.values[i]
		}

		q.first -= q.last
		q.last = 0
		q.values[q.first] = value
		q.first++

		return
	}

	q.values = append(q.values, value)
	q.first++
}

// Dequeue removes a value from the start of the queue
func (q *queue[T]) Dequeue() (value T) {
	if q == nil {
		fmt.Println("queue is nil")

		return
	}

	if q.IsEmpty() {
		fmt.Println("queue is empty")

		// shift values to the left if necessary
		if q.last != 0 {
			q.last = 0
			q.first = 0
		}

		return
	}

	value = q.values[q.last]
	q.last++

	return value
}

// Peek returns the value at the start (first) of the queue
func (q *queue[T]) Peek() (value T) {
	if q == nil {
		fmt.Println("queue is nil")

		return
	}

	if q.IsEmpty() {
		fmt.Println("queue is empty")

		return
	}

	return q.values[q.last]
}

// IsEmpty returns true if the queue is empty or nil
func (q *queue[T]) IsEmpty() bool {
	if q == nil {
		fmt.Println("queue is nil")

		return true
	}

	return q.last == q.first
}
