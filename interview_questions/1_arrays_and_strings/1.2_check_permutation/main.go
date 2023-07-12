package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <string1> <string2>")
		os.Exit(1)
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	isPerm, err := isPermutation(s1, s2)
	if err != nil || !isPerm {
		fmt.Printf("Strings \"%s\" and \"%s\" are not permutations: %s\n", s1, s2, err)
	} else {
		fmt.Printf("Strings \"%s\" and \"%s\" are permutations\n", s1, s2)
	}
}

func isPermutation(s1, s2 string) (bool, error) {
	if len(s1) != len(s2) {
		return false, errors.New("different lenght strings cannot be permutations one of another")
	}

	charMap := make(map[rune]int, 26)

	// adding to a map the character ocurrences of s1
	for _, c := range s1 {
		charMap[c]++
	}

	// substracting to the same map the s2's characters ocurrences
	// if value is zero before substracting, or it doesn't exists, it means that specific character
	// mismatches in the number of ocurrences between the two strings, hence not a permutation
	for _, c := range s2 {
		charCount, ok := charMap[c]
		if !ok {
			return false, fmt.Errorf("character '%c' does not exist in \"%s\"", c, s1)
		}

		if charCount == 0 {
			return false, fmt.Errorf("number ocurrences of '%c' are different that existing ocurrences in \"%s\"", c, s1)
		}

		charMap[c]--
	}

	return true, nil
}
