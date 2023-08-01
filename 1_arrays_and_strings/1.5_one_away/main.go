package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic("usage: go run main.go <string1> <string2")
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	if oneAway(s1, s2) {
		fmt.Printf("Strings \"%s\" and \"%s\" are one (or zero) edit(s) from each other other\n", s1, s2)
	} else {
		fmt.Printf("Strings \"%s\" and \"%s\" are more than one edit from each other\n", s1, s2)
	}
}

// oneAway checks if two strings are at most one edit away from each other
// an edit is defined as an insertion, deletion or replacement of a character
func oneAway(s1, s2 string) bool {
	l1, l2 := len(s1), len(s2)

	// assign to s1 and l1 the value of the longest string
	if l2 > l1 {
		s1, s2 = s2, s1
		l1, l2 = l2, l1
	}

	// both strings can be different on lenght at most by 1 character
	if l1-l2 > 1 {
		return false
	}

	// if same length, the only way to make both strings the same is modifying a distinct character
	if l1 == l2 {
		for i := range s1 {
			if s1[i] != s2[i] {
				// if only one change is present, the rest of the string should be the same
				return s1[:i]+s1[i+1:] == s2[:i]+s2[i+1:]
			}
		}
		return true
	}

	// lenght diff of one, the only wat to make both strings the same is by adding an extra character
	// ranging over lowest length string to panic
	for i := range s2 {
		if s1[i] != s2[i] {
			// aisling the distinct character from s1, the rest of the string should be the same as s2
			return s1[:i]+s1[i+1:] == s2
		}
	}

	return true
}
