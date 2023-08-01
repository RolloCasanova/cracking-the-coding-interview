package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		panic("usage: go run main.go <string> <true_length>")

	}
	s := os.Args[1]
	trueLength, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("second parameter must be an integer")
	}

	urldString := urlify(s, trueLength)
	fmt.Printf("True length: %d\nOriginal string: \"%s\"\nURLified string: \"%s\"\n", trueLength, s, urldString)
}

// urlify replaces all spaces in a string with '%20' and returns the new string
// trueLength is the length of the string without trailing spaces
func urlify(s string, trueLength int) string {
	var urldString strings.Builder

	for i := 0; i < trueLength && i < len(s); i++ {
		if s[i] == ' ' {
			urldString.WriteString("%20")
		} else {
			urldString.WriteByte(s[i])
		}
	}

	return urldString.String()
}
