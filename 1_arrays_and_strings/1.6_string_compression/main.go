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

	compressed := stringCompression(s)

	fmt.Printf("  Original: %s\nCompressed: %s\n", s, compressed)
}

// stringCompression compresses a string using the counts of repeated characters
// if the compressed string is not smaller than the original string, it returns the original string
func stringCompression(s string) string {
	// safe check for empty strings
	if len(s) == 0 {
		return s
	}

	// initialize with first character and counter of that char as 1
	char := rune(s[0])
	counter := 1

	// store compressed string
	var compressed strings.Builder

	for _, c := range s[1:] {
		if c == char {
			counter++
			continue
		}

		// a distinct character than previous one, should concat to compressed
		compressed.WriteString(fmt.Sprintf("%c%d", char, counter))

		// can quick exit program exec - not need to iterate over the whole string
		if compressed.Len() >= len(s) {
			return s
		}

		// start counting next character's iterations
		char = c
		counter = 1
	}

	// concat the last compressed part
	compressed.WriteString(fmt.Sprintf("%c%d", char, counter))

	if compressed.Len() >= len(s) {
		return s
	}

	return compressed.String()
}
