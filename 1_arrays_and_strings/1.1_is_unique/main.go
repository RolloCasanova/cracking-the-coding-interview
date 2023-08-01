package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: go run main.go <string>")
	}

	s := os.Args[1]

	unique, err := isUnique(s)

	if err != nil || !unique {
		fmt.Printf("String \"%s\" does not contain unique characters: %s\n", s, err)
	} else {
		fmt.Printf("String \"%s\" contains unique characters\n", s)
	}
}

// isUnique checks if a string contains only unique characters
func isUnique(s string) (bool, error) {
	// impossible for a string to contain more than 26 characters
	if len(s) > 26 {
		return false, errors.New("string larger than 26")
	}

	// each bit will represent if n-th character exists
	// least significant bit represents 'a' (a=1, b=2 ... z=26)
	var bitMap uint32

	for _, c := range s {
		pos := c - 'a' // this will result in a value from 0 to 25

		if bitMap&(1<<pos) != 0 {
			// if bit is already set then character is not unique
			return false, fmt.Errorf("character '%c' is not unique", c)
		}

		// append character to bitmap
		bitMap = bitMap | (1 << pos)
	}

	return true, nil
}
