package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RolloCasanova/ctci-go/3_stacks_and_queues/utils"
)

// hack to get and use the type of the stackMin struct
type stackMinType = utils.StackMinType

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
	sm.Print()
	fmt.Println("minumum value:", StackMin(sm))
}

// StackMin returns the minimum value in the stack in O(1) time
func StackMin(sm stackMinType) int {
	return sm.Min()
}
