package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <stack values...>")
	}

	sm := utils.NewStackMin(len(os.Args[1:]))

	// convert the arguments to integers and push them onto the stack
	for _, v := range os.Args[1:] {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic("all values must be integers")
		}

		sm.Push(value)
	}

	// return the minimum value in the stack in O(1) time
	sm.Print("stack: ")
	fmt.Println("minumum value:", StackMin(sm))
}

// StackMin returns the minimum value in the stack in O(1) time
func StackMin(sm utils.StackMin) int {
	return sm.Min()
}
