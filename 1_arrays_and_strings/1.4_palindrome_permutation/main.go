package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("usage: go run main.go <string>")
	}

	s := os.Args[1]

	if isPalindromePermutation(s) {
		fmt.Printf("it is possible to create a palindrome with \"%s\" string's characters", s)
	} else {
		fmt.Printf("it is impossible to create a palindrome with \"%s\" string's characters", s)
	}
}

func isPalindromePermutation(s string) bool {
	// vector will hold the parity of the characters (0 means is even, 1 means is odd)
	// least significant bit represents 'a' (a=1, b=2 ... z=26)
	// there are two conditions in which a palindrome can be constructed:
	//   - all characters are even
	//   - all characters are even except for only one character
	var vector uint32

	s = strings.ToLower(s)
	for i := range s {
		// ignoring whitespace
		if s[i] == ' ' {
			continue
		}

		// what bit position is the i-th character
		pos := s[i] - 'a'

		// flip bit of specific character by using XOR
		vector = vector ^ (1 << pos)
	}

	// to be a palindrome permutation, vector must be 0 or a power of 2
	// a trick to check if a number is a power of 2 is to check if it is 0 or if it the number AND the number minus 1 equals 0
	return vector == 0 || (vector&(vector-1) == 0)
}
