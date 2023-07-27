package utils

import "fmt"

// PrintStringMatrix prints a matrix of strings
// leadingZeros is the number of leading zeros to print for each element
func PrintStringMatrix(matrix [][]string, leadingZeros int) {
	var baseSep, sep string = "-", ""

	for i := 0; i < leadingZeros; i++ {
		baseSep += "-"
	}

	for _, row := range matrix {
		for _, c := range row {
			fmt.Printf(fmt.Sprintf("%%0%ds ", leadingZeros), c)
		}

		sep += baseSep
		fmt.Println()
	}

	fmt.Println(sep[:len(sep)-1] + "\n")
}
