package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: go run main.go <string1> <string2>")
	}

	if stringRotation(os.Args[1], os.Args[2]) {
		fmt.Printf("\"%s\" is a rotation of \"%s\"\n", os.Args[2], os.Args[1])
	} else {
		fmt.Printf("\"%s\" is not a rotation of \"%s\"\n", os.Args[2], os.Args[1])
	}
}

// stringRotation checks if s2 is a rotation of s1
func stringRotation(s1, s2 string) bool {
	// if s1 and s2 are not the same length, then s2 cannot be a rotation of s1
	if len(s1) != len(s2) {
		return false
	}

	if s1 == "" && s2 == "" {
		return true
	}

	// if s2 is a rotation of s1, then s2 will be a substring of s1s1
	s1s1 := s1 + s1

	for i := 0; i < len(s1s1)-len(s2); i++ {
		if s1s1[i:i+len(s2)] == s2 {
			return true
		}
	}

	return false
}
